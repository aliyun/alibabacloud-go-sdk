// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOAuth2TokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomParameters(v map[string]interface{}) *GetResourceOAuth2TokenRequest
	GetCustomParameters() map[string]interface{}
	SetCustomState(v string) *GetResourceOAuth2TokenRequest
	GetCustomState() *string
	SetForceAuthentication(v bool) *GetResourceOAuth2TokenRequest
	GetForceAuthentication() *bool
	SetOAuth2Flow(v string) *GetResourceOAuth2TokenRequest
	GetOAuth2Flow() *string
	SetResourceCredentialProviderName(v string) *GetResourceOAuth2TokenRequest
	GetResourceCredentialProviderName() *string
	SetResourceOAuth2ReturnURL(v string) *GetResourceOAuth2TokenRequest
	GetResourceOAuth2ReturnURL() *string
	SetScopes(v []*string) *GetResourceOAuth2TokenRequest
	GetScopes() []*string
	SetSessionURI(v string) *GetResourceOAuth2TokenRequest
	GetSessionURI() *string
	SetWorkloadAccessToken(v string) *GetResourceOAuth2TokenRequest
	GetWorkloadAccessToken() *string
}

type GetResourceOAuth2TokenRequest struct {
	// example:
	//
	// {"param1": "test-param", "param2": "test-param2"}
	CustomParameters map[string]interface{} `json:"CustomParameters,omitempty" xml:"CustomParameters,omitempty"`
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
	ResourceOAuth2ReturnURL *string   `json:"ResourceOAuth2ReturnURL,omitempty" xml:"ResourceOAuth2ReturnURL,omitempty"`
	Scopes                  []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// example:
	//
	// urn:ietf:params:oauth:request_uri:43b7df1a-****-****-****-709375a44d8a
	SessionURI *string `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
	// example:
	//
	// eyAgImFsZyI6ICJSUzI1N****
	WorkloadAccessToken *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s GetResourceOAuth2TokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOAuth2TokenRequest) GoString() string {
	return s.String()
}

func (s *GetResourceOAuth2TokenRequest) GetCustomParameters() map[string]interface{} {
	return s.CustomParameters
}

func (s *GetResourceOAuth2TokenRequest) GetCustomState() *string {
	return s.CustomState
}

func (s *GetResourceOAuth2TokenRequest) GetForceAuthentication() *bool {
	return s.ForceAuthentication
}

func (s *GetResourceOAuth2TokenRequest) GetOAuth2Flow() *string {
	return s.OAuth2Flow
}

func (s *GetResourceOAuth2TokenRequest) GetResourceCredentialProviderName() *string {
	return s.ResourceCredentialProviderName
}

func (s *GetResourceOAuth2TokenRequest) GetResourceOAuth2ReturnURL() *string {
	return s.ResourceOAuth2ReturnURL
}

func (s *GetResourceOAuth2TokenRequest) GetScopes() []*string {
	return s.Scopes
}

func (s *GetResourceOAuth2TokenRequest) GetSessionURI() *string {
	return s.SessionURI
}

func (s *GetResourceOAuth2TokenRequest) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetResourceOAuth2TokenRequest) SetCustomParameters(v map[string]interface{}) *GetResourceOAuth2TokenRequest {
	s.CustomParameters = v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetCustomState(v string) *GetResourceOAuth2TokenRequest {
	s.CustomState = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetForceAuthentication(v bool) *GetResourceOAuth2TokenRequest {
	s.ForceAuthentication = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetOAuth2Flow(v string) *GetResourceOAuth2TokenRequest {
	s.OAuth2Flow = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetResourceCredentialProviderName(v string) *GetResourceOAuth2TokenRequest {
	s.ResourceCredentialProviderName = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetResourceOAuth2ReturnURL(v string) *GetResourceOAuth2TokenRequest {
	s.ResourceOAuth2ReturnURL = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetScopes(v []*string) *GetResourceOAuth2TokenRequest {
	s.Scopes = v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetSessionURI(v string) *GetResourceOAuth2TokenRequest {
	s.SessionURI = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) SetWorkloadAccessToken(v string) *GetResourceOAuth2TokenRequest {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetResourceOAuth2TokenRequest) Validate() error {
	return dara.Validate(s)
}
