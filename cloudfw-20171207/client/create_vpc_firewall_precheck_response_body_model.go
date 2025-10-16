// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallPrecheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrecheckId(v string) *CreateVpcFirewallPrecheckResponseBody
	GetPrecheckId() *string
	SetRequestId(v string) *CreateVpcFirewallPrecheckResponseBody
	GetRequestId() *string
}

type CreateVpcFirewallPrecheckResponseBody struct {
	// example:
	//
	// 4197
	PrecheckId *string `json:"PrecheckId,omitempty" xml:"PrecheckId,omitempty"`
	// example:
	//
	// C5BE1AA4-934A-5085-89CC-9AD1CAC3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcFirewallPrecheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallPrecheckResponseBody) GetPrecheckId() *string {
	return s.PrecheckId
}

func (s *CreateVpcFirewallPrecheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallPrecheckResponseBody) SetPrecheckId(v string) *CreateVpcFirewallPrecheckResponseBody {
	s.PrecheckId = &v
	return s
}

func (s *CreateVpcFirewallPrecheckResponseBody) SetRequestId(v string) *CreateVpcFirewallPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallPrecheckResponseBody) Validate() error {
	return dara.Validate(s)
}
