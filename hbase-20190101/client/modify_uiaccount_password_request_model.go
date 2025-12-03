// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyUIAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ModifyUIAccountPasswordRequest
	GetAccountPassword() *string
	SetClusterId(v string) *ModifyUIAccountPasswordRequest
	GetClusterId() *string
}

type ModifyUIAccountPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **********
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ModifyUIAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyUIAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ModifyUIAccountPasswordRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyUIAccountPasswordRequest) SetAccountName(v string) *ModifyUIAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) SetClusterId(v string) *ModifyUIAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUIAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
