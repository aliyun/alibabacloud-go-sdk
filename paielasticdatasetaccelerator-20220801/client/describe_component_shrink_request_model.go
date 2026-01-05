// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderTemplate(v bool) *DescribeComponentShrinkRequest
	GetRenderTemplate() *bool
	SetValuesShrink(v string) *DescribeComponentShrinkRequest
	GetValuesShrink() *string
}

type DescribeComponentShrinkRequest struct {
	// example:
	//
	// true
	RenderTemplate *bool   `json:"RenderTemplate,omitempty" xml:"RenderTemplate,omitempty"`
	ValuesShrink   *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeComponentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentShrinkRequest) GetRenderTemplate() *bool {
	return s.RenderTemplate
}

func (s *DescribeComponentShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *DescribeComponentShrinkRequest) SetRenderTemplate(v bool) *DescribeComponentShrinkRequest {
	s.RenderTemplate = &v
	return s
}

func (s *DescribeComponentShrinkRequest) SetValuesShrink(v string) *DescribeComponentShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *DescribeComponentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
