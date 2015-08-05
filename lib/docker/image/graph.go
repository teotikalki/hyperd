package image

import (
	"github.com/hyperhq/hyper/lib/docker/daemon/graphdriver"
)

type Graph interface {
	Get(id string) (*Image, error)
	ImageRoot(id string) string
	Driver() graphdriver.Driver
}
