// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *AuthorizationRequest
	GetEndpoint() *string
}

type AuthorizationRequest struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s AuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AuthorizationRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *AuthorizationRequest) SetEndpoint(v string) *AuthorizationRequest {
	s.Endpoint = &v
	return s
}

func (s *AuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
