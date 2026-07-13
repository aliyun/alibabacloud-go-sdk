// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceArn(v string) *CheckRulesRequest
	GetResourceArn() *string
	SetRuleId(v string) *CheckRulesRequest
	GetRuleId() *string
}

type CheckRulesRequest struct {
	// The unique identifier of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the rule to update. If you do not specify this parameter, all rules are updated.
	//
	// example:
	//
	// rule-000***dav
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CheckRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRulesRequest) GoString() string {
	return s.String()
}

func (s *CheckRulesRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *CheckRulesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *CheckRulesRequest) SetResourceArn(v string) *CheckRulesRequest {
	s.ResourceArn = &v
	return s
}

func (s *CheckRulesRequest) SetRuleId(v string) *CheckRulesRequest {
	s.RuleId = &v
	return s
}

func (s *CheckRulesRequest) Validate() error {
	return dara.Validate(s)
}
