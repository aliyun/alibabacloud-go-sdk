// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardScopeSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromGroup(v bool) *TravelStandardScopeSaveRequest
	GetFromGroup() *bool
	SetRuleId(v int64) *TravelStandardScopeSaveRequest
	GetRuleId() *int64
	SetScope(v int32) *TravelStandardScopeSaveRequest
	GetScope() *int32
}

type TravelStandardScopeSaveRequest struct {
	// example:
	//
	// false
	FromGroup *bool `json:"from_group,omitempty" xml:"from_group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6516571
	RuleId *int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scope *int32 `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s TravelStandardScopeSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardScopeSaveRequest) GoString() string {
	return s.String()
}

func (s *TravelStandardScopeSaveRequest) GetFromGroup() *bool {
	return s.FromGroup
}

func (s *TravelStandardScopeSaveRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardScopeSaveRequest) GetScope() *int32 {
	return s.Scope
}

func (s *TravelStandardScopeSaveRequest) SetFromGroup(v bool) *TravelStandardScopeSaveRequest {
	s.FromGroup = &v
	return s
}

func (s *TravelStandardScopeSaveRequest) SetRuleId(v int64) *TravelStandardScopeSaveRequest {
	s.RuleId = &v
	return s
}

func (s *TravelStandardScopeSaveRequest) SetScope(v int32) *TravelStandardScopeSaveRequest {
	s.Scope = &v
	return s
}

func (s *TravelStandardScopeSaveRequest) Validate() error {
	return dara.Validate(s)
}
