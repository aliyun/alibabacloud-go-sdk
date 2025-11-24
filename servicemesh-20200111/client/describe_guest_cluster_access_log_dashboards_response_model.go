// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterAccessLogDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGuestClusterAccessLogDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGuestClusterAccessLogDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGuestClusterAccessLogDashboardsResponseBody) *DescribeGuestClusterAccessLogDashboardsResponse
	GetBody() *DescribeGuestClusterAccessLogDashboardsResponseBody
}

type DescribeGuestClusterAccessLogDashboardsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGuestClusterAccessLogDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGuestClusterAccessLogDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterAccessLogDashboardsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) GetBody() *DescribeGuestClusterAccessLogDashboardsResponseBody {
	return s.Body
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetStatusCode(v int32) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) SetBody(v *DescribeGuestClusterAccessLogDashboardsResponseBody) *DescribeGuestClusterAccessLogDashboardsResponse {
	s.Body = v
	return s
}

func (s *DescribeGuestClusterAccessLogDashboardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
