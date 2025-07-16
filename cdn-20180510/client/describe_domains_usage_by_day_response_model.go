// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsUsageByDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainsUsageByDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainsUsageByDayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainsUsageByDayResponseBody) *DescribeDomainsUsageByDayResponse
	GetBody() *DescribeDomainsUsageByDayResponseBody
}

type DescribeDomainsUsageByDayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsUsageByDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainsUsageByDayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainsUsageByDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainsUsageByDayResponse) GetBody() *DescribeDomainsUsageByDayResponseBody {
	return s.Body
}

func (s *DescribeDomainsUsageByDayResponse) SetHeaders(v map[string]*string) *DescribeDomainsUsageByDayResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) SetStatusCode(v int32) *DescribeDomainsUsageByDayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) SetBody(v *DescribeDomainsUsageByDayResponseBody) *DescribeDomainsUsageByDayResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) Validate() error {
	return dara.Validate(s)
}
