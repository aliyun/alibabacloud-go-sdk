// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOrganizationalUnitParentsRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *ListOrganizationalUnitParentsRequest
	GetOrganizationalUnitId() *string
}

type ListOrganizationalUnitParentsRequest struct {
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

func (s ListOrganizationalUnitParentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitParentsRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitParentsRequest) SetInstanceId(v string) *ListOrganizationalUnitParentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitParentsRequest) SetOrganizationalUnitId(v string) *ListOrganizationalUnitParentsRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitParentsRequest) Validate() error {
	return dara.Validate(s)
}
