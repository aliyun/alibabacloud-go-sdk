// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthRole(v string) *UpdateDataMaskingUsersRequest
	GetAuthRole() *string
	SetExpireTime(v int64) *UpdateDataMaskingUsersRequest
	GetExpireTime() *int64
	SetExpireTimeOperation(v string) *UpdateDataMaskingUsersRequest
	GetExpireTimeOperation() *string
	SetLang(v string) *UpdateDataMaskingUsersRequest
	GetLang() *string
	SetProductCode(v string) *UpdateDataMaskingUsersRequest
	GetProductCode() *string
	SetProductId(v int64) *UpdateDataMaskingUsersRequest
	GetProductId() *int64
	SetUserList(v []*UpdateDataMaskingUsersRequestUserList) *UpdateDataMaskingUsersRequest
	GetUserList() []*UpdateDataMaskingUsersRequestUserList
}

type UpdateDataMaskingUsersRequest struct {
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
	ProductId *int64                                   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	UserList  []*UpdateDataMaskingUsersRequestUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s UpdateDataMaskingUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersRequest) GetAuthRole() *string {
	return s.AuthRole
}

func (s *UpdateDataMaskingUsersRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *UpdateDataMaskingUsersRequest) GetExpireTimeOperation() *string {
	return s.ExpireTimeOperation
}

func (s *UpdateDataMaskingUsersRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataMaskingUsersRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpdateDataMaskingUsersRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *UpdateDataMaskingUsersRequest) GetUserList() []*UpdateDataMaskingUsersRequestUserList {
	return s.UserList
}

func (s *UpdateDataMaskingUsersRequest) SetAuthRole(v string) *UpdateDataMaskingUsersRequest {
	s.AuthRole = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetExpireTime(v int64) *UpdateDataMaskingUsersRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetExpireTimeOperation(v string) *UpdateDataMaskingUsersRequest {
	s.ExpireTimeOperation = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetLang(v string) *UpdateDataMaskingUsersRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetProductCode(v string) *UpdateDataMaskingUsersRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetProductId(v int64) *UpdateDataMaskingUsersRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetUserList(v []*UpdateDataMaskingUsersRequestUserList) *UpdateDataMaskingUsersRequest {
	s.UserList = v
	return s
}

func (s *UpdateDataMaskingUsersRequest) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDataMaskingUsersRequestUserList struct {
	// example:
	//
	// 1001
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDataMaskingUsersRequestUserList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersRequestUserList) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersRequestUserList) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateDataMaskingUsersRequestUserList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDataMaskingUsersRequestUserList) SetAccountId(v string) *UpdateDataMaskingUsersRequestUserList {
	s.AccountId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequestUserList) SetInstanceId(v string) *UpdateDataMaskingUsersRequestUserList {
	s.InstanceId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequestUserList) Validate() error {
	return dara.Validate(s)
}
