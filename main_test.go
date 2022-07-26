//go:build e2e
// +build e2e

// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package main

import (
	"context"
	"testing"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
)

// TestRunMain wraps the main() function in order to build a test binary and collection coverage for
// E2E/Integration tests. Controller CLI flags are also passed in here.
func TestRunMain(t *testing.T) {
	runController(context.TODO(), controllercmd.ControllerContext)
}
