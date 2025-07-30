// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateOrganizationalUnitDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateOrganizationalUnitDescriptionRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitDescriptionRequest
	GetOrganizationalUnitId() *string
}

type UpdateOrganizationalUnitDescriptionRequest struct {
	// The description of the organization. The value can be up to 256 characters in length.
	//
	// example:
	//
	// organizationalUnit_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
}

func (s UpdateOrganizationalUnitDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOrganizationalUnitDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateOrganizationalUnitDescriptionRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetDescription(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetInstanceId(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
