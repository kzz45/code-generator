/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// client-gen makes the individual typed clients using gengo.
package main

import (
	"flag"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	generatorargs "github.com/kzz45/code-generator/cmd/client-gen/args"
	"github.com/kzz45/code-generator/cmd/client-gen/generators"
	"github.com/kzz45/code-generator/pkg/env"
	"github.com/kzz45/code-generator/pkg/util"
)

func main() {
	klog.InitFlags(nil)
	genericArgs, customArgs := generatorargs.NewDefaults()

	klog.Info("starting client-gen")
	// Override defaults.
	// TODO: move this out of client-gen
	genericArgs.OutputPackagePath = env.RepositoryFullModuleEnv() + "/pkg/client-go/"

	genericArgs.AddFlags(pflag.CommandLine)
	customArgs.AddFlags(pflag.CommandLine, env.RepositoryFullModuleEnv()+"/pkg/apis") // TODO: move this input path out of client-gen
	flag.Set("logtostderr", "true")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// add group version package as input dirs for gengo
	for _, pkg := range customArgs.Groups {
		for _, v := range pkg.Versions {
			genericArgs.InputDirs = append(genericArgs.InputDirs, v.Package)
		}
	}

	if err := generatorargs.Validate(genericArgs); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	if err := genericArgs.Execute(
		generators.NameSystems(util.PluralExceptionListToMapOrDie(customArgs.PluralExceptions)),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}
}
