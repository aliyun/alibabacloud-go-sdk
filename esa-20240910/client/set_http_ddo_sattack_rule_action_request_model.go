// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleAction(v string) *SetHttpDDoSAttackRuleActionRequest
	GetRuleAction() *string
	SetRuleIds(v string) *SetHttpDDoSAttackRuleActionRequest
	GetRuleIds() *string
	SetSiteId(v int64) *SetHttpDDoSAttackRuleActionRequest
	GetSiteId() *int64
}

type SetHttpDDoSAttackRuleActionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// js
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100000
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackRuleActionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleActionRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleActionRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *SetHttpDDoSAttackRuleActionRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *SetHttpDDoSAttackRuleActionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackRuleActionRequest) SetRuleAction(v string) *SetHttpDDoSAttackRuleActionRequest {
	s.RuleAction = &v
	return s
}

func (s *SetHttpDDoSAttackRuleActionRequest) SetRuleIds(v string) *SetHttpDDoSAttackRuleActionRequest {
	s.RuleIds = &v
	return s
}

func (s *SetHttpDDoSAttackRuleActionRequest) SetSiteId(v int64) *SetHttpDDoSAttackRuleActionRequest {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackRuleActionRequest) Validate() error {
	return dara.Validate(s)
}
