// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v string) *CreateUserRequest
	GetAccountNo() *string
	SetAccountType(v string) *CreateUserRequest
	GetAccountType() *string
	SetRole(v string) *CreateUserRequest
	GetRole() *string
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
}

type CreateUserRequest struct {
	// UID of the RAM user (sub-account) under the current Alibaba Cloud account (primary account). For information about how to obtain the UID, see [GetUser](https://help.aliyun.com/document_detail/2330970.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 166***980757310
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// Account type. Only ALIYUN is currently supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Role. Valid values:
	//
	// - OPERATOR: Annotator.
	//
	// - ADMIN: Administrator.
	//
	// - LEADER: Annotation team leader.
	//
	// This parameter is required.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Username.
	//
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetAccountNo() *string {
	return s.AccountNo
}

func (s *CreateUserRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateUserRequest) GetRole() *string {
	return s.Role
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) SetAccountNo(v string) *CreateUserRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateUserRequest) SetAccountType(v string) *CreateUserRequest {
	s.AccountType = &v
	return s
}

func (s *CreateUserRequest) SetRole(v string) *CreateUserRequest {
	s.Role = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
