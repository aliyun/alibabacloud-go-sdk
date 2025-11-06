// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayServiceResponseBody) *DeleteGatewayServiceResponse
	GetBody() *DeleteGatewayServiceResponseBody
}

type DeleteGatewayServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayServiceResponse) GetBody() *DeleteGatewayServiceResponseBody {
	return s.Body
}

func (s *DeleteGatewayServiceResponse) SetHeaders(v map[string]*string) *DeleteGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayServiceResponse) SetStatusCode(v int32) *DeleteGatewayServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayServiceResponse) SetBody(v *DeleteGatewayServiceResponseBody) *DeleteGatewayServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
