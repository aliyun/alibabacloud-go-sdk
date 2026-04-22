// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsZoneOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsZoneOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInterAuthStatisticsZoneOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInterAuthStatisticsZoneOverviewResponseBody) *DescribeInterAuthStatisticsZoneOverviewResponse
	GetBody() *DescribeInterAuthStatisticsZoneOverviewResponseBody
}

type DescribeInterAuthStatisticsZoneOverviewResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterAuthStatisticsZoneOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInterAuthStatisticsZoneOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsZoneOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) GetBody() *DescribeInterAuthStatisticsZoneOverviewResponseBody {
	return s.Body
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsZoneOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) SetStatusCode(v int32) *DescribeInterAuthStatisticsZoneOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) SetBody(v *DescribeInterAuthStatisticsZoneOverviewResponseBody) *DescribeInterAuthStatisticsZoneOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
