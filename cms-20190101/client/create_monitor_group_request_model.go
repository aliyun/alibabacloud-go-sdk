// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateMonitorGroupRequest
	GetClientToken() *string
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
	// Ensures the idempotence of the request. You can generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The alert contact group. Alert notifications for the application group are sent to the alert contacts in this alert contact group.
	//
	// > An alert contact group is a collection of alert contacts and can contain one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](~~PutContact~~) and [PutContactGroup](~~PutContactGroup~~).
	//
	// example:
	//
	// ECS_Alert_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// Specifies whether to perform a dry run on the request, including permission and instance status checks. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the name and description of the route table are modified.
	//
	// - **true**: Sends a dry run request. Only validation is performed, and the name and description of the route table are not modified. Check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
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

func (s *CreateMonitorGroupRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateMonitorGroupRequest) SetClientToken(v string) *CreateMonitorGroupRequest {
	s.ClientToken = &v
	return s
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
