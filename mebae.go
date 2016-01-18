package main

import (
  "os"
  "github.com/codegangsta/cli"
  "github.com/google/go-github/github"
)

func main() {
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
        // open https://github.com/login/oauth/authorize?client_id=1c32c0ef97c2f71096fa&scope=repo
        println("mebae init OK.")
      }
    }
  }

  app.Action = func(c *cli.Context) {
  }

  app.Run(os.Args)
}
