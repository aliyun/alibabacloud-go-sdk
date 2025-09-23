// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyFirewallRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyFirewallRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyFirewallRulesResponseBody
	GetSuccess() *bool
}

type ModifyFirewallRulesResponseBody struct {
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A8DBA3A7-82FB-5CBE-A002-8959E47D1D61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFirewallRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyFirewallRulesResponseBody) SetMessage(v string) *ModifyFirewallRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyFirewallRulesResponseBody) SetRequestId(v string) *ModifyFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirewallRulesResponseBody) SetSuccess(v bool) *ModifyFirewallRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyFirewallRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
