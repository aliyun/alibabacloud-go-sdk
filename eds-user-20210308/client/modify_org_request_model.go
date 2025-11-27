// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v string) *ModifyOrgRequest
	GetOrgId() *string
	SetOrgName(v string) *ModifyOrgRequest
	GetOrgName() *string
}

type ModifyOrgRequest struct {
	// The ID of the organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// org-76joc57khvnhdh***
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The name of the organization.
	//
	// This parameter is required.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s ModifyOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOrgRequest) GoString() string {
	return s.String()
}

func (s *ModifyOrgRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *ModifyOrgRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *ModifyOrgRequest) SetOrgId(v string) *ModifyOrgRequest {
	s.OrgId = &v
	return s
}

func (s *ModifyOrgRequest) SetOrgName(v string) *ModifyOrgRequest {
	s.OrgName = &v
	return s
}

func (s *ModifyOrgRequest) Validate() error {
	return dara.Validate(s)
}
