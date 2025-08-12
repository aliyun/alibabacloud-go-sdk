// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupInstancesResponseBody) *DescribeMonitorGroupInstancesResponse
	GetBody() *DescribeMonitorGroupInstancesResponseBody
}

type DescribeMonitorGroupInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupInstancesResponse) GetBody() *DescribeMonitorGroupInstancesResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponse) SetStatusCode(v int32) *DescribeMonitorGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponse) SetBody(v *DescribeMonitorGroupInstancesResponseBody) *DescribeMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponse) Validate() error {
	return dara.Validate(s)
}
