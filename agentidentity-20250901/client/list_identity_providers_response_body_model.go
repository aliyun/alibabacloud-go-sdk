// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviders(v []*ListIdentityProvidersResponseBodyIdentityProviders) *ListIdentityProvidersResponseBody
	GetIdentityProviders() []*ListIdentityProvidersResponseBodyIdentityProviders
	SetMaxResults(v int32) *ListIdentityProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIdentityProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIdentityProvidersResponseBody
	GetTotalCount() *int32
}

type ListIdentityProvidersResponseBody struct {
	IdentityProviders []*ListIdentityProvidersResponseBodyIdentityProviders `json:"IdentityProviders,omitempty" xml:"IdentityProviders,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lsy4U8Dgz6TCndCo6o5TB8
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D325DF9D-7CA8-5C47-BA0C-9A74873F2BE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 33
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIdentityProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBody) GetIdentityProviders() []*ListIdentityProvidersResponseBodyIdentityProviders {
	return s.IdentityProviders
}

func (s *ListIdentityProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIdentityProvidersResponseBody) SetIdentityProviders(v []*ListIdentityProvidersResponseBodyIdentityProviders) *ListIdentityProvidersResponseBody {
	s.IdentityProviders = v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetMaxResults(v int32) *ListIdentityProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetNextToken(v string) *ListIdentityProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetRequestId(v string) *ListIdentityProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetTotalCount(v int32) *ListIdentityProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) Validate() error {
	if s.IdentityProviders != nil {
		for _, item := range s.IdentityProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProvidersResponseBodyIdentityProviders struct {
	AllowedAudience []*string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ods_api_public_ios_wanxin_boxgame_user_behavior_integration_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://oauth.aliyun.com/.well-known/openid-configuration
	DiscoveryURL *string `json:"DiscoveryURL,omitempty" xml:"DiscoveryURL,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:identityprovider/identity-provider-okta
	IdentityProviderArn *string `json:"IdentityProviderArn,omitempty" xml:"IdentityProviderArn,omitempty"`
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// example:
	//
	// 2025-06-09T02:04:23
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListIdentityProvidersResponseBodyIdentityProviders) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyIdentityProviders) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetAllowedAudience() []*string {
	return s.AllowedAudience
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetDescription() *string {
	return s.Description
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderArn() *string {
	return s.IdentityProviderArn
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetAllowedAudience(v []*string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.AllowedAudience = v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetCreateTime(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.CreateTime = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetDescription(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.Description = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetDiscoveryURL(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.DiscoveryURL = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderArn(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderArn = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderName(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderName = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetUpdateTime(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) Validate() error {
	return dara.Validate(s)
}
