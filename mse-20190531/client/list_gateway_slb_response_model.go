// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewaySlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewaySlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewaySlbResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewaySlbResponseBody) *ListGatewaySlbResponse
	GetBody() *ListGatewaySlbResponseBody
}

type ListGatewaySlbResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewaySlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewaySlbResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewaySlbResponse) GoString() string {
	return s.String()
}

func (s *ListGatewaySlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewaySlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewaySlbResponse) GetBody() *ListGatewaySlbResponseBody {
	return s.Body
}

func (s *ListGatewaySlbResponse) SetHeaders(v map[string]*string) *ListGatewaySlbResponse {
	s.Headers = v
	return s
}

func (s *ListGatewaySlbResponse) SetStatusCode(v int32) *ListGatewaySlbResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewaySlbResponse) SetBody(v *ListGatewaySlbResponseBody) *ListGatewaySlbResponse {
	s.Body = v
	return s
}

func (s *ListGatewaySlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
