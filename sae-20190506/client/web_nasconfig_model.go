// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebNASConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*WebNASMountPoint) *WebNASConfig
	GetMountPoints() []*WebNASMountPoint
}

type WebNASConfig struct {
	MountPoints []*WebNASMountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
}

func (s WebNASConfig) String() string {
	return dara.Prettify(s)
}

func (s WebNASConfig) GoString() string {
	return s.String()
}

func (s *WebNASConfig) GetMountPoints() []*WebNASMountPoint {
	return s.MountPoints
}

func (s *WebNASConfig) SetMountPoints(v []*WebNASMountPoint) *WebNASConfig {
	s.MountPoints = v
	return s
}

func (s *WebNASConfig) Validate() error {
	return dara.Validate(s)
}
