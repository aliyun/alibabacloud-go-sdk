// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetRuleByIdRequest
	GetBaseMeAgentId() *int64
	SetRuleId(v int64) *GetRuleByIdRequest
	GetRuleId() *int64
}

type GetRuleByIdRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 53
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRuleByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRuleByIdRequest) GoString() string {
	return s.String()
}

func (s *GetRuleByIdRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetRuleByIdRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetRuleByIdRequest) SetBaseMeAgentId(v int64) *GetRuleByIdRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetRuleByIdRequest) SetRuleId(v int64) *GetRuleByIdRequest {
	s.RuleId = &v
	return s
}

func (s *GetRuleByIdRequest) Validate() error {
	return dara.Validate(s)
}
