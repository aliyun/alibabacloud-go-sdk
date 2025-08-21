// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ValidateConnectionRequest
	GetClientToken() *string
	SetBody(v string) *ValidateConnectionRequest
	GetBody() *string
}

type ValidateConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateConnectionRequest) GoString() string {
	return s.String()
}

func (s *ValidateConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateConnectionRequest) GetBody() *string {
	return s.Body
}

func (s *ValidateConnectionRequest) SetClientToken(v string) *ValidateConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateConnectionRequest) SetBody(v string) *ValidateConnectionRequest {
	s.Body = &v
	return s
}

func (s *ValidateConnectionRequest) Validate() error {
	return dara.Validate(s)
}
