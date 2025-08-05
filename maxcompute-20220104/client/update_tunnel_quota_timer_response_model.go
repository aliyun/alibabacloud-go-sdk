// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelQuotaTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTunnelQuotaTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTunnelQuotaTimerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTunnelQuotaTimerResponseBody) *UpdateTunnelQuotaTimerResponse
	GetBody() *UpdateTunnelQuotaTimerResponseBody
}

type UpdateTunnelQuotaTimerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTunnelQuotaTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTunnelQuotaTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelQuotaTimerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTunnelQuotaTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTunnelQuotaTimerResponse) GetBody() *UpdateTunnelQuotaTimerResponseBody {
	return s.Body
}

func (s *UpdateTunnelQuotaTimerResponse) SetHeaders(v map[string]*string) *UpdateTunnelQuotaTimerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTunnelQuotaTimerResponse) SetStatusCode(v int32) *UpdateTunnelQuotaTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTunnelQuotaTimerResponse) SetBody(v *UpdateTunnelQuotaTimerResponseBody) *UpdateTunnelQuotaTimerResponse {
	s.Body = v
	return s
}

func (s *UpdateTunnelQuotaTimerResponse) Validate() error {
	return dara.Validate(s)
}
