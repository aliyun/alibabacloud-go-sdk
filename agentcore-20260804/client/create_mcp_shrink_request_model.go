// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateMcpShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateMcpShrinkRequest
	GetClientToken() *string
}

type CreateMcpShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client idempotency token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateMcpShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMcpShrinkRequest) SetBodyShrink(v string) *CreateMcpShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateMcpShrinkRequest) SetClientToken(v string) *CreateMcpShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
