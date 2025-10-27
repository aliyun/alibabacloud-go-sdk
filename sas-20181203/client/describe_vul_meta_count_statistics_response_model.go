// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulMetaCountStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulMetaCountStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulMetaCountStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulMetaCountStatisticsResponseBody) *DescribeVulMetaCountStatisticsResponse
	GetBody() *DescribeVulMetaCountStatisticsResponseBody
}

type DescribeVulMetaCountStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulMetaCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulMetaCountStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulMetaCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulMetaCountStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulMetaCountStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulMetaCountStatisticsResponse) GetBody() *DescribeVulMetaCountStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVulMetaCountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVulMetaCountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponse) SetStatusCode(v int32) *DescribeVulMetaCountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponse) SetBody(v *DescribeVulMetaCountStatisticsResponseBody) *DescribeVulMetaCountStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
