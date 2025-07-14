// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebOSSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*WebOSSMountPoint) *WebOSSConfig
	GetMountPoints() []*WebOSSMountPoint
}

type WebOSSConfig struct {
	MountPoints []*WebOSSMountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
}

func (s WebOSSConfig) String() string {
	return dara.Prettify(s)
}

func (s WebOSSConfig) GoString() string {
	return s.String()
}

func (s *WebOSSConfig) GetMountPoints() []*WebOSSMountPoint {
	return s.MountPoints
}

func (s *WebOSSConfig) SetMountPoints(v []*WebOSSMountPoint) *WebOSSConfig {
	s.MountPoints = v
	return s
}

func (s *WebOSSConfig) Validate() error {
	return dara.Validate(s)
}
