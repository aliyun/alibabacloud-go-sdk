// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallAclEngineModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallAclEngineModeResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallAclEngineModeResponseBody struct {
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallAclEngineModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallAclEngineModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallAclEngineModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallAclEngineModeResponseBody) SetRequestId(v string) *ModifyVpcFirewallAclEngineModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeResponseBody) Validate() error {
	return dara.Validate(s)
}
