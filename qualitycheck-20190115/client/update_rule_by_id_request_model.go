// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateRuleByIdRequest
	GetBaseMeAgentId() *int64
	SetIsCopy(v bool) *UpdateRuleByIdRequest
	GetIsCopy() *bool
	SetJsonStrForRule(v string) *UpdateRuleByIdRequest
	GetJsonStrForRule() *string
	SetReturnRelatedSchemes(v bool) *UpdateRuleByIdRequest
	GetReturnRelatedSchemes() *bool
	SetRuleId(v int64) *UpdateRuleByIdRequest
	GetRuleId() *int64
}

type UpdateRuleByIdRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	JsonStrForRule *string `json:"JsonStrForRule,omitempty" xml:"JsonStrForRule,omitempty"`
	// example:
	//
	// 1
	ReturnRelatedSchemes *bool `json:"ReturnRelatedSchemes,omitempty" xml:"ReturnRelatedSchemes,omitempty"`
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateRuleByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleByIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateRuleByIdRequest) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *UpdateRuleByIdRequest) GetJsonStrForRule() *string {
	return s.JsonStrForRule
}

func (s *UpdateRuleByIdRequest) GetReturnRelatedSchemes() *bool {
	return s.ReturnRelatedSchemes
}

func (s *UpdateRuleByIdRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *UpdateRuleByIdRequest) SetBaseMeAgentId(v int64) *UpdateRuleByIdRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetIsCopy(v bool) *UpdateRuleByIdRequest {
	s.IsCopy = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetJsonStrForRule(v string) *UpdateRuleByIdRequest {
	s.JsonStrForRule = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetReturnRelatedSchemes(v bool) *UpdateRuleByIdRequest {
	s.ReturnRelatedSchemes = &v
	return s
}

func (s *UpdateRuleByIdRequest) SetRuleId(v int64) *UpdateRuleByIdRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleByIdRequest) Validate() error {
	return dara.Validate(s)
}
