// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRenderingInstanceGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRenderingInstanceGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRenderingInstanceGatewayResponseBody) *DeleteRenderingInstanceGatewayResponse
	GetBody() *DeleteRenderingInstanceGatewayResponseBody
}

type DeleteRenderingInstanceGatewayResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRenderingInstanceGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRenderingInstanceGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRenderingInstanceGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRenderingInstanceGatewayResponse) GetBody() *DeleteRenderingInstanceGatewayResponseBody {
	return s.Body
}

func (s *DeleteRenderingInstanceGatewayResponse) SetHeaders(v map[string]*string) *DeleteRenderingInstanceGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteRenderingInstanceGatewayResponse) SetStatusCode(v int32) *DeleteRenderingInstanceGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRenderingInstanceGatewayResponse) SetBody(v *DeleteRenderingInstanceGatewayResponseBody) *DeleteRenderingInstanceGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteRenderingInstanceGatewayResponse) Validate() error {
	return dara.Validate(s)
}
