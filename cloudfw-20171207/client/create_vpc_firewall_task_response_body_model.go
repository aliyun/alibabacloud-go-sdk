// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *CreateVpcFirewallTaskResponseBody
	GetModule() *string
	SetRequestId(v string) *CreateVpcFirewallTaskResponseBody
	GetRequestId() *string
}

type CreateVpcFirewallTaskResponseBody struct {
	// example:
	//
	// ips_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 53252B14-BF7C-5A2D-9750-56F827EB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcFirewallTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallTaskResponseBody) GetModule() *string {
	return s.Module
}

func (s *CreateVpcFirewallTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcFirewallTaskResponseBody) SetModule(v string) *CreateVpcFirewallTaskResponseBody {
	s.Module = &v
	return s
}

func (s *CreateVpcFirewallTaskResponseBody) SetRequestId(v string) *CreateVpcFirewallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
