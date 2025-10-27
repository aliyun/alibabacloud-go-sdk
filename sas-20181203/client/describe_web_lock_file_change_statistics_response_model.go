// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileChangeStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockFileChangeStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockFileChangeStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockFileChangeStatisticsResponseBody) *DescribeWebLockFileChangeStatisticsResponse
	GetBody() *DescribeWebLockFileChangeStatisticsResponseBody
}

type DescribeWebLockFileChangeStatisticsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockFileChangeStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockFileChangeStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileChangeStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileChangeStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockFileChangeStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockFileChangeStatisticsResponse) GetBody() *DescribeWebLockFileChangeStatisticsResponseBody {
	return s.Body
}

func (s *DescribeWebLockFileChangeStatisticsResponse) SetHeaders(v map[string]*string) *DescribeWebLockFileChangeStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponse) SetStatusCode(v int32) *DescribeWebLockFileChangeStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponse) SetBody(v *DescribeWebLockFileChangeStatisticsResponseBody) *DescribeWebLockFileChangeStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
