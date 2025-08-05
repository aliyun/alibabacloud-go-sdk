// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsFirewallPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDnsFirewallPolicyResponseBody
	GetRequestId() *string
}

type DeleteDnsFirewallPolicyResponseBody struct {
	// example:
	//
	// 32314C1E-82CF-582C-853A-B1773F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsFirewallPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsFirewallPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsFirewallPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDnsFirewallPolicyResponseBody) SetRequestId(v string) *DeleteDnsFirewallPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDnsFirewallPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
