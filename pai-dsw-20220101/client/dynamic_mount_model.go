// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDynamicMount interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DynamicMount
	GetEnable() *bool
	SetMountPoints(v []*DynamicMountPoint) *DynamicMount
	GetMountPoints() []*DynamicMountPoint
}

type DynamicMount struct {
	Enable      *bool                `json:"Enable,omitempty" xml:"Enable,omitempty"`
	MountPoints []*DynamicMountPoint `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
}

func (s DynamicMount) String() string {
	return dara.Prettify(s)
}

func (s DynamicMount) GoString() string {
	return s.String()
}

func (s *DynamicMount) GetEnable() *bool {
	return s.Enable
}

func (s *DynamicMount) GetMountPoints() []*DynamicMountPoint {
	return s.MountPoints
}

func (s *DynamicMount) SetEnable(v bool) *DynamicMount {
	s.Enable = &v
	return s
}

func (s *DynamicMount) SetMountPoints(v []*DynamicMountPoint) *DynamicMount {
	s.MountPoints = v
	return s
}

func (s *DynamicMount) Validate() error {
	return dara.Validate(s)
}
