// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteDistributionStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteDistributionStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteDistributionStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteDistributionStrategiesResponseBody) *DescribeRouteDistributionStrategiesResponse
	GetBody() *DescribeRouteDistributionStrategiesResponseBody
}

type DescribeRouteDistributionStrategiesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteDistributionStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteDistributionStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteDistributionStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteDistributionStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteDistributionStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteDistributionStrategiesResponse) GetBody() *DescribeRouteDistributionStrategiesResponseBody {
	return s.Body
}

func (s *DescribeRouteDistributionStrategiesResponse) SetHeaders(v map[string]*string) *DescribeRouteDistributionStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponse) SetStatusCode(v int32) *DescribeRouteDistributionStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponse) SetBody(v *DescribeRouteDistributionStrategiesResponseBody) *DescribeRouteDistributionStrategiesResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteDistributionStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
