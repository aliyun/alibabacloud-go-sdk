// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
	SetOrganizationalUnitName(v string) *UpdateOrganizationalUnitRequest
	GetOrganizationalUnitName() *string
}

type UpdateOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organization. The name can be up to 128 characters in length and must be unique in the same parent organization.
	//
	// example:
	//
	// ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
}

func (s UpdateOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *UpdateOrganizationalUnitRequest) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *UpdateOrganizationalUnitRequest) SetInstanceId(v string) *UpdateOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *UpdateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *UpdateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *UpdateOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
