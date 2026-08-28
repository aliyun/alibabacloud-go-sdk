// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentIMChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateAgentIMChannelShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateAgentIMChannelShrinkRequest
	GetClientToken() *string
}

type CreateAgentIMChannelShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// A reserved idempotency token. The backend does not provide persistent idempotency guarantees in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAgentIMChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateAgentIMChannelShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentIMChannelShrinkRequest) SetBodyShrink(v string) *CreateAgentIMChannelShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateAgentIMChannelShrinkRequest) SetClientToken(v string) *CreateAgentIMChannelShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentIMChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
