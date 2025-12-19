// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedAudienceShrink(v string) *UpdateIdentityProviderShrinkRequest
	GetAllowedAudienceShrink() *string
	SetDescription(v string) *UpdateIdentityProviderShrinkRequest
	GetDescription() *string
	SetDiscoveryURL(v string) *UpdateIdentityProviderShrinkRequest
	GetDiscoveryURL() *string
	SetIdentityProviderName(v string) *UpdateIdentityProviderShrinkRequest
	GetIdentityProviderName() *string
}

type UpdateIdentityProviderShrinkRequest struct {
	AllowedAudienceShrink *string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty"`
	// example:
	//
	// example identity provider
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

func (s UpdateIdentityProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderShrinkRequest) GetAllowedAudienceShrink() *string {
	return s.AllowedAudienceShrink
}

func (s *UpdateIdentityProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIdentityProviderShrinkRequest) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *UpdateIdentityProviderShrinkRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *UpdateIdentityProviderShrinkRequest) SetAllowedAudienceShrink(v string) *UpdateIdentityProviderShrinkRequest {
	s.AllowedAudienceShrink = &v
	return s
}

func (s *UpdateIdentityProviderShrinkRequest) SetDescription(v string) *UpdateIdentityProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateIdentityProviderShrinkRequest) SetDiscoveryURL(v string) *UpdateIdentityProviderShrinkRequest {
	s.DiscoveryURL = &v
	return s
}

func (s *UpdateIdentityProviderShrinkRequest) SetIdentityProviderName(v string) *UpdateIdentityProviderShrinkRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *UpdateIdentityProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
