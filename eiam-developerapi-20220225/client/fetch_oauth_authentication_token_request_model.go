// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchOAuthAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderIdentifier(v string) *FetchOAuthAuthenticationTokenRequest
	GetCredentialProviderIdentifier() *string
	SetScope(v string) *FetchOAuthAuthenticationTokenRequest
	GetScope() *string
}

type FetchOAuthAuthenticationTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s FetchOAuthAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenRequest) GetCredentialProviderIdentifier() *string {
	return s.CredentialProviderIdentifier
}

func (s *FetchOAuthAuthenticationTokenRequest) GetScope() *string {
	return s.Scope
}

func (s *FetchOAuthAuthenticationTokenRequest) SetCredentialProviderIdentifier(v string) *FetchOAuthAuthenticationTokenRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) SetScope(v string) *FetchOAuthAuthenticationTokenRequest {
	s.Scope = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
