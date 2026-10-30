// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateManagedAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateManagedAgentShrinkRequest
	GetClientToken() *string
}

type UpdateManagedAgentShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The reserved idempotency token. The backend does not provide idempotency guarantees in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateManagedAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateManagedAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateManagedAgentShrinkRequest) SetBodyShrink(v string) *UpdateManagedAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateManagedAgentShrinkRequest) SetClientToken(v string) *UpdateManagedAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateManagedAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
