// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPasswordRequest
	GetAccountName() *string
	SetClusterId(v string) *ModifyAccountPasswordRequest
	GetClusterId() *string
	SetNewAccountPassword(v string) *ModifyAccountPasswordRequest
	GetNewAccountPassword() *string
}

type ModifyAccountPasswordRequest struct {
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
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **********
	NewAccountPassword *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPasswordRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyAccountPasswordRequest) GetNewAccountPassword() *string {
	return s.NewAccountPassword
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetClusterId(v string) *ModifyAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

func (s *ModifyAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
