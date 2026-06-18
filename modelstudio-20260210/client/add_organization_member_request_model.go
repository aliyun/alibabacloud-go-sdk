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
	SetCallerUacAccountId(v string) *AddOrganizationMemberRequest
	GetCallerUacAccountId() *string
	SetNamespaceId(v string) *AddOrganizationMemberRequest
	GetNamespaceId() *string
	SetOrgId(v string) *AddOrganizationMemberRequest
	GetOrgId() *string
	SetOrgRoleCode(v string) *AddOrganizationMemberRequest
	GetOrgRoleCode() *string
	SetSpecType(v string) *AddOrganizationMemberRequest
	GetSpecType() *string
}

type AddOrganizationMemberRequest struct {
	// The display name. If this parameter is not empty, the name field of the account is also updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 账号名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account information of the TokenPlan management platform.
	//
	// example:
	//
	// acc_123456789
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// The product namespace ID. For the TokenPlan product, this field is fixed to namespace-1.
	//
	// example:
	//
	// namespace-1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// org_123456789
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization role code. Valid values: ORG_ADMIN and ORG_MEMBER. Default value: ORG_MEMBER.
	//
	// This parameter is required.
	//
	// example:
	//
	// ORG_MEMBER
	OrgRoleCode *string `json:"OrgRoleCode,omitempty" xml:"OrgRoleCode,omitempty"`
	// The seat specification.
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

func (s *AddOrganizationMemberRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *AddOrganizationMemberRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *AddOrganizationMemberRequest) GetOrgId() *string {
	return s.OrgId
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

func (s *AddOrganizationMemberRequest) SetCallerUacAccountId(v string) *AddOrganizationMemberRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *AddOrganizationMemberRequest) SetNamespaceId(v string) *AddOrganizationMemberRequest {
	s.NamespaceId = &v
	return s
}

func (s *AddOrganizationMemberRequest) SetOrgId(v string) *AddOrganizationMemberRequest {
	s.OrgId = &v
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
