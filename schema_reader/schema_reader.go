package schema_reader

import "database/sql"
import _ "github.com/go-sql-driver/mysql"


// Hello returns a greeting for the named person.
func fetch(name string) string {

	db, err := sql.Open("mysql", "user:password@/dbname")

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
