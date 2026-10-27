// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateMemoryShrinkRequest
	GetBodyShrink() *string
}

type UpdateMemoryShrinkRequest struct {
	// The update memory request.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateMemoryShrinkRequest) SetBodyShrink(v string) *UpdateMemoryShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
