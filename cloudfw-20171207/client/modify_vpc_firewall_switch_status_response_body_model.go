// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallSwitchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallSwitchStatusResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallSwitchStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125afj2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
