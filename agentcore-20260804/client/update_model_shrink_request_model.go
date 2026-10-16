// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateModelShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateModelShrinkRequest
	GetClientToken() *string
}

type UpdateModelShrinkRequest struct {
	// The model update request body. You must specify at least one non-null parameter among description, contextSize, maxTokens, and capabilities.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token used for idempotence. This parameter is not supported.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateModelShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelShrinkRequest) SetBodyShrink(v string) *UpdateModelShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateModelShrinkRequest) SetClientToken(v string) *UpdateModelShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
