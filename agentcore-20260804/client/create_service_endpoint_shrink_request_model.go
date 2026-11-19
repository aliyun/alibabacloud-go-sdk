// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateServiceEndpointShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateServiceEndpointShrinkRequest
	GetClientToken() *string
}

type CreateServiceEndpointShrinkRequest struct {
	// The name, routing target, and invocation authentication configuration of the service endpoint to be created, along with the optional initial elasticity configuration and public access switch.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// Required when elasticity initialization is included. It is verified by the server. The same token and original request are reused for the same retry. If the logical content changes, a new token is used. The same token with different content returns InvalidParameter. The token must be 1 to 128 characters in length and can contain letters, digits, and the following special characters: . _ : -. The scope is the tenant and workspace. The receipt is retained for 7 days from the time of successful acceptance. Replays do not extend the retention period. After expiration, it is verified again as a new submission.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateServiceEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateServiceEndpointShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceEndpointShrinkRequest) SetBodyShrink(v string) *CreateServiceEndpointShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateServiceEndpointShrinkRequest) SetClientToken(v string) *CreateServiceEndpointShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
