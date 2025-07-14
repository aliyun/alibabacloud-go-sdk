// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNASConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*NASMountConfig) *NASConfig
	GetMountPoints() []*NASMountConfig
}

type NASConfig struct {
	MountPoints []*NASMountConfig `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s NASConfig) String() string {
	return dara.Prettify(s)
}

func (s NASConfig) GoString() string {
	return s.String()
}

func (s *NASConfig) GetMountPoints() []*NASMountConfig {
	return s.MountPoints
}

func (s *NASConfig) SetMountPoints(v []*NASMountConfig) *NASConfig {
	s.MountPoints = v
	return s
}

func (s *NASConfig) Validate() error {
	return dara.Validate(s)
}
