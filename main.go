package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

func generateProject(name string) {
	fmt.Printf("Hello %q", name)
}

func openFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error)
	}
	return string(data), err
}

func writeFile(filename, content string) {
	ioutil.WriteFile(filename, []byte(content), os.ModePerm)
}

func replacePackage(filename, appName string) {
	source, err := openFile(filename)
	if err != nil {
		panic(err.Error)
	}
	source = strings.Replace(source, "examples.spinning_rotate", appName, -1)
	writeFile(filename, source)
}

func Generate(c *cli.Context) error {
	if c.NArg() != 1 {
		fmt.Println("generate needs one params <app-name>\n")
		cli.ShowAppHelpAndExit(c, 1)
	}
	project_name := c.Args().Get(0)
	fmt.Printf("app name %s", project_name)
	err := exec.Command("git", "clone", "https://github.com/daikikojima/flutter-rotate.git", project_name).Run()
	if err != nil {
		return nil
	}
	replacePackage(project_name+"/pubspec.yaml", project_name)
	replacePackage(project_name+"/web/main.dart", project_name)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "flutter web generator"
	app.Usage = "github.com/daikikojima/flutter-web-generator"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:    "create",           // コマンド名
			Aliases: []string{},         // エイリアス一覧
			Usage:   "View saved memo.", //Usage
			Action:  Generate,           // コマンド実施時に実行されるメソッド
		},
	}
	app.Run(os.Args)

}
