// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *UpdateOrganizationMemberRequest
	GetAccountIds() []*string
	SetNewRoleCode(v string) *UpdateOrganizationMemberRequest
	GetNewRoleCode() *string
}

type UpdateOrganizationMemberRequest struct {
	// The list of account IDs for batch operations.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The code of the new role. When you perform a batch operation, all specified accounts are changed to this role.
	//
	// This parameter is required.
	//
	// example:
	//
	// ORG_MEMBER
	NewRoleCode *string `json:"NewRoleCode,omitempty" xml:"NewRoleCode,omitempty"`
}

func (s UpdateOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *UpdateOrganizationMemberRequest) GetNewRoleCode() *string {
	return s.NewRoleCode
}

func (s *UpdateOrganizationMemberRequest) SetAccountIds(v []*string) *UpdateOrganizationMemberRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateOrganizationMemberRequest) SetNewRoleCode(v string) *UpdateOrganizationMemberRequest {
	s.NewRoleCode = &v
	return s
}

func (s *UpdateOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
