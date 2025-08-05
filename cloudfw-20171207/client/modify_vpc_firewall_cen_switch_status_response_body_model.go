// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenSwitchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallCenSwitchStatusResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallCenSwitchStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125afj2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallCenSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
