// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetaStatisticsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetaStatisticsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetaStatisticsListResponseBody) *DescribeMetaStatisticsListResponse
	GetBody() *DescribeMetaStatisticsListResponseBody
}

type DescribeMetaStatisticsListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetaStatisticsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetaStatisticsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetaStatisticsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetaStatisticsListResponse) GetBody() *DescribeMetaStatisticsListResponseBody {
	return s.Body
}

func (s *DescribeMetaStatisticsListResponse) SetHeaders(v map[string]*string) *DescribeMetaStatisticsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaStatisticsListResponse) SetStatusCode(v int32) *DescribeMetaStatisticsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetaStatisticsListResponse) SetBody(v *DescribeMetaStatisticsListResponseBody) *DescribeMetaStatisticsListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetaStatisticsListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
