// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallIPSWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallIPSWhitelistResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallIPSWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F2E8D49A-E5AA-5FF8-8822-25282DCF4BE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallIPSWhitelistResponseBody) SetRequestId(v string) *ModifyVpcFirewallIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
