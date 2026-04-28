// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse
	GetBody() *DeleteGatewayResponseBody
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayResponse) GetBody() *DeleteGatewayResponseBody {
	return s.Body
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
