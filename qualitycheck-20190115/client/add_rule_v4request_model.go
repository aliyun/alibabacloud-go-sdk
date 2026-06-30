// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AddRuleV4Request
	GetBaseMeAgentId() *int64
	SetIsCopy(v bool) *AddRuleV4Request
	GetIsCopy() *bool
	SetJsonStrForRule(v string) *AddRuleV4Request
	GetJsonStrForRule() *string
}

type AddRuleV4Request struct {
	// The workspace ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Specifies whether to copy the rule. Set this to \\`true\\` to copy the rule.
	//
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// For more information about JsonStrForRule, see [`Rule object description`](https://help.aliyun.com/document_detail/453053.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 无
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
}

func (s AddRuleV4Request) String() string {
	return dara.Prettify(s)
}

func (s AddRuleV4Request) GoString() string {
	return s.String()
}

func (s *AddRuleV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AddRuleV4Request) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *AddRuleV4Request) GetJsonStrForRule() *string {
	return s.JsonStrForRule
}

func (s *AddRuleV4Request) SetBaseMeAgentId(v int64) *AddRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *AddRuleV4Request) SetIsCopy(v bool) *AddRuleV4Request {
	s.IsCopy = &v
	return s
}

func (s *AddRuleV4Request) SetJsonStrForRule(v string) *AddRuleV4Request {
	s.JsonStrForRule = &v
	return s
}

func (s *AddRuleV4Request) Validate() error {
	return dara.Validate(s)
}
