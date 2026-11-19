// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateServiceEndpointShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateServiceEndpointShrinkRequest
	GetClientToken() *string
}

type UpdateServiceEndpointShrinkRequest struct {
	// The service endpoint configurations to be modified in this request. Submit the supported modifiable name, bound version, invocation authentication configuration, elastic policy, or public network inbound switch based on the endpoint type. Router and hybrid ingress endpoints cannot be renamed. DEFAULT endpoints do not accept name or authentication configuration updates. When updating the elastic policy, you must specify both updateElasticity=true and clientToken.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token. This parameter is required when updateElasticity is set to true and is verified by the server. Reuse the token and the original request for the same retry. Use a new token if the logical content changes. Passing different content with the same token returns InvalidParameter. The token must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), colons (:), and hyphens (-). The scope is the tenant and workspace. The receipt is retained for 7 days from the time of successful acceptance. Replays do not extend the retention period. After expiration, it is verified again as a new submission.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateServiceEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateServiceEndpointShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceEndpointShrinkRequest) SetBodyShrink(v string) *UpdateServiceEndpointShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateServiceEndpointShrinkRequest) SetClientToken(v string) *UpdateServiceEndpointShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
