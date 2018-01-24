package fargo

// MIT Licensed (see README.md) - Copyright (c) 2013 Hudl <@Hudl>

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("fargo")
var metadataLog = logging.MustGetLogger("fargo.metadata")
var marshalLog = logging.MustGetLogger("fargo.marshal")

func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	lvlBackend := logging.AddModuleLevel(backend)
	log.SetBackend(lvlBackend)
	metadataLog.SetBackend(lvlBackend)
	marshalLog.SetBackend(lvlBackend)
	logging.SetLevel(logging.INFO, "fargo")
	logging.SetLevel(logging.INFO, "fargo.metadata")
	logging.SetLevel(logging.INFO, "fargo.marshal")
}
