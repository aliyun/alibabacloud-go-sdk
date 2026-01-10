// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolarFsConfig interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int32) *PolarFsConfig
	GetGroupId() *int32
	SetMountPoints(v []*PolarFsMountConfig) *PolarFsConfig
	GetMountPoints() []*PolarFsMountConfig
	SetUserId(v int32) *PolarFsConfig
	GetUserId() *int32
}

type PolarFsConfig struct {
	GroupId     *int32                `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*PolarFsMountConfig `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PolarFsConfig) String() string {
	return dara.Prettify(s)
}

func (s PolarFsConfig) GoString() string {
	return s.String()
}

func (s *PolarFsConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *PolarFsConfig) GetMountPoints() []*PolarFsMountConfig {
	return s.MountPoints
}

func (s *PolarFsConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *PolarFsConfig) SetGroupId(v int32) *PolarFsConfig {
	s.GroupId = &v
	return s
}

func (s *PolarFsConfig) SetMountPoints(v []*PolarFsMountConfig) *PolarFsConfig {
	s.MountPoints = v
	return s
}

func (s *PolarFsConfig) SetUserId(v int32) *PolarFsConfig {
	s.UserId = &v
	return s
}

func (s *PolarFsConfig) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
