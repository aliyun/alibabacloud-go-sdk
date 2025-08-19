// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOSSMountConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*OSSMountPoint) *OSSMountConfig
	GetMountPoints() []*OSSMountPoint
}

type OSSMountConfig struct {
	MountPoints []*OSSMountPoint `json:"mountPoints" xml:"mountPoints" type:"Repeated"`
}

func (s OSSMountConfig) String() string {
	return dara.Prettify(s)
}

func (s OSSMountConfig) GoString() string {
	return s.String()
}

func (s *OSSMountConfig) GetMountPoints() []*OSSMountPoint {
	return s.MountPoints
}

func (s *OSSMountConfig) SetMountPoints(v []*OSSMountPoint) *OSSMountConfig {
	s.MountPoints = v
	return s
}

func (s *OSSMountConfig) Validate() error {
	return dara.Validate(s)
}
