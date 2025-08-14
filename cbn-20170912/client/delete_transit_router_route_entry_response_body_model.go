// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterRouteEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D69CCEA-42D0-48B2-8C9A-9BB207F76D6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterRouteEntryResponseBody) SetRequestId(v string) *DeleteTransitRouterRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
