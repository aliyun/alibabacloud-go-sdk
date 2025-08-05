// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCopyVpcFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCopyVpcFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type BatchCopyVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k6j3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCopyVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *BatchCopyVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
