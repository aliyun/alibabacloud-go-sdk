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
	// The detailed configuration of the WAF rule.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// The phase in which the WAF rule runs.
	//
	// - `http_whitelist`: whitelist rule
	//
	// - `http_custom`: custom rule
	//
	// - `http_managed`: managed rule
	//
	// - `http_anti_scan`: anti-scan rule
	//
	// - `http_ratelimit`: rate limit rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_bot`: Advanced Mode Bots
	//
	// - `http_security_level_rule`: Security Rule
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF ruleset. You can obtain this ID by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The ID of the site. You can obtain this ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// If version management is enabled for the site, use this parameter to specify the version to which the configuration applies. The default is 0.
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
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
