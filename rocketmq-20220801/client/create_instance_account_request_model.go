// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *CreateInstanceAccountRequest
	GetPassword() *string
	SetUsername(v string) *CreateInstanceAccountRequest
	GetUsername() *string
}

type CreateInstanceAccountRequest struct {
	// The password of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceAccountRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetUsername(v string) *CreateInstanceAccountRequest {
	s.Username = &v
	return s
}

func (s *CreateInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}
