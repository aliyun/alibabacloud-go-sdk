// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuth2CredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOAuth2CredentialProviderName(v string) *GetOAuth2CredentialProviderRequest
	GetOAuth2CredentialProviderName() *string
}

type GetOAuth2CredentialProviderRequest struct {
	// example:
	//
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
}

func (s GetOAuth2CredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOAuth2CredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *GetOAuth2CredentialProviderRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *GetOAuth2CredentialProviderRequest) SetOAuth2CredentialProviderName(v string) *GetOAuth2CredentialProviderRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *GetOAuth2CredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
