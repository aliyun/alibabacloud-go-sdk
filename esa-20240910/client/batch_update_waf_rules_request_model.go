// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*WafRuleConfig) *BatchUpdateWafRulesRequest
	GetConfigs() []*WafRuleConfig
	SetPhase(v string) *BatchUpdateWafRulesRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchUpdateWafRulesRequest
	GetRulesetId() *int64
	SetShared(v *WafBatchRuleShared) *BatchUpdateWafRulesRequest
	GetShared() *WafBatchRuleShared
	SetSiteId(v int64) *BatchUpdateWafRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchUpdateWafRulesRequest
	GetSiteVersion() *int32
}

type BatchUpdateWafRulesRequest struct {
	// A list of configurations for individual rules.
	Configs []*WafRuleConfig `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The WAF rule runtime phase.
	//
	// - `http_whitelist`: whitelist rule
	//
	// - `http_custom`: custom rule
	//
	// - `http_managed`: managed rule
	//
	// - `http_anti_scan`: scan protection rule
	//
	// - `http_ratelimit`: rate limiting rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_bot`: advanced bot rule
	//
	// - `http_security_level_rule`: security rule
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF ruleset. You can call the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation to obtain this ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The configuration properties that are shared by all rules in this batch update.
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The ID of the site. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version of the site configuration. For sites that have configuration version management enabled, this parameter specifies the version to which the configuration applies. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchUpdateWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesRequest) GetConfigs() []*WafRuleConfig {
	return s.Configs
}

func (s *BatchUpdateWafRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchUpdateWafRulesRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchUpdateWafRulesRequest) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *BatchUpdateWafRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchUpdateWafRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchUpdateWafRulesRequest) SetConfigs(v []*WafRuleConfig) *BatchUpdateWafRulesRequest {
	s.Configs = v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetPhase(v string) *BatchUpdateWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetRulesetId(v int64) *BatchUpdateWafRulesRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetShared(v *WafBatchRuleShared) *BatchUpdateWafRulesRequest {
	s.Shared = v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetSiteId(v int64) *BatchUpdateWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) SetSiteVersion(v int32) *BatchUpdateWafRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchUpdateWafRulesRequest) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Shared != nil {
		if err := s.Shared.Validate(); err != nil {
			return err
		}
	}
	return nil
}
