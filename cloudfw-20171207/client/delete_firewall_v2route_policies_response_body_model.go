// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallV2RoutePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFirewallV2RoutePoliciesResponseBody
	GetRequestId() *string
}

type DeleteFirewallV2RoutePoliciesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 22B6B889-4E9F-56B3-AF3D-53749C477D1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallV2RoutePoliciesResponseBody) SetRequestId(v string) *DeleteFirewallV2RoutePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
