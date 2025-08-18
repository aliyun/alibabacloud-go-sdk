// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *CreateWafRuleShrinkRequest
	GetConfigShrink() *string
	SetPhase(v string) *CreateWafRuleShrinkRequest
	GetPhase() *string
	SetRulesetId(v int64) *CreateWafRuleShrinkRequest
	GetRulesetId() *int64
	SetSiteId(v int64) *CreateWafRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateWafRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateWafRuleShrinkRequest struct {
	// Rule configuration, specifying the detailed configuration for creating a rule.
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// WAF operation phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Ruleset ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateWafRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWafRuleShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *CreateWafRuleShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateWafRuleShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *CreateWafRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWafRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateWafRuleShrinkRequest) SetConfigShrink(v string) *CreateWafRuleShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetPhase(v string) *CreateWafRuleShrinkRequest {
	s.Phase = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetRulesetId(v int64) *CreateWafRuleShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetSiteId(v int64) *CreateWafRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) SetSiteVersion(v int32) *CreateWafRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateWafRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
