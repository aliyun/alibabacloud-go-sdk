// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateIdentityProviderShrinkRequest
	GetBodyShrink() *string
}

type UpdateIdentityProviderShrinkRequest struct {
	// The request body for updating the external identity provider.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdentityProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateIdentityProviderShrinkRequest) SetBodyShrink(v string) *UpdateIdentityProviderShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateIdentityProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
