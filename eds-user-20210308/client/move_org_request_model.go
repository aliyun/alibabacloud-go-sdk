// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewParentOrgId(v string) *MoveOrgRequest
	GetNewParentOrgId() *string
	SetOrgId(v string) *MoveOrgRequest
	GetOrgId() *string
}

type MoveOrgRequest struct {
	// The ID of the parent organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	NewParentOrgId *string `json:"NewParentOrgId,omitempty" xml:"NewParentOrgId,omitempty"`
	// The ID of the organization that you want to move.
	//
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s MoveOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveOrgRequest) GoString() string {
	return s.String()
}

func (s *MoveOrgRequest) GetNewParentOrgId() *string {
	return s.NewParentOrgId
}

func (s *MoveOrgRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *MoveOrgRequest) SetNewParentOrgId(v string) *MoveOrgRequest {
	s.NewParentOrgId = &v
	return s
}

func (s *MoveOrgRequest) SetOrgId(v string) *MoveOrgRequest {
	s.OrgId = &v
	return s
}

func (s *MoveOrgRequest) Validate() error {
	return dara.Validate(s)
}
