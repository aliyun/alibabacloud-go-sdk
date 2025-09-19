// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerStatisticsResponseBody) *DescribeContainerStatisticsResponse
	GetBody() *DescribeContainerStatisticsResponseBody
}

type DescribeContainerStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerStatisticsResponse) GetBody() *DescribeContainerStatisticsResponseBody {
	return s.Body
}

func (s *DescribeContainerStatisticsResponse) SetHeaders(v map[string]*string) *DescribeContainerStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerStatisticsResponse) SetStatusCode(v int32) *DescribeContainerStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerStatisticsResponse) SetBody(v *DescribeContainerStatisticsResponseBody) *DescribeContainerStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
