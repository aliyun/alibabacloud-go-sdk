// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTunnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTunnelResponse
	GetStatusCode() *int32
	SetBody(v *ListTunnelResponseBody) *ListTunnelResponse
	GetBody() *ListTunnelResponseBody
}

type ListTunnelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTunnelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelResponse) GoString() string {
	return s.String()
}

func (s *ListTunnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTunnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTunnelResponse) GetBody() *ListTunnelResponseBody {
	return s.Body
}

func (s *ListTunnelResponse) SetHeaders(v map[string]*string) *ListTunnelResponse {
	s.Headers = v
	return s
}

func (s *ListTunnelResponse) SetStatusCode(v int32) *ListTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTunnelResponse) SetBody(v *ListTunnelResponseBody) *ListTunnelResponse {
	s.Body = v
	return s
}

func (s *ListTunnelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
