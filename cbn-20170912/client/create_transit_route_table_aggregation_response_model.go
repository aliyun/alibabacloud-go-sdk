// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouteTableAggregationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouteTableAggregationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouteTableAggregationResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouteTableAggregationResponseBody) *CreateTransitRouteTableAggregationResponse
	GetBody() *CreateTransitRouteTableAggregationResponseBody
}

type CreateTransitRouteTableAggregationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouteTableAggregationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouteTableAggregationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouteTableAggregationResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouteTableAggregationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouteTableAggregationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouteTableAggregationResponse) GetBody() *CreateTransitRouteTableAggregationResponseBody {
	return s.Body
}

func (s *CreateTransitRouteTableAggregationResponse) SetHeaders(v map[string]*string) *CreateTransitRouteTableAggregationResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouteTableAggregationResponse) SetStatusCode(v int32) *CreateTransitRouteTableAggregationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouteTableAggregationResponse) SetBody(v *CreateTransitRouteTableAggregationResponseBody) *CreateTransitRouteTableAggregationResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouteTableAggregationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
