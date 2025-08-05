// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type DeleteNatFirewallControlPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31306869-C4BC-56F3-BBE6-C377CF9C8E1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatFirewallControlPolicyResponseBody) SetRequestId(v string) *DeleteNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
