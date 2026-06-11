// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *QueryUserInfoByAccountRequest
	GetAccount() *string
	SetParentAccountName(v string) *QueryUserInfoByAccountRequest
	GetParentAccountName() *string
}

type QueryUserInfoByAccountRequest struct {
	// The Alibaba Cloud account name or Alibaba Cloud ID of the user.
	//
	// - If you enter an account name:
	//
	//   - If the organization member is a root account, such as `main_account`, enter the root account name. For example, `main_account`.
	//
	//   - If the organization member is a RAM user, such as `zhangsan@test.onaliyun.com`, enter the prefix of the username before the at sign (@). For example, `zhangsan`.
	//
	// - If you enter an Alibaba Cloud ID:
	//
	//   - Enter the complete user ID (UID) of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1386587****@163.com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// To resolve a "duplicate user" error when querying a RAM user, specify the name of the root account to which the user belongs.
	//
	// example:
	//
	// zhangsan@test.onaliyun.com
	ParentAccountName *string `json:"ParentAccountName,omitempty" xml:"ParentAccountName,omitempty"`
}

func (s QueryUserInfoByAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountRequest) GetAccount() *string {
	return s.Account
}

func (s *QueryUserInfoByAccountRequest) GetParentAccountName() *string {
	return s.ParentAccountName
}

func (s *QueryUserInfoByAccountRequest) SetAccount(v string) *QueryUserInfoByAccountRequest {
	s.Account = &v
	return s
}

func (s *QueryUserInfoByAccountRequest) SetParentAccountName(v string) *QueryUserInfoByAccountRequest {
	s.ParentAccountName = &v
	return s
}

func (s *QueryUserInfoByAccountRequest) Validate() error {
	return dara.Validate(s)
}
