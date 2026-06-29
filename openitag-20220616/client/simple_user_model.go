// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleUser interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v string) *SimpleUser
	GetAccountNo() *string
	SetAccountType(v string) *SimpleUser
	GetAccountType() *string
	SetRole(v string) *SimpleUser
	GetRole() *string
	SetUserId(v int64) *SimpleUser
	GetUserId() *int64
	SetUserName(v string) *SimpleUser
	GetUserName() *string
}

type SimpleUser struct {
	// Account ID
	//
	// example:
	//
	// 166***9980757310
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// Account Type
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Role
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// User ID
	//
	// example:
	//
	// 166***9980757311
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Username
	//
	// example:
	//
	// txdemo@test.aliyunid.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SimpleUser) String() string {
	return dara.Prettify(s)
}

func (s SimpleUser) GoString() string {
	return s.String()
}

func (s *SimpleUser) GetAccountNo() *string {
	return s.AccountNo
}

func (s *SimpleUser) GetAccountType() *string {
	return s.AccountType
}

func (s *SimpleUser) GetRole() *string {
	return s.Role
}

func (s *SimpleUser) GetUserId() *int64 {
	return s.UserId
}

func (s *SimpleUser) GetUserName() *string {
	return s.UserName
}

func (s *SimpleUser) SetAccountNo(v string) *SimpleUser {
	s.AccountNo = &v
	return s
}

func (s *SimpleUser) SetAccountType(v string) *SimpleUser {
	s.AccountType = &v
	return s
}

func (s *SimpleUser) SetRole(v string) *SimpleUser {
	s.Role = &v
	return s
}

func (s *SimpleUser) SetUserId(v int64) *SimpleUser {
	s.UserId = &v
	return s
}

func (s *SimpleUser) SetUserName(v string) *SimpleUser {
	s.UserName = &v
	return s
}

func (s *SimpleUser) Validate() error {
	return dara.Validate(s)
}
