// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowQueryStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowQueryStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowQueryStatsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowQueryStatsResponseBody) *DescribeSlowQueryStatsResponse
	GetBody() *DescribeSlowQueryStatsResponseBody
}

type DescribeSlowQueryStatsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowQueryStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowQueryStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowQueryStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowQueryStatsResponse) GetBody() *DescribeSlowQueryStatsResponseBody {
	return s.Body
}

func (s *DescribeSlowQueryStatsResponse) SetHeaders(v map[string]*string) *DescribeSlowQueryStatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowQueryStatsResponse) SetStatusCode(v int32) *DescribeSlowQueryStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowQueryStatsResponse) SetBody(v *DescribeSlowQueryStatsResponseBody) *DescribeSlowQueryStatsResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowQueryStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
