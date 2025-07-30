// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest
	GetApplicationId() *string
	SetInstanceId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *AuthorizeApplicationToOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
}

type AuthorizeApplicationToOrganizationalUnitsRequest struct {
	// The ID of the application on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations to which you want to grant permissions. You can grant permissions to a maximum of 100 organizations at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s AuthorizeApplicationToOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetApplicationId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetInstanceId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
