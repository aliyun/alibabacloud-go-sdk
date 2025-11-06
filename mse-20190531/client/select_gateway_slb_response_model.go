// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectGatewaySlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectGatewaySlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectGatewaySlbResponse
	GetStatusCode() *int32
	SetBody(v *SelectGatewaySlbResponseBody) *SelectGatewaySlbResponse
	GetBody() *SelectGatewaySlbResponseBody
}

type SelectGatewaySlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectGatewaySlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectGatewaySlbResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectGatewaySlbResponse) GoString() string {
	return s.String()
}

func (s *SelectGatewaySlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectGatewaySlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectGatewaySlbResponse) GetBody() *SelectGatewaySlbResponseBody {
	return s.Body
}

func (s *SelectGatewaySlbResponse) SetHeaders(v map[string]*string) *SelectGatewaySlbResponse {
	s.Headers = v
	return s
}

func (s *SelectGatewaySlbResponse) SetStatusCode(v int32) *SelectGatewaySlbResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectGatewaySlbResponse) SetBody(v *SelectGatewaySlbResponseBody) *SelectGatewaySlbResponse {
	s.Body = v
	return s
}

func (s *SelectGatewaySlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
