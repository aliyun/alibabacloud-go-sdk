// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAPIKeyCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProvider(v *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) *GetAPIKeyCredentialProviderResponseBody
	GetAPIKeyCredentialProvider() *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider
	SetRequestId(v string) *GetAPIKeyCredentialProviderResponseBody
	GetRequestId() *string
}

type GetAPIKeyCredentialProviderResponseBody struct {
	APIKeyCredentialProvider *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider `json:"APIKeyCredentialProvider,omitempty" xml:"APIKeyCredentialProvider,omitempty" type:"Struct"`
	// example:
	//
	// 76BF2570-76F5-5093-87CE-0918DD85752C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAPIKeyCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAPIKeyCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetAPIKeyCredentialProviderResponseBody) GetAPIKeyCredentialProvider() *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	return s.APIKeyCredentialProvider
}

func (s *GetAPIKeyCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAPIKeyCredentialProviderResponseBody) SetAPIKeyCredentialProvider(v *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) *GetAPIKeyCredentialProviderResponseBody {
	s.APIKeyCredentialProvider = v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBody) SetRequestId(v string) *GetAPIKeyCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBody) Validate() error {
	if s.APIKeyCredentialProvider != nil {
		if err := s.APIKeyCredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider struct {
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
	// example:
	//
	// 2025-12-18T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GoString() string {
	return s.String()
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetAPIKeyCredentialProviderName(v string) *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetCreateTime(v string) *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetCredentialProviderArn(v string) *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.CredentialProviderArn = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetDescription(v string) *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.Description = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) SetUpdateTime(v string) *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider {
	s.UpdateTime = &v
	return s
}

func (s *GetAPIKeyCredentialProviderResponseBodyAPIKeyCredentialProvider) Validate() error {
	return dara.Validate(s)
}
