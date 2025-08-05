// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenConfigureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcFirewallCenConfigureResponseBody
	GetRequestId() *string
}

type ModifyVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k6f8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponseBody) Validate() error {
	return dara.Validate(s)
}
