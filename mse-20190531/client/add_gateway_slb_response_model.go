// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewaySlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewaySlbResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewaySlbResponseBody) *AddGatewaySlbResponse
	GetBody() *AddGatewaySlbResponseBody
}

type AddGatewaySlbResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewaySlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewaySlbResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySlbResponse) GoString() string {
	return s.String()
}

func (s *AddGatewaySlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewaySlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewaySlbResponse) GetBody() *AddGatewaySlbResponseBody {
	return s.Body
}

func (s *AddGatewaySlbResponse) SetHeaders(v map[string]*string) *AddGatewaySlbResponse {
	s.Headers = v
	return s
}

func (s *AddGatewaySlbResponse) SetStatusCode(v int32) *AddGatewaySlbResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewaySlbResponse) SetBody(v *AddGatewaySlbResponseBody) *AddGatewaySlbResponse {
	s.Body = v
	return s
}

func (s *AddGatewaySlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
