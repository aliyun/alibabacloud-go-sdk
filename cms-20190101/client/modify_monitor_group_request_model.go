// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v string) *ModifyMonitorGroupRequest
	GetContactGroups() *string
	SetGroupId(v string) *ModifyMonitorGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *ModifyMonitorGroupRequest
	GetGroupName() *string
	SetRegionId(v string) *ModifyMonitorGroupRequest
	GetRegionId() *string
}

type ModifyMonitorGroupRequest struct {
	// The alert groups that can receive alert notifications for the application group.
	//
	// example:
	//
	// alarm_ecs_group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// ecs_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyMonitorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *ModifyMonitorGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyMonitorGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyMonitorGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMonitorGroupRequest) SetContactGroups(v string) *ModifyMonitorGroupRequest {
	s.ContactGroups = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetGroupId(v string) *ModifyMonitorGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetGroupName(v string) *ModifyMonitorGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetRegionId(v string) *ModifyMonitorGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMonitorGroupRequest) Validate() error {
	return dara.Validate(s)
}
