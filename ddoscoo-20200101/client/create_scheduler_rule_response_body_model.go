// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *CreateSchedulerRuleResponseBody
	GetCname() *string
	SetRequestId(v string) *CreateSchedulerRuleResponseBody
	GetRequestId() *string
	SetRuleName(v string) *CreateSchedulerRuleResponseBody
	GetRuleName() *string
}

type CreateSchedulerRuleResponseBody struct {
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
	// 8DFB602D-1AAC-46C4-90F2-C84086E7A6E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateSchedulerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleResponseBody) GetCname() *string {
	return s.Cname
}

func (s *CreateSchedulerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchedulerRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateSchedulerRuleResponseBody) SetCname(v string) *CreateSchedulerRuleResponseBody {
	s.Cname = &v
	return s
}

func (s *CreateSchedulerRuleResponseBody) SetRequestId(v string) *CreateSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchedulerRuleResponseBody) SetRuleName(v string) *CreateSchedulerRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *CreateSchedulerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
