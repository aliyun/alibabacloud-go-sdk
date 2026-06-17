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
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
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
