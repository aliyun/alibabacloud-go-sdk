// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedAudienceShrink(v string) *CreateIdentityProviderShrinkRequest
	GetAllowedAudienceShrink() *string
	SetDescription(v string) *CreateIdentityProviderShrinkRequest
	GetDescription() *string
	SetDiscoveryURL(v string) *CreateIdentityProviderShrinkRequest
	GetDiscoveryURL() *string
	SetIdentityProviderName(v string) *CreateIdentityProviderShrinkRequest
	GetIdentityProviderName() *string
}

type CreateIdentityProviderShrinkRequest struct {
	AllowedAudienceShrink *string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty"`
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
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
}

func (s CreateIdentityProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderShrinkRequest) GetAllowedAudienceShrink() *string {
	return s.AllowedAudienceShrink
}

func (s *CreateIdentityProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIdentityProviderShrinkRequest) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *CreateIdentityProviderShrinkRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateIdentityProviderShrinkRequest) SetAllowedAudienceShrink(v string) *CreateIdentityProviderShrinkRequest {
	s.AllowedAudienceShrink = &v
	return s
}

func (s *CreateIdentityProviderShrinkRequest) SetDescription(v string) *CreateIdentityProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIdentityProviderShrinkRequest) SetDiscoveryURL(v string) *CreateIdentityProviderShrinkRequest {
	s.DiscoveryURL = &v
	return s
}

func (s *CreateIdentityProviderShrinkRequest) SetIdentityProviderName(v string) *CreateIdentityProviderShrinkRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateIdentityProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
