// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteFirewallRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFirewallRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFirewallRulesResponseBody
	GetSuccess() *bool
}

type DeleteFirewallRulesResponseBody struct {
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 049A1520-6BD7-5572-8923-79215D2B4A94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFirewallRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFirewallRulesResponseBody) SetMessage(v string) *DeleteFirewallRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFirewallRulesResponseBody) SetRequestId(v string) *DeleteFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFirewallRulesResponseBody) SetSuccess(v bool) *DeleteFirewallRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFirewallRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
