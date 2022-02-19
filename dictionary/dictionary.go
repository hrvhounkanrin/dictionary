package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

type Dictionary struct {
	db *badger.DB
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	createdAt := e.CreatedAt.Format(time.Stamp)

	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, createdAt)
}

func New(dir string) (*Dictionary, error) {
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	dict := &Dictionary{db: db}
	return dict, nil
}

func (d *Dictionary) Close() {
	d.db.Close()
}
