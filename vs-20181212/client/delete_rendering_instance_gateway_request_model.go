// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *DeleteRenderingInstanceGatewayRequest
	GetRenderingInstanceId() *string
}

type DeleteRenderingInstanceGatewayRequest struct {
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DeleteRenderingInstanceGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceGatewayRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DeleteRenderingInstanceGatewayRequest) SetRenderingInstanceId(v string) *DeleteRenderingInstanceGatewayRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DeleteRenderingInstanceGatewayRequest) Validate() error {
	return dara.Validate(s)
}
