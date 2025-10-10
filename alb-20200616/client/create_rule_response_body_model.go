// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateRuleResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateRuleResponseBody
	GetRuleId() *string
}

type CreateRuleResponseBody struct {
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
	// 365F4154-92F6-4AE4-92F8-7FF34B540750
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rule ID.
	//
	// example:
	//
	// rule-a3x3pg1yohq3lq****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateRuleResponseBody) SetJobId(v string) *CreateRuleResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRuleId(v string) *CreateRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
