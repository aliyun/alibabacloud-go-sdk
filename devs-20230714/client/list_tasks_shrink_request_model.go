// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelSelectorShrink(v string) *ListTasksShrinkRequest
	GetLabelSelectorShrink() *string
}

type ListTasksShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListTasksShrinkRequest) SetLabelSelectorShrink(v string) *ListTasksShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
