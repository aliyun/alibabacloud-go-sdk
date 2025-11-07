// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasProvinceStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyPersonasProvinceStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyPersonasProvinceStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyPersonasProvinceStatisticsResponseBody) *DescribeVerifyPersonasProvinceStatisticsResponse
	GetBody() *DescribeVerifyPersonasProvinceStatisticsResponseBody
}

type DescribeVerifyPersonasProvinceStatisticsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyPersonasProvinceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyPersonasProvinceStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasProvinceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) GetBody() *DescribeVerifyPersonasProvinceStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyPersonasProvinceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyPersonasProvinceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) SetBody(v *DescribeVerifyPersonasProvinceStatisticsResponseBody) *DescribeVerifyPersonasProvinceStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
