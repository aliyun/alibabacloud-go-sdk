// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManagedAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateManagedAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateManagedAgentShrinkRequest
	GetClientToken() *string
}

type CreateManagedAgentShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The reserved idempotency token. The backend does not provide idempotency guarantees in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateManagedAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateManagedAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateManagedAgentShrinkRequest) SetBodyShrink(v string) *CreateManagedAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateManagedAgentShrinkRequest) SetClientToken(v string) *CreateManagedAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateManagedAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
