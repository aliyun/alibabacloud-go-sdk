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
	SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest
	GetOrganizationalUnitExternalId() *string
	SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest
	GetOrganizationalUnitName() *string
	SetParentId(v string) *CreateOrganizationalUnitRequest
	GetParentId() *string
}

type CreateOrganizationalUnitRequest struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// test organizational unit
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The external ID of the organizational unit. The external ID can be used to map external data to the data of the organizational unit in Employee Identity and Access Management (EIAM) of Identity as a Service (IDaaS). By default, the external ID is the organizational unit ID.
	//
	// For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The name of the organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The ID of the parent organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
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
