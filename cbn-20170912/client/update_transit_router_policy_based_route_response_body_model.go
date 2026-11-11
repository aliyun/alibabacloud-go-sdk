// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyBasedRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterPolicyBasedRouteResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterPolicyBasedRouteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 13526224-5780-4426-8ADF-BC8B08700F23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterPolicyBasedRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyBasedRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyBasedRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterPolicyBasedRouteResponseBody) SetRequestId(v string) *UpdateTransitRouterPolicyBasedRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
