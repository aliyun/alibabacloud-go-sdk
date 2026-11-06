// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateCustomAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateCustomAgentShrinkRequest
	GetClientToken() *string
}

type CreateCustomAgentShrinkRequest struct {
	// The JSON configuration of the custom code or image agent. You must specify name, artifact, roleArn, and runtime.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The idempotency token passed as a query parameter. This parameter is required when explicitly submitting a customImage or setting container.type to custom, and is optional for ACR image and code package submissions. The token must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), colons (:), and hyphens (-). Reuse the same token for retries of the same submission. If the same token is used with different parameters, the request is rejected.
	//
	// example:
	//
	// custom-agent-request-20261010-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateCustomAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomAgentShrinkRequest) SetBodyShrink(v string) *CreateCustomAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetClientToken(v string) *CreateCustomAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
