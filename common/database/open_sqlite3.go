//go:build sqlite3
// +build sqlite3

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func openPostgres(source string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  source,
		PreferSimpleProtocol: true,
	})
}

var opens = map[string]func(string) gorm.Dialector{
	"mysql":     mysql.Open,
	"postgres":  openPostgres,
	"sqlite3":   sqlite.Open,
	"sqlserver": sqlserver.Open,
}
