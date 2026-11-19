// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *DeleteEventsShrinkRequest
	GetBodyShrink() *string
}

type DeleteEventsShrinkRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventsShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *DeleteEventsShrinkRequest) SetBodyShrink(v string) *DeleteEventsShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *DeleteEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
