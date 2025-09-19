// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerFieldStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerFieldStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerFieldStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerFieldStatisticsResponseBody) *DescribeContainerFieldStatisticsResponse
	GetBody() *DescribeContainerFieldStatisticsResponseBody
}

type DescribeContainerFieldStatisticsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerFieldStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerFieldStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerFieldStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerFieldStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerFieldStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerFieldStatisticsResponse) GetBody() *DescribeContainerFieldStatisticsResponseBody {
	return s.Body
}

func (s *DescribeContainerFieldStatisticsResponse) SetHeaders(v map[string]*string) *DescribeContainerFieldStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerFieldStatisticsResponse) SetStatusCode(v int32) *DescribeContainerFieldStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerFieldStatisticsResponse) SetBody(v *DescribeContainerFieldStatisticsResponseBody) *DescribeContainerFieldStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerFieldStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
