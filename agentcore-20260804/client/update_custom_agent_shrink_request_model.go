// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateCustomAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateCustomAgentShrinkRequest
	GetClientToken() *string
}

type UpdateCustomAgentShrinkRequest struct {
	// The update configuration for the custom agent. Provide at least one top-level field. Omitted top-level fields retain their original values. When providing a configuration object, specify its required subfields and include any optional configurations that you want to retain. Arrays are replaced in their entirety.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// Idempotency is not currently guaranteed. Do not rely on this parameter to prevent duplicate requests.
	//
	// example:
	//
	// custom-agent-request-20261010-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateCustomAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomAgentShrinkRequest) SetBodyShrink(v string) *UpdateCustomAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetClientToken(v string) *UpdateCustomAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
