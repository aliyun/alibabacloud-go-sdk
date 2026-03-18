// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAuthenticationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ValidateAuthenticationTokenResponseBody
	GetActive() *bool
}

type ValidateAuthenticationTokenResponseBody struct {
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s ValidateAuthenticationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateAuthenticationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateAuthenticationTokenResponseBody) GetActive() *bool {
	return s.Active
}

func (s *ValidateAuthenticationTokenResponseBody) SetActive(v bool) *ValidateAuthenticationTokenResponseBody {
	s.Active = &v
	return s
}

func (s *ValidateAuthenticationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
