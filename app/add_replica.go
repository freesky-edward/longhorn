package app

import (
	"errors"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/rancher/longhorn/sync"
)

func AddReplicaCmd() cli.Command {
	return cli.Command{
		Name:      "add-replica",
		ShortName: "add",
		Action: func(c *cli.Context) {
			if err := addReplica(c); err != nil {
				logrus.Fatalf("Error running add replica command: %v", err)
			}
		},
	}
}

func addReplica(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("replica address is required")
	}
	replica := c.Args()[0]

	url := c.GlobalString("url")
	task := sync.NewTask(url)
	return task.AddReplica(replica)
}
func AutoAddReplica(frontendIP string, replica string) error {
	url := "http://" + frontendIP + ":9501"
	task := sync.NewTask(url)
	for {
		err := task.AddReplica(replica)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		return err
	}
}
