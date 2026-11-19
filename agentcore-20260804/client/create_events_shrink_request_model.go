// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateEventsShrinkRequest
	GetBodyShrink() *string
}

type CreateEventsShrinkRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateEventsShrinkRequest) SetBodyShrink(v string) *CreateEventsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
