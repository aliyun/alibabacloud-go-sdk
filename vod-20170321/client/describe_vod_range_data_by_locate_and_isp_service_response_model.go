// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRangeDataByLocateAndIspServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodRangeDataByLocateAndIspServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodRangeDataByLocateAndIspServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodRangeDataByLocateAndIspServiceResponseBody) *DescribeVodRangeDataByLocateAndIspServiceResponse
	GetBody() *DescribeVodRangeDataByLocateAndIspServiceResponseBody
}

type DescribeVodRangeDataByLocateAndIspServiceResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodRangeDataByLocateAndIspServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodRangeDataByLocateAndIspServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRangeDataByLocateAndIspServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) GetBody() *DescribeVodRangeDataByLocateAndIspServiceResponseBody {
	return s.Body
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) SetHeaders(v map[string]*string) *DescribeVodRangeDataByLocateAndIspServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) SetStatusCode(v int32) *DescribeVodRangeDataByLocateAndIspServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) SetBody(v *DescribeVodRangeDataByLocateAndIspServiceResponseBody) *DescribeVodRangeDataByLocateAndIspServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponse) Validate() error {
	return dara.Validate(s)
}
