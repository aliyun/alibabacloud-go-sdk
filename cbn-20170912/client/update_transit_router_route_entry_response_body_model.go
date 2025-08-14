// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterRouteEntryResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterRouteEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D69CCEA-42D0-48B2-8C9A-9BB207F76D6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterRouteEntryResponseBody) SetRequestId(v string) *UpdateTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
