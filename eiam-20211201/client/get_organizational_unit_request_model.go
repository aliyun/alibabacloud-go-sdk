// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *GetOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
}

type GetOrganizationalUnitRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s GetOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetOrganizationalUnitRequest) SetInstanceId(v string) *GetOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *GetOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
