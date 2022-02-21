/*
Copyright © 2021 Pete Cornish <outofcoffee@gmail.com>

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

package cmd

import (
	"gatehill.io/imposter/workspace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// workspaceNewCmd represents the workspaceNew command
var workspaceNewCmd = &cobra.Command{
	Use:   "new [WORKSPACE_NAME]",
	Short: "Create a workspace",
	Long:  `Creates a new workspace with the given name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createWorkspace(name)
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceNewCmd)
}

func createWorkspace(name string) {
	wd, _ := os.Getwd()
	_, err := workspace.New(wd, name)
	if err != nil {
		logrus.Fatalf("failed to create new workspace: %s", err)
	}
	logrus.Infof("created workspace '%s'", name)
}
