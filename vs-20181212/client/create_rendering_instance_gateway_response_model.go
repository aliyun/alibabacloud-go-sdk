// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRenderingInstanceGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRenderingInstanceGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateRenderingInstanceGatewayResponseBody) *CreateRenderingInstanceGatewayResponse
	GetBody() *CreateRenderingInstanceGatewayResponseBody
}

type CreateRenderingInstanceGatewayResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRenderingInstanceGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRenderingInstanceGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRenderingInstanceGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRenderingInstanceGatewayResponse) GetBody() *CreateRenderingInstanceGatewayResponseBody {
	return s.Body
}

func (s *CreateRenderingInstanceGatewayResponse) SetHeaders(v map[string]*string) *CreateRenderingInstanceGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateRenderingInstanceGatewayResponse) SetStatusCode(v int32) *CreateRenderingInstanceGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRenderingInstanceGatewayResponse) SetBody(v *CreateRenderingInstanceGatewayResponseBody) *CreateRenderingInstanceGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateRenderingInstanceGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
