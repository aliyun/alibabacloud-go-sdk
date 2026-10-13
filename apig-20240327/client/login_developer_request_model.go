// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *LoginDeveloperRequest
	GetPassword() *string
	SetUsername(v string) *LoginDeveloperRequest
	GetUsername() *string
}

type LoginDeveloperRequest struct {
	// The password of the developer account. This field is marked as sensitive in the API definition.
	//
	// example:
	//
	// abcd1234
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The developer logon username.
	//
	// example:
	//
	// developer-demo
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s LoginDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginDeveloperRequest) GoString() string {
	return s.String()
}

func (s *LoginDeveloperRequest) GetPassword() *string {
	return s.Password
}

func (s *LoginDeveloperRequest) GetUsername() *string {
	return s.Username
}

func (s *LoginDeveloperRequest) SetPassword(v string) *LoginDeveloperRequest {
	s.Password = &v
	return s
}

func (s *LoginDeveloperRequest) SetUsername(v string) *LoginDeveloperRequest {
	s.Username = &v
	return s
}

func (s *LoginDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
