// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterRouteTableResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9D6D5548-F271-41C4-AA9F-A62F5599085B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterRouteTableResponseBody) SetRequestId(v string) *UpdateTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
