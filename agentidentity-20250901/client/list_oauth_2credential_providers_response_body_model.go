// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOAuth2CredentialProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOAuth2CredentialProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOAuth2CredentialProvidersResponseBody
	GetNextToken() *string
	SetOAuth2CredentialProviders(v []*ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) *ListOAuth2CredentialProvidersResponseBody
	GetOAuth2CredentialProviders() []*ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders
	SetRequestId(v string) *ListOAuth2CredentialProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOAuth2CredentialProvidersResponseBody
	GetTotalCount() *int32
}

type ListOAuth2CredentialProvidersResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lURp7zQ6t4/95+g/5jN5r/
	NextToken                 *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OAuth2CredentialProviders []*ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders `json:"OAuth2CredentialProviders,omitempty" xml:"OAuth2CredentialProviders,omitempty" type:"Repeated"`
	// example:
	//
	// 690C4C7A-AFB9-54ED-9A06-84426F5C6369
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 400
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOAuth2CredentialProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOAuth2CredentialProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOAuth2CredentialProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOAuth2CredentialProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOAuth2CredentialProvidersResponseBody) GetOAuth2CredentialProviders() []*ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	return s.OAuth2CredentialProviders
}

func (s *ListOAuth2CredentialProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOAuth2CredentialProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOAuth2CredentialProvidersResponseBody) SetMaxResults(v int32) *ListOAuth2CredentialProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBody) SetNextToken(v string) *ListOAuth2CredentialProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBody) SetOAuth2CredentialProviders(v []*ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) *ListOAuth2CredentialProvidersResponseBody {
	s.OAuth2CredentialProviders = v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBody) SetRequestId(v string) *ListOAuth2CredentialProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBody) SetTotalCount(v int32) *ListOAuth2CredentialProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBody) Validate() error {
	if s.OAuth2CredentialProviders != nil {
		for _, item := range s.OAuth2CredentialProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders struct {
	// example:
	//
	// https://agentidentitydata.cn-beijing.aliyuncs.com/oauth2/callback/d51171bc-0dae-3219-8e65-6b4cdafa3beb
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default/oauth2credentialprovider/oauth-provider-aliyun
	CredentialProviderArn *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
	// example:
	//
	// AliyunOAuth2
	CredentialProviderVendor *string `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	// example:
	//
	// example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) String() string {
	return dara.Prettify(s)
}

func (s ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GoString() string {
	return s.String()
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetDescription() *string {
	return s.Description
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetCallbackURL(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.CallbackURL = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetCreateTime(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.CreateTime = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetCredentialProviderArn(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.CredentialProviderArn = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetCredentialProviderVendor(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.CredentialProviderVendor = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetDescription(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.Description = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetOAuth2CredentialProviderName(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.OAuth2ProviderConfig = v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) SetUpdateTime(v string) *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponseBodyOAuth2CredentialProviders) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
