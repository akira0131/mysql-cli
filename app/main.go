package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()

    app.Name = "mysql-cli"
    app.Usage = "cli for backup or restore mysql database."
    app.Version = "1.0.0"

    // Set global options
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "host",
            Value: "localhost",
            Usage: "Set host name for mysql database",
        },
        cli.StringFlag{
            Name:  "user, u",
            Value: "root",
            Usage: "Set user name for mysql database",
        },
        cli.StringFlag{
            Name:   "pass, p",
            EnvVar: "MYSQL_PWD",
            Usage:  "Set user password for mysql database",
        },
    }

    // Run command
    app.Commands = []cli.Command{
        {
            Name:  "backup",
            Usage: "Backup database on mysql",
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "output, o",
                    Value: "/tmp/backup_db",
                    Usage: "Saving path for dump file",
                },
            },
            Action: func(context *cli.Context) error {
                return nil
            },
        },
        {
            Name:  "restore",
            Usage: "Restore database on mysql",
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name:  "output, o",
                    Value: "/tmp/backup_db",
                    Usage: "Saving path for dump file",
                },
            },
            Action: func(context *cli.Context) error {
                fmt.Println("restore database: ", context.Args().First())
                return nil
            },
        },
    }

    // Check args


    app.Run(os.Args)
}
