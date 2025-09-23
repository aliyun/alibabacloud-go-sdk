// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddFirewallRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFirewallRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddFirewallRulesResponseBody
	GetSuccess() *bool
}

type AddFirewallRulesResponseBody struct {
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *AddFirewallRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFirewallRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddFirewallRulesResponseBody) SetMessage(v string) *AddFirewallRulesResponseBody {
	s.Message = &v
	return s
}

func (s *AddFirewallRulesResponseBody) SetRequestId(v string) *AddFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFirewallRulesResponseBody) SetSuccess(v bool) *AddFirewallRulesResponseBody {
	s.Success = &v
	return s
}

func (s *AddFirewallRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
