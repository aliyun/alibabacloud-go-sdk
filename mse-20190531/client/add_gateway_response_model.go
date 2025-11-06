// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayResponseBody) *AddGatewayResponse
	GetBody() *AddGatewayResponseBody
}

type AddGatewayResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayResponse) GetBody() *AddGatewayResponseBody {
	return s.Body
}

func (s *AddGatewayResponse) SetHeaders(v map[string]*string) *AddGatewayResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayResponse) SetStatusCode(v int32) *AddGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayResponse) SetBody(v *AddGatewayResponseBody) *AddGatewayResponse {
	s.Body = v
	return s
}

func (s *AddGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
