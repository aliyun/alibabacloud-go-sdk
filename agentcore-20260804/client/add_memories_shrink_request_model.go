// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *AddMemoriesShrinkRequest
	GetBodyShrink() *string
}

type AddMemoriesShrinkRequest struct {
	// The request body for writing to memory.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMemoriesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *AddMemoriesShrinkRequest) SetBodyShrink(v string) *AddMemoriesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *AddMemoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
