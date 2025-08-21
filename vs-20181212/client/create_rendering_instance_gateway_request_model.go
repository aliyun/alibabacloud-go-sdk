// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayInstanceId(v string) *CreateRenderingInstanceGatewayRequest
	GetGatewayInstanceId() *string
	SetRenderingInstanceId(v string) *CreateRenderingInstanceGatewayRequest
	GetRenderingInstanceId() *string
}

type CreateRenderingInstanceGatewayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// render-xxx
	GatewayInstanceId *string `json:"GatewayInstanceId,omitempty" xml:"GatewayInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s CreateRenderingInstanceGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceGatewayRequest) GetGatewayInstanceId() *string {
	return s.GatewayInstanceId
}

func (s *CreateRenderingInstanceGatewayRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *CreateRenderingInstanceGatewayRequest) SetGatewayInstanceId(v string) *CreateRenderingInstanceGatewayRequest {
	s.GatewayInstanceId = &v
	return s
}

func (s *CreateRenderingInstanceGatewayRequest) SetRenderingInstanceId(v string) *CreateRenderingInstanceGatewayRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *CreateRenderingInstanceGatewayRequest) Validate() error {
	return dara.Validate(s)
}
