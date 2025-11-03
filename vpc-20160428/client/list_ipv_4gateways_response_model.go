// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpv4GatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpv4GatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpv4GatewaysResponse
	GetStatusCode() *int32
	SetBody(v *ListIpv4GatewaysResponseBody) *ListIpv4GatewaysResponse
	GetBody() *ListIpv4GatewaysResponseBody
}

type ListIpv4GatewaysResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpv4GatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpv4GatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpv4GatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpv4GatewaysResponse) GetBody() *ListIpv4GatewaysResponseBody {
	return s.Body
}

func (s *ListIpv4GatewaysResponse) SetHeaders(v map[string]*string) *ListIpv4GatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListIpv4GatewaysResponse) SetStatusCode(v int32) *ListIpv4GatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpv4GatewaysResponse) SetBody(v *ListIpv4GatewaysResponseBody) *ListIpv4GatewaysResponse {
	s.Body = v
	return s
}

func (s *ListIpv4GatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
