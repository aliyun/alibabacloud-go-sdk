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
	// The Alibaba Cloud account name or Alibaba Cloud ID to query.
	//
	// - When you enter an account name:
	//
	//     - If the organization user is a primary account (for example, main_account), the query account format is the primary account. That is, enter the primary account main_account.
	//
	//     - If the organization user is a RAM user (for example, zhangsan@test.onaliyun.com), the query account format is the prefix of the RAM user. That is, enter zhangsan.
	//
	// - When you enter an Alibaba Cloud ID:
	//
	//     - Enter the complete account UID to query the corresponding account information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1386587****@163.com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The account name of the primary account when a duplicate error occurs during RAM user query. For example, zhangsan@test.onaliyun.com.
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
