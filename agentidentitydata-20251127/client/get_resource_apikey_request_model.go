// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceAPIKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCredentialProviderName(v string) *GetResourceAPIKeyRequest
	GetResourceCredentialProviderName() *string
	SetTokenVaultName(v string) *GetResourceAPIKeyRequest
	GetTokenVaultName() *string
	SetWorkloadAccessToken(v string) *GetResourceAPIKeyRequest
	GetWorkloadAccessToken() *string
}

type GetResourceAPIKeyRequest struct {
	ResourceCredentialProviderName *string `json:"ResourceCredentialProviderName,omitempty" xml:"ResourceCredentialProviderName,omitempty"`
	TokenVaultName                 *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	WorkloadAccessToken            *string `json:"WorkloadAccessToken,omitempty" xml:"WorkloadAccessToken,omitempty"`
}

func (s GetResourceAPIKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceAPIKeyRequest) GoString() string {
	return s.String()
}

func (s *GetResourceAPIKeyRequest) GetResourceCredentialProviderName() *string {
	return s.ResourceCredentialProviderName
}

func (s *GetResourceAPIKeyRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *GetResourceAPIKeyRequest) GetWorkloadAccessToken() *string {
	return s.WorkloadAccessToken
}

func (s *GetResourceAPIKeyRequest) SetResourceCredentialProviderName(v string) *GetResourceAPIKeyRequest {
	s.ResourceCredentialProviderName = &v
	return s
}

func (s *GetResourceAPIKeyRequest) SetTokenVaultName(v string) *GetResourceAPIKeyRequest {
	s.TokenVaultName = &v
	return s
}

func (s *GetResourceAPIKeyRequest) SetWorkloadAccessToken(v string) *GetResourceAPIKeyRequest {
	s.WorkloadAccessToken = &v
	return s
}

func (s *GetResourceAPIKeyRequest) Validate() error {
	return dara.Validate(s)
}
