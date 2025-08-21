// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *RebootRenderingInstanceRequest
	GetRenderingInstanceId() *string
}

type RebootRenderingInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RebootRenderingInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootRenderingInstanceRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RebootRenderingInstanceRequest) SetRenderingInstanceId(v string) *RebootRenderingInstanceRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *RebootRenderingInstanceRequest) Validate() error {
	return dara.Validate(s)
}
