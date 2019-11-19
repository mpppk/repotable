package cmd

import (
	"context"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/mpppk/repotable/pkg/run"

	"github.com/ghodss/yaml"

	"github.com/mpppk/repotable/internal/option"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

func newOutFlag() *option.StringFlag {
	return &option.StringFlag{
		Flag: &option.Flag{
			Name:       "file",
			Shorthand:  "f",
			Usage:      "Input file path",
			IsRequired: true,
		},
	}
}

func newRunCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     "run",
		Short:   "Generate markdown table",
		Long:    ``,
		Example: "repotable run --file repos.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := option.NewRunCmdConfigFromViper()
			if err != nil {
				return err
			}

			contents, err := ioutil.ReadFile(conf.File)
			if err != nil {
				return err
			}

			var repoConfig run.RepoTableConfig
			if err := yaml.Unmarshal(contents, &repoConfig); err != nil {
				return err
			}

			client := github.NewClient(nil)

			cmd.Println("Name | CreatedAt | Star | Note")
			cmd.Println("---- | ---- | ---- | ----")
			for _, repository := range repoConfig.Repositories {
				repo, _, err := client.Repositories.Get(context.Background(), repository.Owner, repository.Repo)
				if err != nil {
					return err
				}
				values := []string{
					*repo.FullName,
					repo.CreatedAt.Format("2006 Jan 2"),
					strconv.Itoa(*repo.StargazersCount),
					"-",
				}
				cmd.Println(strings.Join(values, " | "))
			}

			return nil
		},
	}
	if err := option.RegisterStringFlag(cmd, newOutFlag()); err != nil {
		return nil, err
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newRunCmd)
}
