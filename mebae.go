package main

import (
  "os"
  "os/user"
  "log"
  "io/ioutil"
  "bytes"
  "net/http"
  // "github.com/k0kubun/pp"
  "github.com/BurntSushi/toml"
  "github.com/codegangsta/cli"
  "github.com/bitly/go-simplejson"
  "github.com/skratchdot/open-golang/open"
)

type Config struct {
  GithubName string `toml:"GithubName"`
  AccessToken string `toml:"AccessToken"`
}

func main() {
  // Get Config
  usr, _ := user.Current()
  tomlDir := usr.HomeDir + "/.mebae"
  tomlFile := usr.HomeDir + "/.mebae/config.tml"
  var config Config

  // 設定ファイルが存在すれば読み込み
  if _, err := os.Stat(tomlFile); err == nil {
    if _, err := toml.DecodeFile(tomlFile, &config); err != nil {
      panic(err)
    }
  }

  app := cli.NewApp()
  app.Name = "Mebae"
  app.Usage = "Create Github Repository and more."
  app.Version = "1.0"
  app.Commands = []cli.Command {
    {
      Name:    "init",
      Aliases: []string{"i"},
      Usage:   "get access token for Github",
      Action:  func(c *cli.Context) {
        // ディレクトリ作成
        if err := os.Mkdir(tomlDir, 0777); err != nil {
          // TODO: エラーメッセージ表示後にos.Exitで死ぬように
          // println(err)
        }

        // 設定ファイル作成
        if _, err := os.Stat(tomlFile); err != nil {
          if _,err := os.Create(tomlFile); err != nil {
            // TODO: エラーメッセージ表示後にos.Exitで死ぬように
            // println(err)
          }
        }

        open.Run("https://github.com/login/oauth/authorize?client_id=1c32c0ef97c2f71096fa&scope=repo,delete_repo")
        println("Mebae initialized.")
      },
    },
    {
      Name:    "create",
      Aliases: []string{"c"},
      Usage:   "get access token for Github",
      Action:  func(c *cli.Context) {
        client := &http.Client{}
        send_data := []byte(`{"name": "`+ c.Args().First() + `"}`)
        url := "https://api.github.com/user/repos"

        request, err := http.NewRequest("POST", url, bytes.NewBuffer(send_data))
        request.Header.Add("Authorization", "bearer " + config.AccessToken)
        request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        resp, err := client.Do(request)
        if err != nil {
          log.Fatal(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        js, err := simplejson.NewJson(body)
        if err != nil {
          log.Fatal(err)
        }
        println("Created Repository.")
        println(js.Get("git_url").MustString())
      },
    },
    {
      Name:    "delete",
      Aliases: []string{"d"},
      Usage:   "get access token for Github",
      Action:  func(c *cli.Context) {
        client := &http.Client{}
        repo_name := c.Args().First()
        url := "https://api.github.com/repos/"

        request, err := http.NewRequest("DELETE", url + config.GithubName + "/" + repo_name, nil)
        request.Header.Add("Authorization", "bearer " + config.AccessToken)
        request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        // TODO: レスポンスボディのmessageを読み取って
        // エラーであるならエラー処理する
        resp, err := client.Do(request)
        if err != nil {
          log.Fatal(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          log.Fatal(err)
        }
        // log.Print(string(body))

        println("Deleted Repository.")
      },
    },
  }

  app.Action = func(c *cli.Context) {
  }

  app.Run(os.Args)
}
