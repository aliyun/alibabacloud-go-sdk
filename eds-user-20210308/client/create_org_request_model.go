// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgName(v string) *CreateOrgRequest
	GetOrgName() *string
	SetParentOrgId(v string) *CreateOrgRequest
	GetParentOrgId() *string
}

type CreateOrgRequest struct {
	// The name of the organization.
	//
	// This parameter is required.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The ID of the parent organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// org-evk12ozjvmlxl****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s CreateOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *CreateOrgRequest) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *CreateOrgRequest) SetOrgName(v string) *CreateOrgRequest {
	s.OrgName = &v
	return s
}

func (s *CreateOrgRequest) SetParentOrgId(v string) *CreateOrgRequest {
	s.ParentOrgId = &v
	return s
}

func (s *CreateOrgRequest) Validate() error {
	return dara.Validate(s)
}
