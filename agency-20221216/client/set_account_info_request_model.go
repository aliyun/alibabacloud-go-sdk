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
	// Result Code:
	//
	// 	- 200 OK
	//
	// 	- 1109 System error
	//
	// 	- 3030 Sub Account Nickname exceeds maximum length,  maximum length 150 bytes.
	//
	// 	- 3031 Remark exceeds maximum length,  maximum length 3000 bytes.
	//
	// example:
	//
	// Message information
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// Customer manager（limited 50 character）
	//
	// example:
	//
	// abc
	CustomerBd *string `json:"CustomerBd,omitempty" xml:"CustomerBd,omitempty"`
	// success
	//
	// example:
	//
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
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
