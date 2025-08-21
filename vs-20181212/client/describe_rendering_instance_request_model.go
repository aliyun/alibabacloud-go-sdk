// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceRequest
	GetRenderingInstanceId() *string
}

type DescribeRenderingInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DescribeRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceRequest) SetRenderingInstanceId(v string) *DescribeRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
