pacakage argparser

import (
    "github.com/jessevdk/go-flags"
)

// Top level options
type Option struct {
    Version         bool            `long:"version" description:"Show version"`
    Help            bool            `long:"help" description:"Show help"`
    Hostname        string          `short:"h" long:"host" description:"MySQL Database's exec hostname (default: localhost)"`
    Username        string          `short:"u" long:"user" description:"MySQL Database's exec user name (default: root)"`
    Password        string          `short:"p" long:"pass" description:"MySQL Database's exec user password (default: mysql)"`
    Backup          *BackupOption   `command:"backup" subcommands-optional:"true" description:"Command to backup [db name]"`
    Restore         *RestoreOption  `command:"restore" subcommands-optional:"true" description:"Command to restore [db name]"`
}

// Backup options
type BackupOption struct {
    BackupPath      string          `long:"backup-path" description:"Saving path of dump file (default: /tmp/backup_db)"`
}

func newBackupOption() *BackupOption {
    opt := &BackupOption{}
    return opt
}

// Arg parse
func newArgParser(opts *Option) *flags.Parser {
    backup := flags.NewNamedParser("mysql-cli", flags.Default)
    backup.AddGroup("Backup Options", "", &BackupOption{})

    opts.BackupOption = newBackupOption()

    parser := flags.NewParser(opts, flags.Default)
    parser.Name = "mysql-cli"
    parser.Usage = "[OPTIONS] [COMMAND] [DB NAME] [COMMAND OPTIONS]"
    return parser
}
