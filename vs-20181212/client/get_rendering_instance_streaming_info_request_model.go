// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceStreamingInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *GetRenderingInstanceStreamingInfoRequest
	GetRenderingInstanceId() *string
}

type GetRenderingInstanceStreamingInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s GetRenderingInstanceStreamingInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceStreamingInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceStreamingInfoRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *GetRenderingInstanceStreamingInfoRequest) SetRenderingInstanceId(v string) *GetRenderingInstanceStreamingInfoRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *GetRenderingInstanceStreamingInfoRequest) Validate() error {
	return dara.Validate(s)
}
