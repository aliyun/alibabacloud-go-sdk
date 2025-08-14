// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterRouteTableResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EAB2F133-8556-4D7C-9E91-7EE4FE9CC7D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterRouteTableResponseBody) SetRequestId(v string) *DeleteTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
