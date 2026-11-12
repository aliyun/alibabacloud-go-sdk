// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFirewallRulesResponseBody
	GetRequestId() *string
}

type DeleteFirewallRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallRulesResponseBody) SetRequestId(v string) *DeleteFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
