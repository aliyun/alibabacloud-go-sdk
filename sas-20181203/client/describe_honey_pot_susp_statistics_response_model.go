// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotSuspStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHoneyPotSuspStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHoneyPotSuspStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHoneyPotSuspStatisticsResponseBody) *DescribeHoneyPotSuspStatisticsResponse
	GetBody() *DescribeHoneyPotSuspStatisticsResponseBody
}

type DescribeHoneyPotSuspStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHoneyPotSuspStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHoneyPotSuspStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHoneyPotSuspStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHoneyPotSuspStatisticsResponse) GetBody() *DescribeHoneyPotSuspStatisticsResponseBody {
	return s.Body
}

func (s *DescribeHoneyPotSuspStatisticsResponse) SetHeaders(v map[string]*string) *DescribeHoneyPotSuspStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponse) SetStatusCode(v int32) *DescribeHoneyPotSuspStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponse) SetBody(v *DescribeHoneyPotSuspStatisticsResponseBody) *DescribeHoneyPotSuspStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
