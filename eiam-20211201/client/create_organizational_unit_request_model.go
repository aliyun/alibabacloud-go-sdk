// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateOrganizationalUnitRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest
	GetOrganizationalUnitExternalId() *string
	SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest
	GetOrganizationalUnitName() *string
	SetParentId(v string) *CreateOrganizationalUnitRequest
	GetParentId() *string
}

type CreateOrganizationalUnitRequest struct {
	// The description. The maximum length is 256 characters.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// External ID of the organization, which is used for association with an external system. The default value is the organization ID. The maximum length is 64 characters.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// Organization name. The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// Parent organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOrganizationalUnitRequest) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *CreateOrganizationalUnitRequest) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *CreateOrganizationalUnitRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateOrganizationalUnitRequest) SetDescription(v string) *CreateOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetInstanceId(v string) *CreateOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetParentId(v string) *CreateOrganizationalUnitRequest {
	s.ParentId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
