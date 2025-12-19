// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedAudience(v []*string) *UpdateIdentityProviderRequest
	GetAllowedAudience() []*string
	SetDescription(v string) *UpdateIdentityProviderRequest
	GetDescription() *string
	SetDiscoveryURL(v string) *UpdateIdentityProviderRequest
	GetDiscoveryURL() *string
	SetIdentityProviderName(v string) *UpdateIdentityProviderRequest
	GetIdentityProviderName() *string
}

type UpdateIdentityProviderRequest struct {
	AllowedAudience []*string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty" type:"Repeated"`
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

func (s UpdateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequest) GetAllowedAudience() []*string {
	return s.AllowedAudience
}

func (s *UpdateIdentityProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateIdentityProviderRequest) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *UpdateIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *UpdateIdentityProviderRequest) SetAllowedAudience(v []*string) *UpdateIdentityProviderRequest {
	s.AllowedAudience = v
	return s
}

func (s *UpdateIdentityProviderRequest) SetDescription(v string) *UpdateIdentityProviderRequest {
	s.Description = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetDiscoveryURL(v string) *UpdateIdentityProviderRequest {
	s.DiscoveryURL = &v
	return s
}

func (s *UpdateIdentityProviderRequest) SetIdentityProviderName(v string) *UpdateIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
