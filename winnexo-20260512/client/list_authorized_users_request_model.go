// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGranteeType(v string) *ListAuthorizedUsersRequest
	GetGranteeType() *string
	SetKeyword(v string) *ListAuthorizedUsersRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListAuthorizedUsersRequest
	GetOperatingObjectName() *string
	SetPermission(v string) *ListAuthorizedUsersRequest
	GetPermission() *string
	SetTenantId(v string) *ListAuthorizedUsersRequest
	GetTenantId() *string
}

type ListAuthorizedUsersRequest struct {
	// The filter type. Valid values: USER, USER_GROUP. If not specified, all types are returned.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// Permission
	//
	// example:
	//
	// USE
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAuthorizedUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersRequest) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAuthorizedUsersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAuthorizedUsersRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListAuthorizedUsersRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListAuthorizedUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAuthorizedUsersRequest) SetGranteeType(v string) *ListAuthorizedUsersRequest {
	s.GranteeType = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetKeyword(v string) *ListAuthorizedUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetOperatingObjectName(v string) *ListAuthorizedUsersRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPermission(v string) *ListAuthorizedUsersRequest {
	s.Permission = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetTenantId(v string) *ListAuthorizedUsersRequest {
	s.TenantId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) Validate() error {
	return dara.Validate(s)
}
