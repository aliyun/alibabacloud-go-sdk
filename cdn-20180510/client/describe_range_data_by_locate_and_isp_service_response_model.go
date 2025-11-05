// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRangeDataByLocateAndIspServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRangeDataByLocateAndIspServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRangeDataByLocateAndIspServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRangeDataByLocateAndIspServiceResponseBody) *DescribeRangeDataByLocateAndIspServiceResponse
	GetBody() *DescribeRangeDataByLocateAndIspServiceResponseBody
}

type DescribeRangeDataByLocateAndIspServiceResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRangeDataByLocateAndIspServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRangeDataByLocateAndIspServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) GetBody() *DescribeRangeDataByLocateAndIspServiceResponseBody {
	return s.Body
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) SetHeaders(v map[string]*string) *DescribeRangeDataByLocateAndIspServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) SetStatusCode(v int32) *DescribeRangeDataByLocateAndIspServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) SetBody(v *DescribeRangeDataByLocateAndIspServiceResponseBody) *DescribeRangeDataByLocateAndIspServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
