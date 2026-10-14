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
	// The column encryption access permissions. Valid values:
	//
	// - fullAccess: plaintext access.
	//
	// - restrictedAccess: restricted access.
	//
	// - noneAccess: decryption denied.
	//
	// noneAccess applies only to ApsaraDB RDS for MySQL instances whose active key pattern is client_key.
	//
	// example:
	//
	// fullAccess
	AuthRole *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// The expiration time of the plaintext access permissions. Specify a UNIX timestamp in milliseconds since 1970-01-01 00:00:00 UTC. This parameter is required only when ExpireTimeOperation is set to SET. The value must be later than the current time and no later than 2145953410000. The server stores the value with second-level precision. Do not specify this parameter when ExpireTimeOperation is set to PRESERVE or CLEAR.
	//
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The action to take on the permission expiration time. Valid values:
	//
	// - PRESERVE: Retains the existing expiration time.
	//
	// - CLEAR: Clears the expiration time.
	//
	// - SET: Sets the expiration time to the value specified by ExpireTime.
	//
	// Default value: PRESERVE. Only fullAccess supports configuring a validity period.
	//
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two parameters must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type identifier. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If both ProductId and ProductCode are specified, they must correspond to each other.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of database accounts for which to set permissions.
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
