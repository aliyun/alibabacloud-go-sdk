// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*WafRuleConfig) *BatchCreateWafRulesRequest
	GetConfigs() []*WafRuleConfig
	SetPhase(v string) *BatchCreateWafRulesRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchCreateWafRulesRequest
	GetRulesetId() *int64
	SetShared(v *WafBatchRuleShared) *BatchCreateWafRulesRequest
	GetShared() *WafBatchRuleShared
	SetSiteId(v int64) *BatchCreateWafRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchCreateWafRulesRequest
	GetSiteVersion() *int32
}

type BatchCreateWafRulesRequest struct {
	// A list of configurations for each rule, specifying detailed configurations for each rule.
	Configs []*WafRuleConfig `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// WAF rule type, with values:
	//
	// - **http_anti_scan**: Scan protection.
	//
	// - **http_bot**: Bots.
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Ruleset ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// Shared configuration for multiple rules, specifying common attributes of multiple rules.
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
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

func (s BatchCreateWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesRequest) GetConfigs() []*WafRuleConfig {
	return s.Configs
}

func (s *BatchCreateWafRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchCreateWafRulesRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchCreateWafRulesRequest) GetShared() *WafBatchRuleShared {
	return s.Shared
}

func (s *BatchCreateWafRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchCreateWafRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchCreateWafRulesRequest) SetConfigs(v []*WafRuleConfig) *BatchCreateWafRulesRequest {
	s.Configs = v
	return s
}

func (s *BatchCreateWafRulesRequest) SetPhase(v string) *BatchCreateWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *BatchCreateWafRulesRequest) SetRulesetId(v int64) *BatchCreateWafRulesRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchCreateWafRulesRequest) SetShared(v *WafBatchRuleShared) *BatchCreateWafRulesRequest {
	s.Shared = v
	return s
}

func (s *BatchCreateWafRulesRequest) SetSiteId(v int64) *BatchCreateWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateWafRulesRequest) SetSiteVersion(v int32) *BatchCreateWafRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchCreateWafRulesRequest) Validate() error {
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
