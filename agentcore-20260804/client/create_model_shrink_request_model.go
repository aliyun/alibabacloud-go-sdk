// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateModelShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateModelShrinkRequest
	GetClientToken() *string
}

type CreateModelShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token for idempotence. Not currently supported.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateModelShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelShrinkRequest) SetBodyShrink(v string) *CreateModelShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateModelShrinkRequest) SetClientToken(v string) *CreateModelShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
