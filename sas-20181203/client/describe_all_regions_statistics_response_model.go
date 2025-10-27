// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllRegionsStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllRegionsStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllRegionsStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllRegionsStatisticsResponseBody) *DescribeAllRegionsStatisticsResponse
	GetBody() *DescribeAllRegionsStatisticsResponseBody
}

type DescribeAllRegionsStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllRegionsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllRegionsStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRegionsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllRegionsStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllRegionsStatisticsResponse) GetBody() *DescribeAllRegionsStatisticsResponseBody {
	return s.Body
}

func (s *DescribeAllRegionsStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAllRegionsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllRegionsStatisticsResponse) SetStatusCode(v int32) *DescribeAllRegionsStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponse) SetBody(v *DescribeAllRegionsStatisticsResponseBody) *DescribeAllRegionsStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeAllRegionsStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
