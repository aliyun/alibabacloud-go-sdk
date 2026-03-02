// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDeveloperRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CheckDeveloperRoleRequest
	GetAccountId() *string
	SetCompanyId(v int64) *CheckDeveloperRoleRequest
	GetCompanyId() *int64
	SetPlatform(v string) *CheckDeveloperRoleRequest
	GetPlatform() *string
	SetRoleName(v string) *CheckDeveloperRoleRequest
	GetRoleName() *string
}

type CheckDeveloperRoleRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Platform  *string `json:"platform,omitempty" xml:"platform,omitempty"`
	RoleName  *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s CheckDeveloperRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDeveloperRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckDeveloperRoleRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CheckDeveloperRoleRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CheckDeveloperRoleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CheckDeveloperRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CheckDeveloperRoleRequest) SetAccountId(v string) *CheckDeveloperRoleRequest {
	s.AccountId = &v
	return s
}

func (s *CheckDeveloperRoleRequest) SetCompanyId(v int64) *CheckDeveloperRoleRequest {
	s.CompanyId = &v
	return s
}

func (s *CheckDeveloperRoleRequest) SetPlatform(v string) *CheckDeveloperRoleRequest {
	s.Platform = &v
	return s
}

func (s *CheckDeveloperRoleRequest) SetRoleName(v string) *CheckDeveloperRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CheckDeveloperRoleRequest) Validate() error {
	return dara.Validate(s)
}
