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
	// Enter the name or ID of the Alibaba Cloud account that you want to query.
	//
	// 	- When you enter an account name:
	//
	//     	- If the organization user is a master account, such as main_account, the query account format is master account. That is, the main account main_account to be entered.
	//
	//     	- If the organization user is a RAM user, such as a <zhangsan@test.onaliyun.com>, the query account format is the head of the RAM user, that is, the RAM user to be entered is zhangsan.
	//
	// 	- ID:
	//
	//     	- Enter the UID of the account to query the account information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1386587****@163.com
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// When a duplicate error occurs while querying the sub-account, enter the primary account\\"s username, for example, zhangsan@test.onaliyun.com.
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
