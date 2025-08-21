// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeNamesShrink(v string) *DescribeRenderingInstanceSettingsShrinkRequest
	GetAttributeNamesShrink() *string
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceSettingsShrinkRequest
	GetRenderingInstanceId() *string
}

type DescribeRenderingInstanceSettingsShrinkRequest struct {
	AttributeNamesShrink *string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DescribeRenderingInstanceSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceSettingsShrinkRequest) GetAttributeNamesShrink() *string {
	return s.AttributeNamesShrink
}

func (s *DescribeRenderingInstanceSettingsShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceSettingsShrinkRequest) SetAttributeNamesShrink(v string) *DescribeRenderingInstanceSettingsShrinkRequest {
	s.AttributeNamesShrink = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsShrinkRequest) SetRenderingInstanceId(v string) *DescribeRenderingInstanceSettingsShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
