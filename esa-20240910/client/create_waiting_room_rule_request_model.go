// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v string) *CreateWaitingRoomRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateWaitingRoomRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateWaitingRoomRuleRequest
	GetRuleName() *string
	SetSiteId(v int64) *CreateWaitingRoomRuleRequest
	GetSiteId() *int64
	SetWaitingRoomId(v string) *CreateWaitingRoomRuleRequest
	GetWaitingRoomId() *string
}

type CreateWaitingRoomRuleRequest struct {
	// This parameter is required.
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	WaitingRoomId *string `json:"WaitingRoomId,omitempty" xml:"WaitingRoomId,omitempty"`
}

func (s CreateWaitingRoomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateWaitingRoomRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateWaitingRoomRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateWaitingRoomRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWaitingRoomRuleRequest) GetWaitingRoomId() *string {
	return s.WaitingRoomId
}

func (s *CreateWaitingRoomRuleRequest) SetRule(v string) *CreateWaitingRoomRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleEnable(v string) *CreateWaitingRoomRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetRuleName(v string) *CreateWaitingRoomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetSiteId(v int64) *CreateWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) SetWaitingRoomId(v string) *CreateWaitingRoomRuleRequest {
	s.WaitingRoomId = &v
	return s
}

func (s *CreateWaitingRoomRuleRequest) Validate() error {
	return dara.Validate(s)
}
