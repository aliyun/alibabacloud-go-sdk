// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmMonitorStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDnsGtmMonitorStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDnsGtmMonitorStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDnsGtmMonitorStatusResponseBody) *SetDnsGtmMonitorStatusResponse
	GetBody() *SetDnsGtmMonitorStatusResponseBody
}

type SetDnsGtmMonitorStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDnsGtmMonitorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDnsGtmMonitorStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmMonitorStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDnsGtmMonitorStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDnsGtmMonitorStatusResponse) GetBody() *SetDnsGtmMonitorStatusResponseBody {
	return s.Body
}

func (s *SetDnsGtmMonitorStatusResponse) SetHeaders(v map[string]*string) *SetDnsGtmMonitorStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDnsGtmMonitorStatusResponse) SetStatusCode(v int32) *SetDnsGtmMonitorStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDnsGtmMonitorStatusResponse) SetBody(v *SetDnsGtmMonitorStatusResponseBody) *SetDnsGtmMonitorStatusResponse {
	s.Body = v
	return s
}

func (s *SetDnsGtmMonitorStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
