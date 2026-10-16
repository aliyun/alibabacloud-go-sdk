// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProvider(v *GetIdentityProviderResponseBodyIdentityProvider) *GetIdentityProviderResponseBody
	GetIdentityProvider() *GetIdentityProviderResponseBodyIdentityProvider
	SetRequestId(v string) *GetIdentityProviderResponseBody
	GetRequestId() *string
}

type GetIdentityProviderResponseBody struct {
	IdentityProvider *GetIdentityProviderResponseBodyIdentityProvider `json:"IdentityProvider,omitempty" xml:"IdentityProvider,omitempty" type:"Struct"`
	// example:
	//
	// 117E1427-1D4C-59CF-A0A8-D7DCDFE88863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBody) GetIdentityProvider() *GetIdentityProviderResponseBodyIdentityProvider {
	return s.IdentityProvider
}

func (s *GetIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderResponseBody) SetIdentityProvider(v *GetIdentityProviderResponseBodyIdentityProvider) *GetIdentityProviderResponseBody {
	s.IdentityProvider = v
	return s
}

func (s *GetIdentityProviderResponseBody) SetRequestId(v string) *GetIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderResponseBody) Validate() error {
	if s.IdentityProvider != nil {
		if err := s.IdentityProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyIdentityProvider struct {
	AllowedAudience []*string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example agent
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
	// 2025-12-18T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProvider) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProvider) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetAllowedAudience() []*string {
	return s.AllowedAudience
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetDescription() *string {
	return s.Description
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetIdentityProviderArn() *string {
	return s.IdentityProviderArn
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetAllowedAudience(v []*string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.AllowedAudience = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetCreateTime(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.CreateTime = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetDescription(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.Description = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetDiscoveryURL(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.DiscoveryURL = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetIdentityProviderArn(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.IdentityProviderArn = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetIdentityProviderName(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.IdentityProviderName = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) SetUpdateTime(v string) *GetIdentityProviderResponseBodyIdentityProvider {
	s.UpdateTime = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProvider) Validate() error {
	return dara.Validate(s)
}
