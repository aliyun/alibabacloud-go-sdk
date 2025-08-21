// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySchedulerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *ModifySchedulerRuleResponseBody
	GetCname() *string
	SetRequestId(v string) *ModifySchedulerRuleResponseBody
	GetRequestId() *string
	SetRuleName(v string) *ModifySchedulerRuleResponseBody
	GetRuleName() *string
}

type ModifySchedulerRuleResponseBody struct {
	// The CNAME that is assigned by Sec-Traffic Manager for the scheduling rule.
	//
	// > To enable the scheduling rule, you must map the domain name of the service to the CNAME.
	//
	// example:
	//
	// 48k7b372gpl4****.aliyunddos0001.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FFC77501-BDF8-4BC8-9BF5-B295FBC3189B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifySchedulerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySchedulerRuleResponseBody) GetCname() *string {
	return s.Cname
}

func (s *ModifySchedulerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySchedulerRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifySchedulerRuleResponseBody) SetCname(v string) *ModifySchedulerRuleResponseBody {
	s.Cname = &v
	return s
}

func (s *ModifySchedulerRuleResponseBody) SetRequestId(v string) *ModifySchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySchedulerRuleResponseBody) SetRuleName(v string) *ModifySchedulerRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *ModifySchedulerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
