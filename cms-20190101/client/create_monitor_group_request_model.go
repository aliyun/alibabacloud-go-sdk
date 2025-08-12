// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v string) *CreateMonitorGroupRequest
	GetContactGroups() *string
	SetGroupName(v string) *CreateMonitorGroupRequest
	GetGroupName() *string
	SetRegionId(v string) *CreateMonitorGroupRequest
	GetRegionId() *string
}

type CreateMonitorGroupRequest struct {
	// The alert contact group. The alert notifications of the application group are sent to the alert contacts that belong to the alert contact group.
	//
	// >  An alert contact group can contain one or more alert contacts. For information about how to create alert contacts and alert contact groups, see [PutContact](~~PutContact~~) and [PutContactGroup](~~PutContactGroup~~).
	//
	// example:
	//
	// ECS_Alert_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *CreateMonitorGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateMonitorGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupRequest) SetContactGroups(v string) *CreateMonitorGroupRequest {
	s.ContactGroups = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetGroupName(v string) *CreateMonitorGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetRegionId(v string) *CreateMonitorGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupRequest) Validate() error {
	return dara.Validate(s)
}
