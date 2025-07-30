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
	// The description of the organization. The value can be up to 256 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The external ID of the organization, which can be used to associate the organization with an external system. By default, the external ID is the organization ID. The value can be up to 64 characters in length.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// The name of the organization. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The parent organization ID.
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
