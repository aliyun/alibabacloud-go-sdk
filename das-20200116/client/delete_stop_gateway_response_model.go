// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStopGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStopGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStopGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStopGatewayResponseBody) *DeleteStopGatewayResponse
	GetBody() *DeleteStopGatewayResponseBody
}

type DeleteStopGatewayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStopGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStopGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStopGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteStopGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStopGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStopGatewayResponse) GetBody() *DeleteStopGatewayResponseBody {
	return s.Body
}

func (s *DeleteStopGatewayResponse) SetHeaders(v map[string]*string) *DeleteStopGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteStopGatewayResponse) SetStatusCode(v int32) *DeleteStopGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStopGatewayResponse) SetBody(v *DeleteStopGatewayResponseBody) *DeleteStopGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteStopGatewayResponse) Validate() error {
	return dara.Validate(s)
}
