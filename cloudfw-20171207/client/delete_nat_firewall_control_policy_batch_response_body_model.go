// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatFirewallControlPolicyBatchResponseBody
	GetRequestId() *string
}

type DeleteNatFirewallControlPolicyBatchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 450D47F5-956E-543E-8502-2F71C8C54E72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatFirewallControlPolicyBatchResponseBody) SetRequestId(v string) *DeleteNatFirewallControlPolicyBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
