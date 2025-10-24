// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayResponseBody) *UpdateGatewayResponse
	GetBody() *UpdateGatewayResponseBody
}

type UpdateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayResponse) GetBody() *UpdateGatewayResponseBody {
	return s.Body
}

func (s *UpdateGatewayResponse) SetHeaders(v map[string]*string) *UpdateGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayResponse) SetStatusCode(v int32) *UpdateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayResponse) SetBody(v *UpdateGatewayResponseBody) *UpdateGatewayResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
