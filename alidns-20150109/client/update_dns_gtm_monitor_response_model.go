// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsGtmMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsGtmMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsGtmMonitorResponseBody) *UpdateDnsGtmMonitorResponse
	GetBody() *UpdateDnsGtmMonitorResponseBody
}

type UpdateDnsGtmMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsGtmMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsGtmMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsGtmMonitorResponse) GetBody() *UpdateDnsGtmMonitorResponseBody {
	return s.Body
}

func (s *UpdateDnsGtmMonitorResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmMonitorResponse) SetStatusCode(v int32) *UpdateDnsGtmMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsGtmMonitorResponse) SetBody(v *UpdateDnsGtmMonitorResponseBody) *UpdateDnsGtmMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsGtmMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
