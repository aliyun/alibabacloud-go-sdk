// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayApiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayApiResponseBody) *DeleteGatewayApiResponse
	GetBody() *DeleteGatewayApiResponseBody
}

type DeleteGatewayApiResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayApiResponse) GetBody() *DeleteGatewayApiResponseBody {
	return s.Body
}

func (s *DeleteGatewayApiResponse) SetHeaders(v map[string]*string) *DeleteGatewayApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayApiResponse) SetStatusCode(v int32) *DeleteGatewayApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayApiResponse) SetBody(v *DeleteGatewayApiResponseBody) *DeleteGatewayApiResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
