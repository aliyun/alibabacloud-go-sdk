// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateASMGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateASMGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateASMGatewayResponseBody) *UpdateASMGatewayResponse
	GetBody() *UpdateASMGatewayResponseBody
}

type UpdateASMGatewayResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateASMGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateASMGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateASMGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateASMGatewayResponse) GetBody() *UpdateASMGatewayResponseBody {
	return s.Body
}

func (s *UpdateASMGatewayResponse) SetHeaders(v map[string]*string) *UpdateASMGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMGatewayResponse) SetStatusCode(v int32) *UpdateASMGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMGatewayResponse) SetBody(v *UpdateASMGatewayResponseBody) *UpdateASMGatewayResponse {
	s.Body = v
	return s
}

func (s *UpdateASMGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
