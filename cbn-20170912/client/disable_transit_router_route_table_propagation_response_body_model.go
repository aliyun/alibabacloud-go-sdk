// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTransitRouterRouteTablePropagationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableTransitRouterRouteTablePropagationResponseBody
	GetRequestId() *string
}

type DisableTransitRouterRouteTablePropagationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A7C43F99-B1E5-4A53-AB64-4BAE8AF4484E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableTransitRouterRouteTablePropagationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableTransitRouterRouteTablePropagationResponseBody) SetRequestId(v string) *DisableTransitRouterRouteTablePropagationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationResponseBody) Validate() error {
	return dara.Validate(s)
}
