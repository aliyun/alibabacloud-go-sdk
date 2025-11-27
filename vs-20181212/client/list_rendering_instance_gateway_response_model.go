// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstanceGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingInstanceGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingInstanceGatewayResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingInstanceGatewayResponseBody) *ListRenderingInstanceGatewayResponse
	GetBody() *ListRenderingInstanceGatewayResponseBody
}

type ListRenderingInstanceGatewayResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingInstanceGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingInstanceGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstanceGatewayResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingInstanceGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingInstanceGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingInstanceGatewayResponse) GetBody() *ListRenderingInstanceGatewayResponseBody {
	return s.Body
}

func (s *ListRenderingInstanceGatewayResponse) SetHeaders(v map[string]*string) *ListRenderingInstanceGatewayResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingInstanceGatewayResponse) SetStatusCode(v int32) *ListRenderingInstanceGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingInstanceGatewayResponse) SetBody(v *ListRenderingInstanceGatewayResponseBody) *ListRenderingInstanceGatewayResponse {
	s.Body = v
	return s
}

func (s *ListRenderingInstanceGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
