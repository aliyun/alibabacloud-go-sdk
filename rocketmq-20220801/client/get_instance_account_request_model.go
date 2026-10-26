// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUsername(v string) *GetInstanceAccountRequest
	GetUsername() *string
}

type GetInstanceAccountRequest struct {
	// The username.
	//
	// If you do not specify this parameter, the system returns details of the default account.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountRequest) GetUsername() *string {
	return s.Username
}

func (s *GetInstanceAccountRequest) SetUsername(v string) *GetInstanceAccountRequest {
	s.Username = &v
	return s
}

func (s *GetInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}
