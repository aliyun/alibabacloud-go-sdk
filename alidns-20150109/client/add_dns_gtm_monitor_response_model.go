// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnsGtmMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnsGtmMonitorResponse
	GetStatusCode() *int32
	SetBody(v *AddDnsGtmMonitorResponseBody) *AddDnsGtmMonitorResponse
	GetBody() *AddDnsGtmMonitorResponseBody
}

type AddDnsGtmMonitorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnsGtmMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnsGtmMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmMonitorResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnsGtmMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnsGtmMonitorResponse) GetBody() *AddDnsGtmMonitorResponseBody {
	return s.Body
}

func (s *AddDnsGtmMonitorResponse) SetHeaders(v map[string]*string) *AddDnsGtmMonitorResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmMonitorResponse) SetStatusCode(v int32) *AddDnsGtmMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnsGtmMonitorResponse) SetBody(v *AddDnsGtmMonitorResponseBody) *AddDnsGtmMonitorResponse {
	s.Body = v
	return s
}

func (s *AddDnsGtmMonitorResponse) Validate() error {
	return dara.Validate(s)
}
