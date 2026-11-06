// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateIdentityProviderShrinkRequest
	GetBodyShrink() *string
}

type CreateIdentityProviderShrinkRequest struct {
	// The request body for binding an external identity provider.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentityProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateIdentityProviderShrinkRequest) SetBodyShrink(v string) *CreateIdentityProviderShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateIdentityProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
