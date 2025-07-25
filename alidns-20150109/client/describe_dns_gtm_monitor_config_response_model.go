// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmMonitorConfigResponseBody) *DescribeDnsGtmMonitorConfigResponse
	GetBody() *DescribeDnsGtmMonitorConfigResponseBody
}

type DescribeDnsGtmMonitorConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmMonitorConfigResponse) GetBody() *DescribeDnsGtmMonitorConfigResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmMonitorConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponse) SetStatusCode(v int32) *DescribeDnsGtmMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponse) SetBody(v *DescribeDnsGtmMonitorConfigResponseBody) *DescribeDnsGtmMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponse) Validate() error {
	return dara.Validate(s)
}
