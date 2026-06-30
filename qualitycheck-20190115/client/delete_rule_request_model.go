// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *DeleteRuleRequest
	GetBaseMeAgentId() *int64
	SetForceDelete(v bool) *DeleteRuleRequest
	GetForceDelete() *bool
	SetIsSchemeData(v int32) *DeleteRuleRequest
	GetIsSchemeData() *int32
	SetRuleId(v int64) *DeleteRuleRequest
	GetRuleId() *int64
}

type DeleteRuleRequest struct {
	// Workspace ID.
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// Whether to delete the rule even if it is associated with a quality check task.
	//
	// example:
	//
	// true
	ForceDelete *bool `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// Indicates whether this is the new version of quality check. 0: old version. 1: new version. Default is 1.
	//
	// example:
	//
	// 1
	IsSchemeData *int32 `json:"IsSchemeData,omitempty" xml:"IsSchemeData,omitempty"`
	// Rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *DeleteRuleRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteRuleRequest) GetIsSchemeData() *int32 {
	return s.IsSchemeData
}

func (s *DeleteRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteRuleRequest) SetBaseMeAgentId(v int64) *DeleteRuleRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *DeleteRuleRequest) SetForceDelete(v bool) *DeleteRuleRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteRuleRequest) SetIsSchemeData(v int32) *DeleteRuleRequest {
	s.IsSchemeData = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v int64) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
