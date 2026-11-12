// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateExternalAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateExternalAgentShrinkRequest
	GetClientToken() *string
}

type CreateExternalAgentShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The reserved query parameter. Currently, the backend does not provide idempotency guarantees for deduplication or restoring the original acceptance result based on this identifier.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateExternalAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateExternalAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExternalAgentShrinkRequest) SetBodyShrink(v string) *CreateExternalAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateExternalAgentShrinkRequest) SetClientToken(v string) *CreateExternalAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExternalAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
