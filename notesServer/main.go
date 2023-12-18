package main

import (
	"flag"
	"myModules/notesServer/controller/httpserver"
	"myModules/notesServer/gates/storage"
	"myModules/notesServer/gates/storage/list"
	"myModules/notesServer/gates/storage/mp"
)

func main() {
	var st storage.Storage
	var useList bool
	flag.BoolVar(&useList, "list", false, "A boolean flag")
	flag.Parse()

	if useList {
		st = list.NewList()
	} else {
		st = mp.NewMap()
	}
	hs := httpserver.NewHttpServer(":8080", st)
	hs.Start()
}
