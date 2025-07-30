// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitParentIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateOrganizationalUnitParentIdRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitParentIdRequest
	GetOrganizationalUnitId() *string
	SetParentId(v string) *UpdateOrganizationalUnitParentIdRequest
	GetParentId() *string
}

type UpdateOrganizationalUnitParentIdRequest struct {
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
	// The parent organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateOrganizationalUnitParentIdRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *UpdateOrganizationalUnitParentIdRequest) GetParentId() *string {
	return s.ParentId
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetInstanceId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetParentId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdRequest) Validate() error {
	return dara.Validate(s)
}
