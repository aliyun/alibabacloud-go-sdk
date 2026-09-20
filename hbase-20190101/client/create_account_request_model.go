// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetClusterId(v string) *CreateAccountRequest
	GetClusterId() *string
}

type CreateAccountRequest struct {
	// The account name. The name must meet the following requirements:
	//
	// 	- Starts with a lowercase letter and ends with a letter or digit.
	//
	// 	- Contains only lowercase letters, digits, or underscores.
	//
	// 	- Is 2 to 16 characters in length.
	//
	// 	- Cannot be a reserved username such as root or admin.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account. The password must meet the following requirements:
	//
	// - Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The supported special characters are `!@#$%^&*()_+-=`.
	//
	// - Is 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test*****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetClusterId(v string) *CreateAccountRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
