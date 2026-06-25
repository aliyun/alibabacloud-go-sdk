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
	SetLabelKeysShrink(v string) *DeleteServiceLabelShrinkRequest
	GetLabelKeysShrink() *string
}

type DeleteServiceLabelShrinkRequest struct {
	// Deprecated
	//
	// The keys of the service labels that you want to delete. This parameter is deprecated. Use the LabelKeys parameter instead.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The keys of the service labels that you want to delete.
	LabelKeysShrink *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
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

func (s *DeleteServiceLabelShrinkRequest) GetLabelKeysShrink() *string {
	return s.LabelKeysShrink
}

func (s *DeleteServiceLabelShrinkRequest) SetKeysShrink(v string) *DeleteServiceLabelShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *DeleteServiceLabelShrinkRequest) SetLabelKeysShrink(v string) *DeleteServiceLabelShrinkRequest {
	s.LabelKeysShrink = &v
	return s
}

func (s *DeleteServiceLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
