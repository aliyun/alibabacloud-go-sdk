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

func (s *UpdateInstanceAccountRequest) SetAccountStatus(v string) *UpdateInstanceAccountRequest {
	s.AccountStatus = &v
	return s
}

func (s *UpdateInstanceAccountRequest) SetPassword(v string) *UpdateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *UpdateInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}
