// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelCredentialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateAgentIMChannelCredentialShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateAgentIMChannelCredentialShrinkRequest
	GetClientToken() *string
}

type UpdateAgentIMChannelCredentialShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// A reserved idempotency token. The backend does not provide persistent idempotence guarantee in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAgentIMChannelCredentialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateAgentIMChannelCredentialShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentIMChannelCredentialShrinkRequest) SetBodyShrink(v string) *UpdateAgentIMChannelCredentialShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialShrinkRequest) SetClientToken(v string) *UpdateAgentIMChannelCredentialShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
