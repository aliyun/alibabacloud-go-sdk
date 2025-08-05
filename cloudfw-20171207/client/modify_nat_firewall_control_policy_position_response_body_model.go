// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyPositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatFirewallControlPolicyPositionResponseBody
	GetRequestId() *string
}

type ModifyNatFirewallControlPolicyPositionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 60617208-F5F7-5B44-BB1E-3AC1B6FCD627
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatFirewallControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyNatFirewallControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponseBody) Validate() error {
	return dara.Validate(s)
}
