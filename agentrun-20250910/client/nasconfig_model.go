// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNASConfig interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int32) *NASConfig
	GetGroupId() *int32
	SetMountPoints(v []*NASMountConfig) *NASConfig
	GetMountPoints() []*NASMountConfig
	SetUserId(v int32) *NASConfig
	GetUserId() *int32
}

type NASConfig struct {
	// example:
	//
	// 100
	GroupId     *int32            `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*NASMountConfig `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NASConfig) String() string {
	return dara.Prettify(s)
}

func (s NASConfig) GoString() string {
	return s.String()
}

func (s *NASConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *NASConfig) GetMountPoints() []*NASMountConfig {
	return s.MountPoints
}

func (s *NASConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *NASConfig) SetGroupId(v int32) *NASConfig {
	s.GroupId = &v
	return s
}

func (s *NASConfig) SetMountPoints(v []*NASMountConfig) *NASConfig {
	s.MountPoints = v
	return s
}

func (s *NASConfig) SetUserId(v int32) *NASConfig {
	s.UserId = &v
	return s
}

func (s *NASConfig) Validate() error {
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
