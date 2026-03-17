// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserOnlineClientStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserOnlineClientStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserOnlineClientStatisticsResponseBody) *DescribeUserOnlineClientStatisticsResponse
	GetBody() *DescribeUserOnlineClientStatisticsResponseBody
}

type DescribeUserOnlineClientStatisticsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserOnlineClientStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserOnlineClientStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserOnlineClientStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserOnlineClientStatisticsResponse) GetBody() *DescribeUserOnlineClientStatisticsResponseBody {
	return s.Body
}

func (s *DescribeUserOnlineClientStatisticsResponse) SetHeaders(v map[string]*string) *DescribeUserOnlineClientStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponse) SetStatusCode(v int32) *DescribeUserOnlineClientStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponse) SetBody(v *DescribeUserOnlineClientStatisticsResponseBody) *DescribeUserOnlineClientStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserOnlineClientStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
