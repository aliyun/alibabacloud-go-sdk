// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayServiceVersionResponseBody) *AddGatewayServiceVersionResponse
	GetBody() *AddGatewayServiceVersionResponseBody
}

type AddGatewayServiceVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayServiceVersionResponse) GetBody() *AddGatewayServiceVersionResponseBody {
	return s.Body
}

func (s *AddGatewayServiceVersionResponse) SetHeaders(v map[string]*string) *AddGatewayServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayServiceVersionResponse) SetStatusCode(v int32) *AddGatewayServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayServiceVersionResponse) SetBody(v *AddGatewayServiceVersionResponseBody) *AddGatewayServiceVersionResponse {
	s.Body = v
	return s
}

func (s *AddGatewayServiceVersionResponse) Validate() error {
	return dara.Validate(s)
}
