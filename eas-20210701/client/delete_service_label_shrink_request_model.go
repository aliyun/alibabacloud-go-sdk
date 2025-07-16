// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLabelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeysShrink(v string) *DeleteServiceLabelShrinkRequest
	GetKeysShrink() *string
}

type DeleteServiceLabelShrinkRequest struct {
	// The service tags that you want to delete.
	//
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteServiceLabelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *DeleteServiceLabelShrinkRequest) SetKeysShrink(v string) *DeleteServiceLabelShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *DeleteServiceLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
