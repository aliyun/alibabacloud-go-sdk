// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetaStatisticsPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetaStatisticsPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetaStatisticsPageListResponseBody) *DescribeMetaStatisticsPageListResponse
	GetBody() *DescribeMetaStatisticsPageListResponseBody
}

type DescribeMetaStatisticsPageListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetaStatisticsPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetaStatisticsPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetaStatisticsPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetaStatisticsPageListResponse) GetBody() *DescribeMetaStatisticsPageListResponseBody {
	return s.Body
}

func (s *DescribeMetaStatisticsPageListResponse) SetHeaders(v map[string]*string) *DescribeMetaStatisticsPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaStatisticsPageListResponse) SetStatusCode(v int32) *DescribeMetaStatisticsPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponse) SetBody(v *DescribeMetaStatisticsPageListResponseBody) *DescribeMetaStatisticsPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetaStatisticsPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
