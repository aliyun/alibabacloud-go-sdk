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
	SetDryRun(v bool) *CreateMonitorGroupRequest
	GetDryRun() *bool
	SetGroupName(v string) *CreateMonitorGroupRequest
	GetGroupName() *string
	SetRegionId(v string) *CreateMonitorGroupRequest
	GetRegionId() *string
}

type CreateMonitorGroupRequest struct {
	// The alert contact group. Alert notifications for the application group are sent to the alert contacts in this alert contact group.
	//
	// > An alert contact group contains one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](~~PutContact~~) and [PutContactGroup](~~PutContactGroup~~).
	//
	// example:
	//
	// ECS_Alert_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// Specifies whether to perform a dry run on the request. A dry run checks items such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the name and description of the route table are modified.
	//
	// - **true**: Sends a dry run request. Only the check is performed, and the name and description of the route table are not modified. The check items include whether required parameters are specified and whether the request format is valid. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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

func (s *CreateMonitorGroupRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateMonitorGroupRequest) SetDryRun(v bool) *CreateMonitorGroupRequest {
	s.DryRun = &v
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
