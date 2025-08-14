// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouteTableAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouteTableAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouteTableAggregationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouteTableAggregationResponseBody) *DeleteTransitRouteTableAggregationResponse
	GetBody() *DeleteTransitRouteTableAggregationResponseBody
}

type DeleteTransitRouteTableAggregationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouteTableAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouteTableAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouteTableAggregationResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouteTableAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouteTableAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouteTableAggregationResponse) GetBody() *DeleteTransitRouteTableAggregationResponseBody {
	return s.Body
}

func (s *DeleteTransitRouteTableAggregationResponse) SetHeaders(v map[string]*string) *DeleteTransitRouteTableAggregationResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouteTableAggregationResponse) SetStatusCode(v int32) *DeleteTransitRouteTableAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouteTableAggregationResponse) SetBody(v *DeleteTransitRouteTableAggregationResponseBody) *DeleteTransitRouteTableAggregationResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouteTableAggregationResponse) Validate() error {
	return dara.Validate(s)
}
