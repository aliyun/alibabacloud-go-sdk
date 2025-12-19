// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOAuth2TokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomParametersShrink(v string) *GetResourceOAuth2TokenShrinkRequest
	GetCustomParametersShrink() *string
	SetCustomState(v string) *GetResourceOAuth2TokenShrinkRequest
	GetCustomState() *string
	SetForceAuthentication(v bool) *GetResourceOAuth2TokenShrinkRequest
	GetForceAuthentication() *bool
	SetOAuth2Flow(v string) *GetResourceOAuth2TokenShrinkRequest
	GetOAuth2Flow() *string
	SetResourceCredentialProviderName(v string) *GetResourceOAuth2TokenShrinkRequest
	GetResourceCredentialProviderName() *string
	SetResourceOAuth2ReturnURL(v string) *GetResourceOAuth2TokenShrinkRequest
	GetResourceOAuth2ReturnURL() *string
	SetScopesShrink(v string) *GetResourceOAuth2TokenShrinkRequest
	GetScopesShrink() *string
	SetSessionURI(v string) *GetResourceOAuth2TokenShrinkRequest
	GetSessionURI() *string
	SetWorkloadAccessToken(v string) *GetResourceOAuth2TokenShrinkRequest
	GetWorkloadAccessToken() *string
}

type GetResourceOAuth2TokenShrinkRequest struct {
	// example:
	//
	// {"param1": "test-param", "param2": "test-param2"}
	CustomParametersShrink *string `json:"CustomParameters,omitempty" xml:"CustomParameters,omitempty"`
	// example:
	//
	// user-defined-value
	CustomState *string `json:"CustomState,omitempty" xml:"CustomState,omitempty"`
	// example:
	//
	// False
	ForceAuthentication *bool `json:"ForceAuthentication,omitempty" xml:"ForceAuthentication,omitempty"`
	// example:
	//
	// USER_FEDERATION
	OAuth2Flow *string `json:"OAuth2Flow,omitempty" xml:"OAuth2Flow,omitempty"`
	// example:
	//
	// oauth_credential_provider
	ResourceCredentialProviderName *string `json:"ResourceCredentialProviderName,omitempty" xml:"ResourceCredentialProviderName,omitempty"`
	// example:
	//
	// https://example.com
	ResourceOAuth2ReturnURL *string `json:"ResourceOAuth2ReturnURL,omitempty" xml:"ResourceOAuth2ReturnURL,omitempty"`
	ScopesShrink            *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
	// example:
	//
	// urn:ietf:params:oauth:request_uri:43b7df1a-****-****-****-709375a44d8a
	SessionURI *string `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1N****
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s GetResourceOAuth2TokenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOAuth2TokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetCustomParametersShrink() *string {
	return s.CustomParametersShrink
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetCustomState() *string {
	return s.CustomState
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetForceAuthentication() *bool {
	return s.ForceAuthentication
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetOAuth2Flow() *string {
	return s.OAuth2Flow
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetResourceCredentialProviderName() *string {
	return s.ResourceCredentialProviderName
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetResourceOAuth2ReturnURL() *string {
	return s.ResourceOAuth2ReturnURL
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetScopesShrink() *string {
	return s.ScopesShrink
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetSessionURI() *string {
	return s.SessionURI
}

func (s *GetResourceOAuth2TokenShrinkRequest) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetCustomParametersShrink(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.CustomParametersShrink = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetCustomState(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.CustomState = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetForceAuthentication(v bool) *GetResourceOAuth2TokenShrinkRequest {
	s.ForceAuthentication = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetOAuth2Flow(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.OAuth2Flow = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetResourceCredentialProviderName(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.ResourceCredentialProviderName = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetResourceOAuth2ReturnURL(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.ResourceOAuth2ReturnURL = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetScopesShrink(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.ScopesShrink = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetSessionURI(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.SessionURI = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) SetWorkloadAccessToken(v string) *GetResourceOAuth2TokenShrinkRequest {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetResourceOAuth2TokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
