// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateOidcAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvider(v string) *InitiateOidcAuthRequest
	GetProvider() *string
}

type InitiateOidcAuthRequest struct {
	// The OIDC provider identifier.
	//
	// example:
	//
	// provider-xxxxxxxx
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s InitiateOidcAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s InitiateOidcAuthRequest) GoString() string {
	return s.String()
}

func (s *InitiateOidcAuthRequest) GetProvider() *string {
	return s.Provider
}

func (s *InitiateOidcAuthRequest) SetProvider(v string) *InitiateOidcAuthRequest {
	s.Provider = &v
	return s
}

func (s *InitiateOidcAuthRequest) Validate() error {
	return dara.Validate(s)
}
