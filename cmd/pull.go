/*
Copyright © 2022 Felipe Macias felipem1210@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/felipem1210/git-helper/githelper"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes of a branch from remote",
	Long:  `Pull changes of a branch from remote. It will make the pull of current branch`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		repoNames := githelper.ListDirectories(target)
		githelper.GitPull(target, repoNames)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
