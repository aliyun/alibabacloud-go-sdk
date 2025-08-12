// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupsResponseBody) *DescribeMonitorGroupsResponse
	GetBody() *DescribeMonitorGroupsResponseBody
}

type DescribeMonitorGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupsResponse) GetBody() *DescribeMonitorGroupsResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupsResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupsResponse) SetStatusCode(v int32) *DescribeMonitorGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupsResponse) SetBody(v *DescribeMonitorGroupsResponseBody) *DescribeMonitorGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupsResponse) Validate() error {
	return dara.Validate(s)
}
