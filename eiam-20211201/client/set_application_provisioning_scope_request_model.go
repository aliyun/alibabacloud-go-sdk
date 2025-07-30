// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationProvisioningScopeRequest
	GetApplicationId() *string
	SetGroupIds(v []*string) *SetApplicationProvisioningScopeRequest
	GetGroupIds() []*string
	SetInstanceId(v string) *SetApplicationProvisioningScopeRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *SetApplicationProvisioningScopeRequest
	GetOrganizationalUnitIds() []*string
}

type SetApplicationProvisioningScopeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// List of groups that are authorized to be synchronized from
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of organizational units that are authorized for account synchronization.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s SetApplicationProvisioningScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningScopeRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationProvisioningScopeRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *SetApplicationProvisioningScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationProvisioningScopeRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *SetApplicationProvisioningScopeRequest) SetApplicationId(v string) *SetApplicationProvisioningScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetGroupIds(v []*string) *SetApplicationProvisioningScopeRequest {
	s.GroupIds = v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetInstanceId(v string) *SetApplicationProvisioningScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetOrganizationalUnitIds(v []*string) *SetApplicationProvisioningScopeRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) Validate() error {
	return dara.Validate(s)
}
