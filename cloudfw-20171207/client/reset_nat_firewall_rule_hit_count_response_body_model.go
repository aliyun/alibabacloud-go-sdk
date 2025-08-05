// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNatFirewallRuleHitCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetNatFirewallRuleHitCountResponseBody
	GetRequestId() *string
}

type ResetNatFirewallRuleHitCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5CAA0FFD-4B94-5BB9-8B0A-ECFC86A0E666
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNatFirewallRuleHitCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetNatFirewallRuleHitCountResponseBody) SetRequestId(v string) *ResetNatFirewallRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponseBody) Validate() error {
	return dara.Validate(s)
}
