// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainsUsageByDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainsUsageByDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainsUsageByDayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainsUsageByDayResponseBody) *DescribeVodDomainsUsageByDayResponse
	GetBody() *DescribeVodDomainsUsageByDayResponseBody
}

type DescribeVodDomainsUsageByDayResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainsUsageByDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainsUsageByDayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainsUsageByDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainsUsageByDayResponse) GetBody() *DescribeVodDomainsUsageByDayResponseBody {
	return s.Body
}

func (s *DescribeVodDomainsUsageByDayResponse) SetHeaders(v map[string]*string) *DescribeVodDomainsUsageByDayResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponse) SetStatusCode(v int32) *DescribeVodDomainsUsageByDayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponse) SetBody(v *DescribeVodDomainsUsageByDayResponseBody) *DescribeVodDomainsUsageByDayResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponse) Validate() error {
	return dara.Validate(s)
}
