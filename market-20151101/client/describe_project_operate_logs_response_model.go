// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectOperateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectOperateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectOperateLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectOperateLogsResponseBody) *DescribeProjectOperateLogsResponse
	GetBody() *DescribeProjectOperateLogsResponseBody
}

type DescribeProjectOperateLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectOperateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectOperateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectOperateLogsResponse) GetBody() *DescribeProjectOperateLogsResponseBody {
	return s.Body
}

func (s *DescribeProjectOperateLogsResponse) SetHeaders(v map[string]*string) *DescribeProjectOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectOperateLogsResponse) SetStatusCode(v int32) *DescribeProjectOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectOperateLogsResponse) SetBody(v *DescribeProjectOperateLogsResponseBody) *DescribeProjectOperateLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectOperateLogsResponse) Validate() error {
	return dara.Validate(s)
}
