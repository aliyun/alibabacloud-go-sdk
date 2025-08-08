// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelSelectorShrink(v string) *ListPipelinesShrinkRequest
	GetLabelSelectorShrink() *string
}

type ListPipelinesShrinkRequest struct {
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListPipelinesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListPipelinesShrinkRequest) SetLabelSelectorShrink(v string) *ListPipelinesShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListPipelinesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
