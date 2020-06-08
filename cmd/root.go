package cmd

import (
	"21bit-gcr/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var mirrorUrl = "registry.cn-beijing.aliyuncs.com/21bit/"
var rootCmd = &cobra.Command{
	Use:   "21bit-gcr",
	Short: "usage: 21bit-gcr kube-apiserver:v1.18.3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("usage: 21bit-gcr kube-apiserver:v1.18.3")
			return
		}

		// pull images
		if err := utils.RunCommand("docker", "pull", mirrorUrl+args[0]); err != nil {
			fmt.Printf("pull image %s failed", mirrorUrl+args[0])
			return
		}

		// tag new image
		if err := utils.RunCommand("docker", "tag", mirrorUrl+args[0], "k8s.gcr.io/"+args[0]); err != nil {
			fmt.Printf("tag image %s failed", args[0])
			return
		}

		// remove image
		if err := utils.RunCommand("docker", "image", "rm", mirrorUrl+args[0]); err != nil {
			fmt.Printf("remove image %s failed", mirrorUrl+args[0])
			return
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
