// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasDeviceModelStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyPersonasDeviceModelStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyPersonasDeviceModelStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) *DescribeVerifyPersonasDeviceModelStatisticsResponse
	GetBody() *DescribeVerifyPersonasDeviceModelStatisticsResponseBody
}

type DescribeVerifyPersonasDeviceModelStatisticsResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyPersonasDeviceModelStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasDeviceModelStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) GetBody() *DescribeVerifyPersonasDeviceModelStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyPersonasDeviceModelStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyPersonasDeviceModelStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) SetBody(v *DescribeVerifyPersonasDeviceModelStatisticsResponseBody) *DescribeVerifyPersonasDeviceModelStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
