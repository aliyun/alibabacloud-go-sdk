// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthRole(v string) *UpdateDataMaskingUsersShrinkRequest
	GetAuthRole() *string
	SetExpireTime(v int64) *UpdateDataMaskingUsersShrinkRequest
	GetExpireTime() *int64
	SetExpireTimeOperation(v string) *UpdateDataMaskingUsersShrinkRequest
	GetExpireTimeOperation() *string
	SetLang(v string) *UpdateDataMaskingUsersShrinkRequest
	GetLang() *string
	SetProductCode(v string) *UpdateDataMaskingUsersShrinkRequest
	GetProductCode() *string
	SetProductId(v int64) *UpdateDataMaskingUsersShrinkRequest
	GetProductId() *int64
	SetUserListShrink(v string) *UpdateDataMaskingUsersShrinkRequest
	GetUserListShrink() *string
}

type UpdateDataMaskingUsersShrinkRequest struct {
	// example:
	//
	// fullAccess
	AuthRole *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId      *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	UserListShrink *string `json:"UserList,omitempty" xml:"UserList,omitempty"`
}

func (s UpdateDataMaskingUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetAuthRole() *string {
	return s.AuthRole
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetExpireTimeOperation() *string {
	return s.ExpireTimeOperation
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *UpdateDataMaskingUsersShrinkRequest) GetUserListShrink() *string {
	return s.UserListShrink
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetAuthRole(v string) *UpdateDataMaskingUsersShrinkRequest {
	s.AuthRole = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetExpireTime(v int64) *UpdateDataMaskingUsersShrinkRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetExpireTimeOperation(v string) *UpdateDataMaskingUsersShrinkRequest {
	s.ExpireTimeOperation = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetLang(v string) *UpdateDataMaskingUsersShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetProductCode(v string) *UpdateDataMaskingUsersShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetProductId(v int64) *UpdateDataMaskingUsersShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) SetUserListShrink(v string) *UpdateDataMaskingUsersShrinkRequest {
	s.UserListShrink = &v
	return s
}

func (s *UpdateDataMaskingUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
