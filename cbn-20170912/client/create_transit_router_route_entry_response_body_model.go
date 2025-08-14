// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterRouteEntryResponseBody
	GetRequestId() *string
	SetTransitRouterRouteEntryId(v string) *CreateTransitRouterRouteEntryResponseBody
	GetTransitRouterRouteEntryId() *string
}

type CreateTransitRouterRouteEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 835E7F4B-B380-4E0F-96A5-6EA572388047
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// rte-75eg4jprkvk0pw****
	TransitRouterRouteEntryId *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
}

func (s CreateTransitRouterRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterRouteEntryResponseBody) GetTransitRouterRouteEntryId() *string {
	return s.TransitRouterRouteEntryId
}

func (s *CreateTransitRouterRouteEntryResponseBody) SetRequestId(v string) *CreateTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryResponseBody) SetTransitRouterRouteEntryId(v string) *CreateTransitRouterRouteEntryResponseBody {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
