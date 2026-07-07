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
	// The specific configuration of the WAF rule (`WafRuleConfig` data structure). The required fields vary depending on the Phase value:
	//
	// - `http_custom`: `Expression` (match expression) and `Action` (action upon match) are required. Setting `Name` is recommended for easier identification.
	//
	// - `http_whitelist`: `Expression` is required. Matched requests are allowed directly (no Action).
	//
	// - `http_ratelimit`: `Expression` and `RateLimit` (rate limiting parameters) are required.
	//
	// - `ip_access_rule`: `Expression` (containing IP match) and `Action` are required.
	//
	// > The complete field definitions are based on the `WafRuleConfig` data structure. If required fields are missing, the service returns `InvalidParameter(400)` / `Rule.Config.Malformed`.
	Config *WafRuleConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	// The WAF rule execution phase. This **single-creation operation*	- supports the following phases (it does not support `http_anti_scan` or `http_bot`. For these two phases, use the batch operation `BatchCreateWafRules`):
	//
	// - `http_whitelist`: whitelist rule
	//
	// - `http_custom`: custom rule
	//
	// - `http_managed`: managed rule
	//
	// - `http_ratelimit`: rate limiting rule
	//
	// - `ip_access_rule`: IP access rule
	//
	// - `http_security_level_rule`: security rule
	//
	// > Note: `http_anti_scan` and `http_bot` can only be created through the batch operation. Passing these two values to this operation returns an error.
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF ruleset. You can call the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation to obtain the ruleset ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is 0.
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
