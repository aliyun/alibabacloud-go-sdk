// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleIds(v string) *SetHttpDDoSAttackRuleStatusRequest
	GetRuleIds() *string
	SetSiteId(v int64) *SetHttpDDoSAttackRuleStatusRequest
	GetSiteId() *int64
	SetStatus(v string) *SetHttpDDoSAttackRuleStatusRequest
	GetStatus() *string
}

type SetHttpDDoSAttackRuleStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 87570
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetHttpDDoSAttackRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleStatusRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *SetHttpDDoSAttackRuleStatusRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackRuleStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetHttpDDoSAttackRuleStatusRequest) SetRuleIds(v string) *SetHttpDDoSAttackRuleStatusRequest {
	s.RuleIds = &v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusRequest) SetSiteId(v int64) *SetHttpDDoSAttackRuleStatusRequest {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusRequest) SetStatus(v string) *SetHttpDDoSAttackRuleStatusRequest {
	s.Status = &v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
