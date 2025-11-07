// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasOsStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyPersonasOsStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyPersonasOsStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyPersonasOsStatisticsResponseBody) *DescribeVerifyPersonasOsStatisticsResponse
	GetBody() *DescribeVerifyPersonasOsStatisticsResponseBody
}

type DescribeVerifyPersonasOsStatisticsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyPersonasOsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyPersonasOsStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasOsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) GetBody() *DescribeVerifyPersonasOsStatisticsResponseBody {
	return s.Body
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) SetHeaders(v map[string]*string) *DescribeVerifyPersonasOsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) SetStatusCode(v int32) *DescribeVerifyPersonasOsStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) SetBody(v *DescribeVerifyPersonasOsStatisticsResponseBody) *DescribeVerifyPersonasOsStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
