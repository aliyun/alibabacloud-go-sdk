// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *ReleaseRenderingInstanceRequest
	GetRenderingInstanceId() *string
}

type ReleaseRenderingInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ReleaseRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ReleaseRenderingInstanceRequest) SetRenderingInstanceId(v string) *ReleaseRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ReleaseRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
