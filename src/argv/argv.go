package argv_P

import (
	"os"

	gleam_P "example.com/todo/gleam"
)

func load() gleam_P.Tuple3_t[gleam_P.String_t, gleam_P.String_t, gleam_P.List_t[gleam_P.String_t]] {
	var args []gleam_P.String_t
	for _, arg := range os.Args {
		args = append(args, gleam_P.String_t(arg))
	}
	return gleam_P.Tuple3_t[gleam_P.String_t, gleam_P.String_t, gleam_P.List_t[gleam_P.String_t]]{"", args[0], gleam_P.ToList(args[1:]...)}
}
