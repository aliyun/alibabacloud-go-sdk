// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouteTableAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouteTableAggregationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouteTableAggregationResponseBody) *DescribeTransitRouteTableAggregationResponse
	GetBody() *DescribeTransitRouteTableAggregationResponseBody
}

type DescribeTransitRouteTableAggregationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouteTableAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouteTableAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouteTableAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouteTableAggregationResponse) GetBody() *DescribeTransitRouteTableAggregationResponseBody {
	return s.Body
}

func (s *DescribeTransitRouteTableAggregationResponse) SetHeaders(v map[string]*string) *DescribeTransitRouteTableAggregationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponse) SetStatusCode(v int32) *DescribeTransitRouteTableAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponse) SetBody(v *DescribeTransitRouteTableAggregationResponseBody) *DescribeTransitRouteTableAggregationResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponse) Validate() error {
	return dara.Validate(s)
}
