// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderSCIMShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateIdentityProviderSCIMShrinkRequest
	GetBodyShrink() *string
}

type UpdateIdentityProviderSCIMShrinkRequest struct {
	// The business parameters submitted in the request body. The body parameters of public APIs are mapped to backend wrappers by the gateway and must not be mixed with direct backend body wrappers.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdentityProviderSCIMShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateIdentityProviderSCIMShrinkRequest) SetBodyShrink(v string) *UpdateIdentityProviderSCIMShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateIdentityProviderSCIMShrinkRequest) Validate() error {
	return dara.Validate(s)
}
