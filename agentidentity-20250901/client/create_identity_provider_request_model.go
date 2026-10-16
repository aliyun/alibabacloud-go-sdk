// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedAudience(v []*string) *CreateIdentityProviderRequest
	GetAllowedAudience() []*string
	SetDescription(v string) *CreateIdentityProviderRequest
	GetDescription() *string
	SetDiscoveryURL(v string) *CreateIdentityProviderRequest
	GetDiscoveryURL() *string
	SetIdentityProviderName(v string) *CreateIdentityProviderRequest
	GetIdentityProviderName() *string
}

type CreateIdentityProviderRequest struct {
	AllowedAudience []*string `json:"AllowedAudience,omitempty" xml:"AllowedAudience,omitempty" type:"Repeated"`
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

func (s CreateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequest) GetAllowedAudience() []*string {
	return s.AllowedAudience
}

func (s *CreateIdentityProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIdentityProviderRequest) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *CreateIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateIdentityProviderRequest) SetAllowedAudience(v []*string) *CreateIdentityProviderRequest {
	s.AllowedAudience = v
	return s
}

func (s *CreateIdentityProviderRequest) SetDescription(v string) *CreateIdentityProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetDiscoveryURL(v string) *CreateIdentityProviderRequest {
	s.DiscoveryURL = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetIdentityProviderName(v string) *CreateIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateIdentityProviderRequest) Validate() error {
	return dara.Validate(s)
}
