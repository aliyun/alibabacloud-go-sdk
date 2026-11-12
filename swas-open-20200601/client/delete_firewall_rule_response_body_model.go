// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFirewallRuleResponseBody
	GetRequestId() *string
}

type DeleteFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallRuleResponseBody) SetRequestId(v string) *DeleteFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
