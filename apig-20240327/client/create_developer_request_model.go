// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *CreateDeveloperRequest
	GetPassword() *string
	SetUsername(v string) *CreateDeveloperRequest
	GetUsername() *string
}

type CreateDeveloperRequest struct {
	// The password of the developer account. The length must be 6 to 32 characters.
	//
	// example:
	//
	// abcd1234
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The logon username of the developer. The length must not exceed 64 characters.
	//
	// example:
	//
	// developer-demo
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeveloperRequest) GoString() string {
	return s.String()
}

func (s *CreateDeveloperRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDeveloperRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateDeveloperRequest) SetPassword(v string) *CreateDeveloperRequest {
	s.Password = &v
	return s
}

func (s *CreateDeveloperRequest) SetUsername(v string) *CreateDeveloperRequest {
	s.Username = &v
	return s
}

func (s *CreateDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
