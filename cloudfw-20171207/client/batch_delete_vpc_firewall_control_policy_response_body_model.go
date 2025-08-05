// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVpcFirewallControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteVpcFirewallControlPolicyResponseBody
	GetRequestId() *string
}

type BatchDeleteVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF7334F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *BatchDeleteVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
