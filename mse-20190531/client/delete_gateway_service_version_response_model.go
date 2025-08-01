// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayServiceVersionResponseBody) *DeleteGatewayServiceVersionResponse
	GetBody() *DeleteGatewayServiceVersionResponseBody
}

type DeleteGatewayServiceVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayServiceVersionResponse) GetBody() *DeleteGatewayServiceVersionResponseBody {
	return s.Body
}

func (s *DeleteGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *DeleteGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayServiceVersionResponse) SetStatusCode(v int32) *DeleteGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponse) SetBody(v *DeleteGatewayServiceVersionResponseBody) *DeleteGatewayServiceVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayServiceVersionResponse) Validate() error {
	return dara.Validate(s)
}
