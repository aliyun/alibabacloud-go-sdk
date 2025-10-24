// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecStatisticsResponseBody) *DescribeApisecStatisticsResponse
	GetBody() *DescribeApisecStatisticsResponseBody
}

type DescribeApisecStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecStatisticsResponse) GetBody() *DescribeApisecStatisticsResponseBody {
	return s.Body
}

func (s *DescribeApisecStatisticsResponse) SetHeaders(v map[string]*string) *DescribeApisecStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecStatisticsResponse) SetStatusCode(v int32) *DescribeApisecStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecStatisticsResponse) SetBody(v *DescribeApisecStatisticsResponseBody) *DescribeApisecStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
