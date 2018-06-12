package models

import (
	"github.com/SSHZ-ORG/dedicatus"
	"github.com/SSHZ-ORG/dedicatus/utils"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Config struct {
	Admins       []int
	Contributors []int
}

const stringKey = "necessarius"

func configKey(ctx context.Context) *datastore.Key {
	return datastore.NewKey(ctx, configEntityKind, stringKey, 0, nil)
}

func GetConfig(ctx context.Context) Config {
	c := Config{}
	datastore.Get(ctx, configKey(ctx), &c)
	return c
}

func CreateConfig(ctx context.Context) error {
	c := GetConfig(ctx)

	admins := utils.NewIntSetFromSlice(c.Admins)
	admins.Add(dedicatus.InitAdminID)

	contributors := utils.NewIntSetFromSlice(c.Contributors)
	contributors.Add(dedicatus.InitAdminID)

	c.Admins = admins.ToSlice()
	c.Contributors = contributors.ToSlice()

	_, err := datastore.Put(ctx, configKey(ctx), &c)
	return err
}
