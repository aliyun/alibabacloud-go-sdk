// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallControlPolicyPositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallControlPolicyPositionResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallControlPolicyPositionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyVpcFirewallControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponseBody) Validate() error {
	return dara.Validate(s)
}
