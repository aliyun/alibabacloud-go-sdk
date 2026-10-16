// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAPIKeyCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProviderName(v string) *GetAPIKeyCredentialProviderRequest
	GetAPIKeyCredentialProviderName() *string
	SetTokenVaultName(v string) *GetAPIKeyCredentialProviderRequest
	GetTokenVaultName() *string
}

type GetAPIKeyCredentialProviderRequest struct {
	// The name of the credential provider to retrieve.
	//
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	// The name of the TokenVault.
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s GetAPIKeyCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAPIKeyCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *GetAPIKeyCredentialProviderRequest) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *GetAPIKeyCredentialProviderRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *GetAPIKeyCredentialProviderRequest) SetAPIKeyCredentialProviderName(v string) *GetAPIKeyCredentialProviderRequest {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *GetAPIKeyCredentialProviderRequest) SetTokenVaultName(v string) *GetAPIKeyCredentialProviderRequest {
	s.TokenVaultName = &v
	return s
}

func (s *GetAPIKeyCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
