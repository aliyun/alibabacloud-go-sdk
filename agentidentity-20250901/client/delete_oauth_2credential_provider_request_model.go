// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOAuth2CredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOAuth2CredentialProviderName(v string) *DeleteOAuth2CredentialProviderRequest
	GetOAuth2CredentialProviderName() *string
}

type DeleteOAuth2CredentialProviderRequest struct {
	// example:
	//
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
}

func (s DeleteOAuth2CredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOAuth2CredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteOAuth2CredentialProviderRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *DeleteOAuth2CredentialProviderRequest) SetOAuth2CredentialProviderName(v string) *DeleteOAuth2CredentialProviderRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *DeleteOAuth2CredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
