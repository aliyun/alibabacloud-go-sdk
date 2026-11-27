// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMissionConnectorTarget interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *MissionConnectorTarget
	GetDisplayName() *string
	SetId(v string) *MissionConnectorTarget
	GetId() *string
	SetTargetType(v string) *MissionConnectorTarget
	GetTargetType() *string
}

type MissionConnectorTarget struct {
	// The display name of the IM contact or group.
	//
	// example:
	//
	// Alert notification group
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The case-sensitive ID of the IM contact or group.
	//
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The recipient type. Valid values: user for a user, and group for a group.
	//
	// example:
	//
	// group
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s MissionConnectorTarget) String() string {
	return dara.Prettify(s)
}

func (s MissionConnectorTarget) GoString() string {
	return s.String()
}

func (s *MissionConnectorTarget) GetDisplayName() *string {
	return s.DisplayName
}

func (s *MissionConnectorTarget) GetId() *string {
	return s.Id
}

func (s *MissionConnectorTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *MissionConnectorTarget) SetDisplayName(v string) *MissionConnectorTarget {
	s.DisplayName = &v
	return s
}

func (s *MissionConnectorTarget) SetId(v string) *MissionConnectorTarget {
	s.Id = &v
	return s
}

func (s *MissionConnectorTarget) SetTargetType(v string) *MissionConnectorTarget {
	s.TargetType = &v
	return s
}

func (s *MissionConnectorTarget) Validate() error {
	return dara.Validate(s)
}
