// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateAccessApplicationShrinkRequest
	GetBodyShrink() *string
}

type UpdateAccessApplicationShrinkRequest struct {
	// The request body for updating the access application.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessApplicationShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateAccessApplicationShrinkRequest) SetBodyShrink(v string) *UpdateAccessApplicationShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateAccessApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
