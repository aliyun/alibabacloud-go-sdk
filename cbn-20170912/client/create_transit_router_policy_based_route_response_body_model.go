// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyBasedRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyBasedRouteId(v string) *CreateTransitRouterPolicyBasedRouteResponseBody
	GetPolicyBasedRouteId() *string
	SetRequestId(v string) *CreateTransitRouterPolicyBasedRouteResponseBody
	GetRequestId() *string
}

type CreateTransitRouterPolicyBasedRouteResponseBody struct {
	// The instance ID of the policy-based routing rule.
	//
	// example:
	//
	// pbr-123abczxc
	PolicyBasedRouteId *string `json:"PolicyBasedRouteId,omitempty" xml:"PolicyBasedRouteId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13526224-5780-4426-8BDF-BC8B08700F22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransitRouterPolicyBasedRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyBasedRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyBasedRouteResponseBody) GetPolicyBasedRouteId() *string {
	return s.PolicyBasedRouteId
}

func (s *CreateTransitRouterPolicyBasedRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterPolicyBasedRouteResponseBody) SetPolicyBasedRouteId(v string) *CreateTransitRouterPolicyBasedRouteResponseBody {
	s.PolicyBasedRouteId = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteResponseBody) SetRequestId(v string) *CreateTransitRouterPolicyBasedRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
