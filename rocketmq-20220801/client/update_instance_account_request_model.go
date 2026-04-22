// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountStatus(v string) *UpdateInstanceAccountRequest
	GetAccountStatus() *string
	SetPassword(v string) *UpdateInstanceAccountRequest
	GetPassword() *string
	SetRemark(v string) *UpdateInstanceAccountRequest
	GetRemark() *string
}

type UpdateInstanceAccountRequest struct {
	// The status of the account.
	//
	// Valid values:
	//
	// 	- DISABLE
	//
	// 	- ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The password of the account.
	//
	// example:
	//
	// test
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountRequest) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *UpdateInstanceAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateInstanceAccountRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateInstanceAccountRequest) SetAccountStatus(v string) *UpdateInstanceAccountRequest {
	s.AccountStatus = &v
	return s
}

func (s *UpdateInstanceAccountRequest) SetPassword(v string) *UpdateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *UpdateInstanceAccountRequest) SetRemark(v string) *UpdateInstanceAccountRequest {
	s.Remark = &v
	return s
}

func (s *UpdateInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}
