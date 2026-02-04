// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshTransitRouteTableAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshTransitRouteTableAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshTransitRouteTableAggregationResponse
	GetStatusCode() *int32
	SetBody(v *RefreshTransitRouteTableAggregationResponseBody) *RefreshTransitRouteTableAggregationResponse
	GetBody() *RefreshTransitRouteTableAggregationResponseBody
}

type RefreshTransitRouteTableAggregationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshTransitRouteTableAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshTransitRouteTableAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshTransitRouteTableAggregationResponse) GoString() string {
	return s.String()
}

func (s *RefreshTransitRouteTableAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshTransitRouteTableAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshTransitRouteTableAggregationResponse) GetBody() *RefreshTransitRouteTableAggregationResponseBody {
	return s.Body
}

func (s *RefreshTransitRouteTableAggregationResponse) SetHeaders(v map[string]*string) *RefreshTransitRouteTableAggregationResponse {
	s.Headers = v
	return s
}

func (s *RefreshTransitRouteTableAggregationResponse) SetStatusCode(v int32) *RefreshTransitRouteTableAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshTransitRouteTableAggregationResponse) SetBody(v *RefreshTransitRouteTableAggregationResponseBody) *RefreshTransitRouteTableAggregationResponse {
	s.Body = v
	return s
}

func (s *RefreshTransitRouteTableAggregationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
