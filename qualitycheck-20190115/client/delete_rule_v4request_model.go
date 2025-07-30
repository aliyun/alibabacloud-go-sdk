// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteRuleV4Request
	GetBaseMeAgentId() *int64
	SetForceDelete(v bool) *DeleteRuleV4Request
	GetForceDelete() *bool
	SetRuleId(v int64) *DeleteRuleV4Request
	GetRuleId() *int64
}

type DeleteRuleV4Request struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// false
	ForceDelete *bool `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleV4Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleV4Request) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4Request) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteRuleV4Request) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteRuleV4Request) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteRuleV4Request) SetBaseMeAgentId(v int64) *DeleteRuleV4Request {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteRuleV4Request) SetForceDelete(v bool) *DeleteRuleV4Request {
	s.ForceDelete = &v
	return s
}

func (s *DeleteRuleV4Request) SetRuleId(v int64) *DeleteRuleV4Request {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleV4Request) Validate() error {
	return dara.Validate(s)
}
