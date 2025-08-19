// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartStatisticsPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartStatisticsPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartStatisticsPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartStatisticsPageListResponseBody) *DescribeSmartStatisticsPageListResponse
	GetBody() *DescribeSmartStatisticsPageListResponseBody
}

type DescribeSmartStatisticsPageListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartStatisticsPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartStatisticsPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartStatisticsPageListResponse) GetBody() *DescribeSmartStatisticsPageListResponseBody {
	return s.Body
}

func (s *DescribeSmartStatisticsPageListResponse) SetHeaders(v map[string]*string) *DescribeSmartStatisticsPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartStatisticsPageListResponse) SetStatusCode(v int32) *DescribeSmartStatisticsPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponse) SetBody(v *DescribeSmartStatisticsPageListResponseBody) *DescribeSmartStatisticsPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartStatisticsPageListResponse) Validate() error {
	return dara.Validate(s)
}
