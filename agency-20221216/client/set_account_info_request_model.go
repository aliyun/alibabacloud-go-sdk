// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNickname(v string) *SetAccountInfoRequest
	GetAccountNickname() *string
	SetCustomerBd(v string) *SetAccountInfoRequest
	GetCustomerBd() *string
	SetRemark(v string) *SetAccountInfoRequest
	GetRemark() *string
	SetUid(v int64) *SetAccountInfoRequest
	GetUid() *int64
}

type SetAccountInfoRequest struct {
	// The nickname of the sub-account. For enterprise customers, this is the company name. For T2 partners, this is the partner name.
	//
	// example:
	//
	// 我的账号
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The customer business manager. Maximum length: 50 characters.
	//
	// example:
	//
	// 张三
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// The description of the sub-account.
	//
	// example:
	//
	// 测试的账号
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The UID of the sub-account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1133166938931507
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *SetAccountInfoRequest) GetAccountNickname() *string {
	return s.AccountNickname
}

func (s *SetAccountInfoRequest) GetCustomerBd() *string {
	return s.CustomerBd
}

func (s *SetAccountInfoRequest) GetRemark() *string {
	return s.Remark
}

func (s *SetAccountInfoRequest) GetUid() *int64 {
	return s.Uid
}

func (s *SetAccountInfoRequest) SetAccountNickname(v string) *SetAccountInfoRequest {
	s.AccountNickname = &v
	return s
}

func (s *SetAccountInfoRequest) SetCustomerBd(v string) *SetAccountInfoRequest {
	s.CustomerBd = &v
	return s
}

func (s *SetAccountInfoRequest) SetRemark(v string) *SetAccountInfoRequest {
	s.Remark = &v
	return s
}

func (s *SetAccountInfoRequest) SetUid(v int64) *SetAccountInfoRequest {
	s.Uid = &v
	return s
}

func (s *SetAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
