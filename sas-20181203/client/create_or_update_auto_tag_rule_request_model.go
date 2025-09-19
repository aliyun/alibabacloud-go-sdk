// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAutoTagRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckAll(v bool) *CreateOrUpdateAutoTagRuleRequest
	GetCheckAll() *bool
	SetExpression(v string) *CreateOrUpdateAutoTagRuleRequest
	GetExpression() *string
	SetRuleDesc(v string) *CreateOrUpdateAutoTagRuleRequest
	GetRuleDesc() *string
	SetRuleId(v int64) *CreateOrUpdateAutoTagRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *CreateOrUpdateAutoTagRuleRequest
	GetRuleName() *string
	SetTagContext(v string) *CreateOrUpdateAutoTagRuleRequest
	GetTagContext() *string
	SetTagType(v string) *CreateOrUpdateAutoTagRuleRequest
	GetTagType() *string
}

type CreateOrUpdateAutoTagRuleRequest struct {
	// Specifies whether to check the rule on the backend. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CheckAll *bool `json:"CheckAll,omitempty" xml:"CheckAll,omitempty"`
	// The expression of the rule.
	//
	// example:
	//
	// [{"groups":"0","fieldValueType":"string","field":"internetIp","operator":"equals","value":"12.0.0.1"}]
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// describe
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the rule.
	//
	// >  You can call the [ListAutoTagRules](~~ListAutoTagRules~~) operation to query the ID.
	//
	// example:
	//
	// 300566
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tag specified by the operation type of the rule.
	//
	// 	- If TagType is set to group, set this parameter to {"groupId":XXX}. XXX specifies the ID of the group. You can call the [DescribeGroupStruct](~~DescribeGroupStruct~~) operation to query the ID.
	//
	// 	- If TagType is set to tag, set this parameter to {"tagId":XXX}. XXX specifies the ID of the tag. You can call the [DescribeGroupedTags](~~DescribeGroupedTags~~) operation to query the ID.
	//
	// example:
	//
	// {"tagId":7804789}
	TagContext *string `json:"TagContext,omitempty" xml:"TagContext,omitempty"`
	// The operation type of the rule. Valid values:
	//
	// 	- **group**
	//
	// 	- **tag**
	//
	// This parameter is required.
	//
	// example:
	//
	// tag
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s CreateOrUpdateAutoTagRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAutoTagRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetCheckAll() *bool {
	return s.CheckAll
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetTagContext() *string {
	return s.TagContext
}

func (s *CreateOrUpdateAutoTagRuleRequest) GetTagType() *string {
	return s.TagType
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetCheckAll(v bool) *CreateOrUpdateAutoTagRuleRequest {
	s.CheckAll = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetExpression(v string) *CreateOrUpdateAutoTagRuleRequest {
	s.Expression = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetRuleDesc(v string) *CreateOrUpdateAutoTagRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetRuleId(v int64) *CreateOrUpdateAutoTagRuleRequest {
	s.RuleId = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetRuleName(v string) *CreateOrUpdateAutoTagRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetTagContext(v string) *CreateOrUpdateAutoTagRuleRequest {
	s.TagContext = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) SetTagType(v string) *CreateOrUpdateAutoTagRuleRequest {
	s.TagType = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleRequest) Validate() error {
	return dara.Validate(s)
}
