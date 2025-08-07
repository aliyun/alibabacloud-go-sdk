// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterMonitorResponseBody) *DescribeDBClusterMonitorResponse
	GetBody() *DescribeDBClusterMonitorResponseBody
}

type DescribeDBClusterMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterMonitorResponse) GetBody() *DescribeDBClusterMonitorResponseBody {
	return s.Body
}

func (s *DescribeDBClusterMonitorResponse) SetHeaders(v map[string]*string) *DescribeDBClusterMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterMonitorResponse) SetStatusCode(v int32) *DescribeDBClusterMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterMonitorResponse) SetBody(v *DescribeDBClusterMonitorResponseBody) *DescribeDBClusterMonitorResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterMonitorResponse) Validate() error {
	return dara.Validate(s)
}
