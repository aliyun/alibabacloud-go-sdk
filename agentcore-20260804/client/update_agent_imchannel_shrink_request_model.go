// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateAgentIMChannelShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateAgentIMChannelShrinkRequest
	GetClientToken() *string
}

type UpdateAgentIMChannelShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The reserved idempotency token. The backend does not provide persistent idempotency guarantees in this phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAgentIMChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateAgentIMChannelShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentIMChannelShrinkRequest) SetBodyShrink(v string) *UpdateAgentIMChannelShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateAgentIMChannelShrinkRequest) SetClientToken(v string) *UpdateAgentIMChannelShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentIMChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
