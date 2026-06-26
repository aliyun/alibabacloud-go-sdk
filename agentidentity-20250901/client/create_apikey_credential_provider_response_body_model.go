// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPIKeyCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProvider(v *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) *CreateAPIKeyCredentialProviderResponseBody
	GetAPIKeyCredentialProvider() *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider
	SetRequestId(v string) *CreateAPIKeyCredentialProviderResponseBody
	GetRequestId() *string
}

type CreateAPIKeyCredentialProviderResponseBody struct {
	APIKeyCredentialProvider *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider `json:"APIKeyCredentialProvider,omitempty" xml:"APIKeyCredentialProvider,omitempty" type:"Struct"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAPIKeyCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAPIKeyCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAPIKeyCredentialProviderResponseBody) GetAPIKeyCredentialProvider() *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	return s.APIKeyCredentialProvider
}

func (s *CreateAPIKeyCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAPIKeyCredentialProviderResponseBody) SetAPIKeyCredentialProvider(v *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) *CreateAPIKeyCredentialProviderResponseBody {
	s.APIKeyCredentialProvider = v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBody) SetRequestId(v string) *CreateAPIKeyCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBody) Validate() error {
	if s.APIKeyCredentialProvider != nil {
		if err := s.APIKeyCredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider struct {
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	CreateTime                   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialProviderArn        *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TokenVaultName               *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GoString() string {
	return s.String()
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetAPIKeyCredentialProviderName(v string) *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetCreateTime(v string) *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetCredentialProviderArn(v string) *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.CredentialProviderArn = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetDescription(v string) *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.Description = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetTokenVaultName(v string) *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.TokenVaultName = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) Validate() error {
	return dara.Validate(s)
}
