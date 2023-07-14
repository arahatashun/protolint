package main

import (
	"flag"

	"github.com/yoheimuta/protolint/_example/plugin/customrules"
	"github.com/yoheimuta/protolint/plugin"
)

func main() {
	flag.Parse()

	plugin.RegisterCustomRules(
		customrules.NewPackageNameDefinedRule(),
	)
}
