// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouteTableAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTransitRouteTableAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTransitRouteTableAggregationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTransitRouteTableAggregationResponseBody) *ModifyTransitRouteTableAggregationResponse
	GetBody() *ModifyTransitRouteTableAggregationResponseBody
}

type ModifyTransitRouteTableAggregationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTransitRouteTableAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTransitRouteTableAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouteTableAggregationResponse) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouteTableAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTransitRouteTableAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTransitRouteTableAggregationResponse) GetBody() *ModifyTransitRouteTableAggregationResponseBody {
	return s.Body
}

func (s *ModifyTransitRouteTableAggregationResponse) SetHeaders(v map[string]*string) *ModifyTransitRouteTableAggregationResponse {
	s.Headers = v
	return s
}

func (s *ModifyTransitRouteTableAggregationResponse) SetStatusCode(v int32) *ModifyTransitRouteTableAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationResponse) SetBody(v *ModifyTransitRouteTableAggregationResponseBody) *ModifyTransitRouteTableAggregationResponse {
	s.Body = v
	return s
}

func (s *ModifyTransitRouteTableAggregationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
