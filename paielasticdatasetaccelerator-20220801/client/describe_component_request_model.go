// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderTemplate(v bool) *DescribeComponentRequest
	GetRenderTemplate() *bool
	SetValues(v map[string]interface{}) *DescribeComponentRequest
	GetValues() map[string]interface{}
}

type DescribeComponentRequest struct {
	// example:
	//
	// true
	RenderTemplate *bool                  `json:"RenderTemplate,omitempty" xml:"RenderTemplate,omitempty"`
	Values         map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentRequest) GetRenderTemplate() *bool {
	return s.RenderTemplate
}

func (s *DescribeComponentRequest) GetValues() map[string]interface{} {
	return s.Values
}

func (s *DescribeComponentRequest) SetRenderTemplate(v bool) *DescribeComponentRequest {
	s.RenderTemplate = &v
	return s
}

func (s *DescribeComponentRequest) SetValues(v map[string]interface{}) *DescribeComponentRequest {
	s.Values = v
	return s
}

func (s *DescribeComponentRequest) Validate() error {
	return dara.Validate(s)
}
