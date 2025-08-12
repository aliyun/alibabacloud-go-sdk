// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogMonitorListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogMonitorListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogMonitorListResponseBody) *DescribeLogMonitorListResponse
	GetBody() *DescribeLogMonitorListResponseBody
}

type DescribeLogMonitorListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogMonitorListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogMonitorListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogMonitorListResponse) GetBody() *DescribeLogMonitorListResponseBody {
	return s.Body
}

func (s *DescribeLogMonitorListResponse) SetHeaders(v map[string]*string) *DescribeLogMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogMonitorListResponse) SetStatusCode(v int32) *DescribeLogMonitorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogMonitorListResponse) SetBody(v *DescribeLogMonitorListResponseBody) *DescribeLogMonitorListResponse {
	s.Body = v
	return s
}

func (s *DescribeLogMonitorListResponse) Validate() error {
	return dara.Validate(s)
}
