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
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of target instance. You can call the DescribeInstances operation to obtain target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The new password of the account. The password must meet the following requirements:
	//
	// 	- Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Is 8 to 32 characters in length.
	//
	// 	- Special characters include `!@#$%^&*()_+-=`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test*****
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
