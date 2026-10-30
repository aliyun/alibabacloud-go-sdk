// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCapacityReservationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTransitRouterCapacityReservationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTransitRouterCapacityReservationsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTransitRouterCapacityReservationsResponseBody) *ModifyTransitRouterCapacityReservationsResponse
	GetBody() *ModifyTransitRouterCapacityReservationsResponseBody
}

type ModifyTransitRouterCapacityReservationsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTransitRouterCapacityReservationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTransitRouterCapacityReservationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCapacityReservationsResponse) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCapacityReservationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTransitRouterCapacityReservationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTransitRouterCapacityReservationsResponse) GetBody() *ModifyTransitRouterCapacityReservationsResponseBody {
	return s.Body
}

func (s *ModifyTransitRouterCapacityReservationsResponse) SetHeaders(v map[string]*string) *ModifyTransitRouterCapacityReservationsResponse {
	s.Headers = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsResponse) SetStatusCode(v int32) *ModifyTransitRouterCapacityReservationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsResponse) SetBody(v *ModifyTransitRouterCapacityReservationsResponseBody) *ModifyTransitRouterCapacityReservationsResponse {
	s.Body = v
	return s
}

func (s *ModifyTransitRouterCapacityReservationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
