// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallConfigureResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k6f8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
