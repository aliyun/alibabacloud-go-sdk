// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyDeviceRiskStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyDeviceRiskStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyDeviceRiskStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyDeviceRiskStatisticsResponseBody) *DescribeVerifyDeviceRiskStatisticsResponse
	GetBody() *DescribeVerifyDeviceRiskStatisticsResponseBody
}

type DescribeVerifyDeviceRiskStatisticsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyDeviceRiskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyDeviceRiskStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyDeviceRiskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) GetBody() *DescribeVerifyDeviceRiskStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyDeviceRiskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyDeviceRiskStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) SetBody(v *DescribeVerifyDeviceRiskStatisticsResponseBody) *DescribeVerifyDeviceRiskStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
