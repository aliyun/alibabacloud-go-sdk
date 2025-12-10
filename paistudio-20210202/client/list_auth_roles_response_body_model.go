// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAuthRolesResponseBody
	GetRequestId() *string
	SetRoles(v []*ListAuthRolesResponseBodyRoles) *ListAuthRolesResponseBody
	GetRoles() []*ListAuthRolesResponseBodyRoles
}

type ListAuthRolesResponseBody struct {
	// example:
	//
	// 8E2C1BB9-57C4-5051-9EF2-570ADC03A164
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     []*ListAuthRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s ListAuthRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthRolesResponseBody) GetRoles() []*ListAuthRolesResponseBodyRoles {
	return s.Roles
}

func (s *ListAuthRolesResponseBody) SetRequestId(v string) *ListAuthRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthRolesResponseBody) SetRoles(v []*ListAuthRolesResponseBodyRoles) *ListAuthRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListAuthRolesResponseBody) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthRolesResponseBodyRoles struct {
	// example:
	//
	// true
	IsEnabled *string `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	// example:
	//
	// acs:ram::1557********904:role/aliyunodpspaidefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// example:
	//
	// AliyunODPSPAIDefaultRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// PaiStudio。
	RoleType *string                              `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	Token    *ListAuthRolesResponseBodyRolesToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s ListAuthRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListAuthRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBodyRoles) GetIsEnabled() *string {
	return s.IsEnabled
}

func (s *ListAuthRolesResponseBodyRoles) GetRoleARN() *string {
	return s.RoleARN
}

func (s *ListAuthRolesResponseBodyRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *ListAuthRolesResponseBodyRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *ListAuthRolesResponseBodyRoles) GetToken() *ListAuthRolesResponseBodyRolesToken {
	return s.Token
}

func (s *ListAuthRolesResponseBodyRoles) SetIsEnabled(v string) *ListAuthRolesResponseBodyRoles {
	s.IsEnabled = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleARN(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleARN = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleName(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleName = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetRoleType(v string) *ListAuthRolesResponseBodyRoles {
	s.RoleType = &v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) SetToken(v *ListAuthRolesResponseBodyRolesToken) *ListAuthRolesResponseBodyRoles {
	s.Token = v
	return s
}

func (s *ListAuthRolesResponseBodyRoles) Validate() error {
	if s.Token != nil {
		if err := s.Token.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthRolesResponseBodyRolesToken struct {
	// example:
	//
	// STS.NU************TT5LoC
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// E1h2n66Duo1D**********c79JVk59R6i
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// 2021-03-19T19:14:42Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// CAISggJ1q6Ft5B2yf***************aAaDf+bmceH2MNtNe9XtmTXJytadQ2T0RT8uOA+4kSypOPxSHjdjmnQjbdA/Q9MyNtTErQ/m45RNsg==
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAuthRolesResponseBodyRolesToken) String() string {
	return dara.Prettify(s)
}

func (s ListAuthRolesResponseBodyRolesToken) GoString() string {
	return s.String()
}

func (s *ListAuthRolesResponseBodyRolesToken) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ListAuthRolesResponseBodyRolesToken) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *ListAuthRolesResponseBodyRolesToken) GetExpiration() *string {
	return s.Expiration
}

func (s *ListAuthRolesResponseBodyRolesToken) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAuthRolesResponseBodyRolesToken) SetAccessKeyId(v string) *ListAuthRolesResponseBodyRolesToken {
	s.AccessKeyId = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetAccessKeySecret(v string) *ListAuthRolesResponseBodyRolesToken {
	s.AccessKeySecret = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetExpiration(v string) *ListAuthRolesResponseBodyRolesToken {
	s.Expiration = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) SetSecurityToken(v string) *ListAuthRolesResponseBodyRolesToken {
	s.SecurityToken = &v
	return s
}

func (s *ListAuthRolesResponseBodyRolesToken) Validate() error {
	return dara.Validate(s)
}
