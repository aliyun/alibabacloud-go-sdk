// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v string) *UpdateWaitingRoomRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateWaitingRoomRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateWaitingRoomRuleRequest
	GetRuleName() *string
	SetSiteId(v int64) *UpdateWaitingRoomRuleRequest
	GetSiteId() *int64
	SetWaitingRoomRuleId(v int64) *UpdateWaitingRoomRuleRequest
	GetWaitingRoomRuleId() *int64
}

type UpdateWaitingRoomRuleRequest struct {
	// This parameter is required.
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s UpdateWaitingRoomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateWaitingRoomRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateWaitingRoomRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateWaitingRoomRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateWaitingRoomRuleRequest) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *UpdateWaitingRoomRuleRequest) SetRule(v string) *UpdateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleEnable(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetRuleName(v string) *UpdateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetSiteId(v int64) *UpdateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) SetWaitingRoomRuleId(v int64) *UpdateWaitingRoomRuleRequest {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *UpdateWaitingRoomRuleRequest) Validate() error {
	return dara.Validate(s)
}
