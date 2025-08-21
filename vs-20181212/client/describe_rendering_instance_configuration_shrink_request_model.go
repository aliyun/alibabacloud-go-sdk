// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationShrink(v string) *DescribeRenderingInstanceConfigurationShrinkRequest
	GetConfigurationShrink() *string
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceConfigurationShrinkRequest
	GetRenderingInstanceId() *string
}

type DescribeRenderingInstanceConfigurationShrinkRequest struct {
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *DescribeRenderingInstanceConfigurationShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceConfigurationShrinkRequest) SetConfigurationShrink(v string) *DescribeRenderingInstanceConfigurationShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationShrinkRequest) SetRenderingInstanceId(v string) *DescribeRenderingInstanceConfigurationShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
