// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewaySlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewaySlbResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewaySlbResponseBody) *DeleteGatewaySlbResponse
	GetBody() *DeleteGatewaySlbResponseBody
}

type DeleteGatewaySlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewaySlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewaySlbResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySlbResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewaySlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewaySlbResponse) GetBody() *DeleteGatewaySlbResponseBody {
	return s.Body
}

func (s *DeleteGatewaySlbResponse) SetHeaders(v map[string]*string) *DeleteGatewaySlbResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewaySlbResponse) SetStatusCode(v int32) *DeleteGatewaySlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewaySlbResponse) SetBody(v *DeleteGatewaySlbResponseBody) *DeleteGatewaySlbResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewaySlbResponse) Validate() error {
	return dara.Validate(s)
}
