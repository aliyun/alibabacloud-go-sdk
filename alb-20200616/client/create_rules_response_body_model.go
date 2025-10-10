// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateRulesResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateRulesResponseBody
	GetRequestId() *string
	SetRuleIds(v []*CreateRulesResponseBodyRuleIds) *CreateRulesResponseBody
	GetRuleIds() []*CreateRulesResponseBodyRuleIds
}

type CreateRulesResponseBody struct {
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A lower value specifies a higher priority.
	//
	// > The priorities of the forwarding rules created for the same listener is unique.
	RuleIds []*CreateRulesResponseBodyRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s CreateRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRulesResponseBody) GetRuleIds() []*CreateRulesResponseBodyRuleIds {
	return s.RuleIds
}

func (s *CreateRulesResponseBody) SetJobId(v string) *CreateRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRequestId(v string) *CreateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRuleIds(v []*CreateRulesResponseBodyRuleIds) *CreateRulesResponseBody {
	s.RuleIds = v
	return s
}

func (s *CreateRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRulesResponseBodyRuleIds struct {
	// The priority of the forwarding rule. Valid values: **1 to 10000**. A smaller value indicates a higher priority.
	//
	// > The priorities of the forwarding rules created for the same listener must be unique.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The forwarding rule ID.
	//
	// example:
	//
	// rule-a3x3pg1yohq3lq****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRulesResponseBodyRuleIds) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponseBodyRuleIds) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRuleIds) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRulesResponseBodyRuleIds) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateRulesResponseBodyRuleIds) SetPriority(v int32) *CreateRulesResponseBodyRuleIds {
	s.Priority = &v
	return s
}

func (s *CreateRulesResponseBodyRuleIds) SetRuleId(v string) *CreateRulesResponseBodyRuleIds {
	s.RuleId = &v
	return s
}

func (s *CreateRulesResponseBodyRuleIds) Validate() error {
	return dara.Validate(s)
}
