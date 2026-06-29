// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrganizationMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *AddOrganizationMemberRequest
	GetAccountName() *string
	SetOrgRoleCode(v string) *AddOrganizationMemberRequest
	GetOrgRoleCode() *string
	SetSpecType(v string) *AddOrganizationMemberRequest
	GetSpecType() *string
}

type AddOrganizationMemberRequest struct {
	// The display name. If this parameter is not empty, the name field of the account is updated synchronously.
	//
	// This parameter is required.
	//
	// example:
	//
	// 账号名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The organization role code. Only ORG_ADMIN or ORG_MEMBER is allowed. Default value: ORG_MEMBER.
	//
	// This parameter is required.
	//
	// example:
	//
	// ORG_MEMBER
	OrgRoleCode *string `json:"OrgRoleCode,omitempty" xml:"OrgRoleCode,omitempty"`
	// The seat specification. Valid values:
	//
	// - standard: standard seat.
	//
	// - pro: premium seat.
	//
	// - max: exclusive seat.
	//
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s AddOrganizationMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOrganizationMemberRequest) GoString() string {
	return s.String()
}

func (s *AddOrganizationMemberRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *AddOrganizationMemberRequest) GetOrgRoleCode() *string {
	return s.OrgRoleCode
}

func (s *AddOrganizationMemberRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *AddOrganizationMemberRequest) SetAccountName(v string) *AddOrganizationMemberRequest {
	s.AccountName = &v
	return s
}

func (s *AddOrganizationMemberRequest) SetOrgRoleCode(v string) *AddOrganizationMemberRequest {
	s.OrgRoleCode = &v
	return s
}

func (s *AddOrganizationMemberRequest) SetSpecType(v string) *AddOrganizationMemberRequest {
	s.SpecType = &v
	return s
}

func (s *AddOrganizationMemberRequest) Validate() error {
	return dara.Validate(s)
}
