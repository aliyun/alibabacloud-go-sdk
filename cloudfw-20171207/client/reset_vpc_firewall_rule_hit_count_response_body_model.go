// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetVpcFirewallRuleHitCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetVpcFirewallRuleHitCountResponseBody
	GetRequestId() *string
}

type ResetVpcFirewallRuleHitCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A918B4F6-482F-5A91-8F65-AFFFF1FC04EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetVpcFirewallRuleHitCountResponseBody) SetRequestId(v string) *ResetVpcFirewallRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponseBody) Validate() error {
	return dara.Validate(s)
}
