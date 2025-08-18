// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *WafRuleConfig) *CreateWafRuleRequest
	GetConfig() *WafRuleConfig
	SetPhase(v string) *CreateWafRuleRequest
	GetPhase() *string
	SetRulesetId(v int64) *CreateWafRuleRequest
	GetRulesetId() *int64
	SetSiteId(v int64) *CreateWafRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateWafRuleRequest
	GetSiteVersion() *int32
}

type CreateWafRuleRequest struct {
	// Rule configuration, specifying the detailed configuration for creating a rule.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
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

func (s CreateWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWafRuleRequest) GetConfig() *WafRuleConfig {
	return s.Config
}

func (s *CreateWafRuleRequest) GetPhase() *string {
	return s.Phase
}

func (s *CreateWafRuleRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *CreateWafRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateWafRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateWafRuleRequest) SetConfig(v *WafRuleConfig) *CreateWafRuleRequest {
	s.Config = v
	return s
}

func (s *CreateWafRuleRequest) SetPhase(v string) *CreateWafRuleRequest {
	s.Phase = &v
	return s
}

func (s *CreateWafRuleRequest) SetRulesetId(v int64) *CreateWafRuleRequest {
	s.RulesetId = &v
	return s
}

func (s *CreateWafRuleRequest) SetSiteId(v int64) *CreateWafRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateWafRuleRequest) SetSiteVersion(v int32) *CreateWafRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateWafRuleRequest) Validate() error {
	return dara.Validate(s)
}
