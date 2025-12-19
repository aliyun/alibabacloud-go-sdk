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
	// example:
	//
	// 0B447F15-7037-512D-8EFC-A4188FC9E9E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default/apikeycredentialprovider/api-key-dash-scope
	CredentialProviderArn *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
	// example:
	//
	// example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s *CreateAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) Validate() error {
	return dara.Validate(s)
}
