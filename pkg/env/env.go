package env

import (
	"os"

	"k8s.io/klog/v2"
)

func RepositoryEnv() string {
	rep := os.Getenv("GENERATOR_REPOSITORY")
	if rep == "" {
		klog.Fatalf("failed to load GENERATOR_REPOSITORY from envs")
	}
	return rep
}

func RepositoryFullModuleEnv() string {
	rep := os.Getenv("GENERATOR_REPOSITORY_FULL_MODULE")
	if rep == "" {
		klog.Fatalf("failed to load GENERATOR_REPOSITORY_FULL_MODULE from envs")
	}
	return rep
}
