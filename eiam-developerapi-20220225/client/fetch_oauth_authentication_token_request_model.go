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
	SetCustomParameters(v map[string]*string) *FetchOAuthAuthenticationTokenRequest
	GetCustomParameters() map[string]*string
	SetForceAuthentication(v bool) *FetchOAuthAuthenticationTokenRequest
	GetForceAuthentication() *bool
	SetScope(v string) *FetchOAuthAuthenticationTokenRequest
	GetScope() *string
}

type FetchOAuthAuthenticationTokenRequest struct {
	// The credential provider identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_example_identifier
	CredentialProviderIdentifier *string `json:"credentialProviderIdentifier,omitempty" xml:"credentialProviderIdentifier,omitempty"`
	// Custom key-value pairs appended to the OAuth authorization URL to pass additional parameters supported by the OAuth provider.
	CustomParameters map[string]*string `json:"customParameters,omitempty" xml:"customParameters,omitempty"`
	// Specifies whether to ignore existing valid tokens and force re-authorization. Default value: false.
	//
	// example:
	//
	// false
	ForceAuthentication *bool `json:"forceAuthentication,omitempty" xml:"forceAuthentication,omitempty"`
	// The scope corresponding to the OAuth protocol.
	//
	// > If not specified, the scope of the issued OAuth Access Token defaults to the scope configuration of the corresponding credential provider.
	//
	// 	Notice: Multiple scope values are separated by spaces.
	//
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

func (s *FetchOAuthAuthenticationTokenRequest) GetCustomParameters() map[string]*string {
	return s.CustomParameters
}

func (s *FetchOAuthAuthenticationTokenRequest) GetForceAuthentication() *bool {
	return s.ForceAuthentication
}

func (s *FetchOAuthAuthenticationTokenRequest) GetScope() *string {
	return s.Scope
}

func (s *FetchOAuthAuthenticationTokenRequest) SetCredentialProviderIdentifier(v string) *FetchOAuthAuthenticationTokenRequest {
	s.CredentialProviderIdentifier = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) SetCustomParameters(v map[string]*string) *FetchOAuthAuthenticationTokenRequest {
	s.CustomParameters = v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) SetForceAuthentication(v bool) *FetchOAuthAuthenticationTokenRequest {
	s.ForceAuthentication = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) SetScope(v string) *FetchOAuthAuthenticationTokenRequest {
	s.Scope = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
