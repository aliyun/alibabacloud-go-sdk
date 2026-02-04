// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouteTableAggregationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouteTableAggregationDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouteTableAggregationDetailResponseBody) *DescribeTransitRouteTableAggregationDetailResponse
	GetBody() *DescribeTransitRouteTableAggregationDetailResponseBody
}

type DescribeTransitRouteTableAggregationDetailResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouteTableAggregationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouteTableAggregationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) GetBody() *DescribeTransitRouteTableAggregationDetailResponseBody {
	return s.Body
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) SetHeaders(v map[string]*string) *DescribeTransitRouteTableAggregationDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) SetStatusCode(v int32) *DescribeTransitRouteTableAggregationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) SetBody(v *DescribeTransitRouteTableAggregationDetailResponseBody) *DescribeTransitRouteTableAggregationDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
