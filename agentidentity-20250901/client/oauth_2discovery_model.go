// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2Discovery interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServerMetadata(v *AuthorizationServerMetadata) *OAuth2Discovery
	GetAuthorizationServerMetadata() *AuthorizationServerMetadata
	SetDiscoveryURL(v string) *OAuth2Discovery
	GetDiscoveryURL() *string
}

type OAuth2Discovery struct {
	AuthorizationServerMetadata *AuthorizationServerMetadata `json:"AuthorizationServerMetadata,omitempty" xml:"AuthorizationServerMetadata,omitempty"`
	DiscoveryURL                *string                      `json:"DiscoveryURL,omitempty" xml:"DiscoveryURL,omitempty"`
}

func (s OAuth2Discovery) String() string {
	return dara.Prettify(s)
}

func (s OAuth2Discovery) GoString() string {
	return s.String()
}

func (s *OAuth2Discovery) GetAuthorizationServerMetadata() *AuthorizationServerMetadata {
	return s.AuthorizationServerMetadata
}

func (s *OAuth2Discovery) GetDiscoveryURL() *string {
	return s.DiscoveryURL
}

func (s *OAuth2Discovery) SetAuthorizationServerMetadata(v *AuthorizationServerMetadata) *OAuth2Discovery {
	s.AuthorizationServerMetadata = v
	return s
}

func (s *OAuth2Discovery) SetDiscoveryURL(v string) *OAuth2Discovery {
	s.DiscoveryURL = &v
	return s
}

func (s *OAuth2Discovery) Validate() error {
	if s.AuthorizationServerMetadata != nil {
		if err := s.AuthorizationServerMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
