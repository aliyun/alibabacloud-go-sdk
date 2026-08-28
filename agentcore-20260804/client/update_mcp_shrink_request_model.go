// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateMcpShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateMcpShrinkRequest
	GetClientToken() *string
}

type UpdateMcpShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client idempotency token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateMcpShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpShrinkRequest) SetBodyShrink(v string) *UpdateMcpShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateMcpShrinkRequest) SetClientToken(v string) *UpdateMcpShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
