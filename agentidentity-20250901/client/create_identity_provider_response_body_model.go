// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProvider(v *CreateIdentityProviderResponseBodyIdentityProvider) *CreateIdentityProviderResponseBody
	GetIdentityProvider() *CreateIdentityProviderResponseBodyIdentityProvider
	SetRequestId(v string) *CreateIdentityProviderResponseBody
	GetRequestId() *string
}

type CreateIdentityProviderResponseBody struct {
	IdentityProvider *CreateIdentityProviderResponseBodyIdentityProvider `json:"IdentityProvider,omitempty" xml:"IdentityProvider,omitempty" type:"Struct"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBody) GetIdentityProvider() *CreateIdentityProviderResponseBodyIdentityProvider {
	return s.IdentityProvider
}

func (s *CreateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderResponseBody) SetIdentityProvider(v *CreateIdentityProviderResponseBodyIdentityProvider) *CreateIdentityProviderResponseBody {
	s.IdentityProvider = v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetRequestId(v string) *CreateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) Validate() error {
	if s.IdentityProvider != nil {
		if err := s.IdentityProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderResponseBodyIdentityProvider struct {
	AllowedAudience      []*string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty" type:"Repeated"`
	CreateTime           *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DiscoveryURL         *string   `json:"DiscoveryURL,omitempty" xml:"DiscoveryURL,omitempty"`
	IdentityProviderArn  *string   `json:"IdentityProviderArn,omitempty" xml:"IdentityProviderArn,omitempty"`
	IdentityProviderName *string   `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	UpdateTime           *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateIdentityProviderResponseBodyIdentityProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyIdentityProvider) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetAllowedAudience() []*string {
	return s.AllowedAudience
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetDescription() *string {
	return s.Description
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetIdentityProviderArn() *string {
	return s.IdentityProviderArn
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetAllowedAudience(v []*string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.AllowedAudience = v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetCreateTime(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.CreateTime = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetDescription(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.Description = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetDiscoveryURL(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.DiscoveryURL = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetIdentityProviderArn(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.IdentityProviderArn = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetIdentityProviderName(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) SetUpdateTime(v string) *CreateIdentityProviderResponseBodyIdentityProvider {
	s.UpdateTime = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyIdentityProvider) Validate() error {
	return dara.Validate(s)
}
