// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterOperateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterOperateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterOperateLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterOperateLogsResponseBody) *DescribeClusterOperateLogsResponse
	GetBody() *DescribeClusterOperateLogsResponseBody
}

type DescribeClusterOperateLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterOperateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterOperateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterOperateLogsResponse) GetBody() *DescribeClusterOperateLogsResponseBody {
	return s.Body
}

func (s *DescribeClusterOperateLogsResponse) SetHeaders(v map[string]*string) *DescribeClusterOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterOperateLogsResponse) SetStatusCode(v int32) *DescribeClusterOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterOperateLogsResponse) SetBody(v *DescribeClusterOperateLogsResponseBody) *DescribeClusterOperateLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterOperateLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
