// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyBasedRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterPolicyBasedRouteResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterPolicyBasedRouteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E4B345CD-2CBA-4881-AF6D-E5D9BAE1CA7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterPolicyBasedRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyBasedRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyBasedRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterPolicyBasedRouteResponseBody) SetRequestId(v string) *DeleteTransitRouterPolicyBasedRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
