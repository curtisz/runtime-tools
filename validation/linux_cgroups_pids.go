package main

import (
	"github.com/opencontainers/runtime-tools/cgroups"
	"github.com/opencontainers/runtime-tools/validation/util"
)

func main() {
	var limit int64 = 1000
	g := util.GetDefaultGenerator()
	g.SetLinuxCgroupsPath(cgroups.AbsCgroupPath)
	g.SetLinuxResourcesPidsLimit(limit)
	err := util.RuntimeOutsideValidate(g, util.ValidateLinuxResourcesPids)
	if err != nil {
		util.Fatal(err)
	}
}
