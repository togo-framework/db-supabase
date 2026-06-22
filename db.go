// Package dbsupabase registers the pgx Postgres driver so togo's ORM can talk to
// Postgres / Supabase. Set DB_DRIVER=pgx and DATABASE_URL to your connection
// string. Install: `togo install togo-framework/db-supabase`.
package dbsupabase

import (
	"github.com/togo-framework/togo"

	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" database/sql driver
)

func init() {
	togo.RegisterProviderFunc("db-supabase", togo.PriorityCore, func(k *togo.Kernel) error {
		if k.Log != nil {
			k.Log.Info("db-supabase: pgx driver registered (set DB_DRIVER=pgx)")
		}
		return nil
	})
}
