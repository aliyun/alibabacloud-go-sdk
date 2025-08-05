// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2RoutePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrFirewallV2RoutePolicyResponseBody
	GetRequestId() *string
	SetTrFirewallRoutePolicyId(v string) *CreateTrFirewallV2RoutePolicyResponseBody
	GetTrFirewallRoutePolicyId() *string
}

type CreateTrFirewallV2RoutePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C91D68BA-A0BE-51BF-A0F1-1CB5C57FE58D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-8ebed27e13e14ce2****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) SetRequestId(v string) *CreateTrFirewallV2RoutePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) SetTrFirewallRoutePolicyId(v string) *CreateTrFirewallV2RoutePolicyResponseBody {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
