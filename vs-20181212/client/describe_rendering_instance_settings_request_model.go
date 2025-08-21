// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeNames(v []*string) *DescribeRenderingInstanceSettingsRequest
	GetAttributeNames() []*string
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceSettingsRequest
	GetRenderingInstanceId() *string
}

type DescribeRenderingInstanceSettingsRequest struct {
	AttributeNames []*string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty" type:"Repeated"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DescribeRenderingInstanceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceSettingsRequest) GetAttributeNames() []*string {
	return s.AttributeNames
}

func (s *DescribeRenderingInstanceSettingsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceSettingsRequest) SetAttributeNames(v []*string) *DescribeRenderingInstanceSettingsRequest {
	s.AttributeNames = v
	return s
}

func (s *DescribeRenderingInstanceSettingsRequest) SetRenderingInstanceId(v string) *DescribeRenderingInstanceSettingsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsRequest) Validate() error {
	return dara.Validate(s)
}
