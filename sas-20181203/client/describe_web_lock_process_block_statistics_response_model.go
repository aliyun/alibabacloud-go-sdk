// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessBlockStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockProcessBlockStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockProcessBlockStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockProcessBlockStatisticsResponseBody) *DescribeWebLockProcessBlockStatisticsResponse
	GetBody() *DescribeWebLockProcessBlockStatisticsResponseBody
}

type DescribeWebLockProcessBlockStatisticsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockProcessBlockStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockProcessBlockStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessBlockStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) GetBody() *DescribeWebLockProcessBlockStatisticsResponseBody {
	return s.Body
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) SetHeaders(v map[string]*string) *DescribeWebLockProcessBlockStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) SetStatusCode(v int32) *DescribeWebLockProcessBlockStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) SetBody(v *DescribeWebLockProcessBlockStatisticsResponseBody) *DescribeWebLockProcessBlockStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
