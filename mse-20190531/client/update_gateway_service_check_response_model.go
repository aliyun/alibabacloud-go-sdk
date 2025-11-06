// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayServiceCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayServiceCheckResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayServiceCheckResponseBody) *UpdateGatewayServiceCheckResponse
	GetBody() *UpdateGatewayServiceCheckResponseBody
}

type UpdateGatewayServiceCheckResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayServiceCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayServiceCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceCheckResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayServiceCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayServiceCheckResponse) GetBody() *UpdateGatewayServiceCheckResponseBody {
	return s.Body
}

func (s *UpdateGatewayServiceCheckResponse) SetHeaders(v map[string]*string) *UpdateGatewayServiceCheckResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayServiceCheckResponse) SetStatusCode(v int32) *UpdateGatewayServiceCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayServiceCheckResponse) SetBody(v *UpdateGatewayServiceCheckResponseBody) *UpdateGatewayServiceCheckResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayServiceCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
