// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExternalAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateExternalAgentShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateExternalAgentShrinkRequest
	GetClientToken() *string
}

type UpdateExternalAgentShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// A reserved idempotency token. The backend does not guarantee idempotency in the current version.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateExternalAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExternalAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateExternalAgentShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateExternalAgentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExternalAgentShrinkRequest) SetBodyShrink(v string) *UpdateExternalAgentShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateExternalAgentShrinkRequest) SetClientToken(v string) *UpdateExternalAgentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExternalAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
