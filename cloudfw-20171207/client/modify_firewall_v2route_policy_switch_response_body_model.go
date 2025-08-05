// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallV2RoutePolicySwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFirewallV2RoutePolicySwitchResponseBody
	GetRequestId() *string
}

type ModifyFirewallV2RoutePolicySwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1F6D4A8D-EC01-5996-A61A-AA3B56490C00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFirewallV2RoutePolicySwitchResponseBody) SetRequestId(v string) *ModifyFirewallV2RoutePolicySwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
