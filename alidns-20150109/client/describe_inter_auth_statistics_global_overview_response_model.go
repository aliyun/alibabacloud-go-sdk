// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsGlobalOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsGlobalOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInterAuthStatisticsGlobalOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInterAuthStatisticsGlobalOverviewResponseBody) *DescribeInterAuthStatisticsGlobalOverviewResponse
	GetBody() *DescribeInterAuthStatisticsGlobalOverviewResponseBody
}

type DescribeInterAuthStatisticsGlobalOverviewResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterAuthStatisticsGlobalOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsGlobalOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) GetBody() *DescribeInterAuthStatisticsGlobalOverviewResponseBody {
	return s.Body
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsGlobalOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) SetStatusCode(v int32) *DescribeInterAuthStatisticsGlobalOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) SetBody(v *DescribeInterAuthStatisticsGlobalOverviewResponseBody) *DescribeInterAuthStatisticsGlobalOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeInterAuthStatisticsGlobalOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
