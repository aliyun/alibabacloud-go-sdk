// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmMonitorAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmMonitorAvailableConfigResponseBody) *DescribeDnsGtmMonitorAvailableConfigResponse
	GetBody() *DescribeDnsGtmMonitorAvailableConfigResponseBody
}

type DescribeDnsGtmMonitorAvailableConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmMonitorAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) GetBody() *DescribeDnsGtmMonitorAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) SetStatusCode(v int32) *DescribeDnsGtmMonitorAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) SetBody(v *DescribeDnsGtmMonitorAvailableConfigResponseBody) *DescribeDnsGtmMonitorAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigResponse) Validate() error {
	return dara.Validate(s)
}
