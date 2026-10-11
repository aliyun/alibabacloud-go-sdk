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
	// The JSON configuration of the custom agent. You must specify name, artifact, roleArn, and runtime.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// Idempotency is not currently guaranteed. Do not rely on this parameter to prevent duplicate requests.
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
