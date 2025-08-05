// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelQuotaTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTunnelQuotaTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTunnelQuotaTimerResponse
	GetStatusCode() *int32
	SetBody(v *ListTunnelQuotaTimerResponseBody) *ListTunnelQuotaTimerResponse
	GetBody() *ListTunnelQuotaTimerResponseBody
}

type ListTunnelQuotaTimerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTunnelQuotaTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTunnelQuotaTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelQuotaTimerResponse) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTunnelQuotaTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTunnelQuotaTimerResponse) GetBody() *ListTunnelQuotaTimerResponseBody {
	return s.Body
}

func (s *ListTunnelQuotaTimerResponse) SetHeaders(v map[string]*string) *ListTunnelQuotaTimerResponse {
	s.Headers = v
	return s
}

func (s *ListTunnelQuotaTimerResponse) SetStatusCode(v int32) *ListTunnelQuotaTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponse) SetBody(v *ListTunnelQuotaTimerResponseBody) *ListTunnelQuotaTimerResponse {
	s.Body = v
	return s
}

func (s *ListTunnelQuotaTimerResponse) Validate() error {
	return dara.Validate(s)
}
