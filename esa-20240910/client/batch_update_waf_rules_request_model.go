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
	// The list of rule configurations. Specifies the detailed configuration for each rule.
	//
	// **Required subfields for each phase*	- (applicable only to the two phases supported by this batch operation):
	//
	// - `http_anti_scan`: You must provide `Type` and at least one of `ManagedList` or `RateLimit`.
	//
	// - `http_bot`: You must provide the advanced mode bots configuration. The subfields are defined in the `WafRuleConfig` data structure.
	//
	// > Note: Other phases such as `http_custom` and `http_whitelist` cannot use this batch operation. Use the single-rule operation `UpdateWafRule` instead. The subfield constraints for those phases are described in the single-rule operation documentation.
	//
	// > Important: If `Configs` is missing or subfields are incomplete, the server returns `InvalidParameter(400)` or `Rule.Config.Malformed`.
	//
	// example:
	//
	// 10000001
	Configs []*WafRuleConfig `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The WAF rule execution phase. This **batch operation supports only*	- the following two phases. For other phases, use the single-rule operation `UpdateWafRule`:
	//
	// - `http_anti_scan`: scan protection rules
	//
	// - `http_bot`: advanced mode bots
	//
	// > Note: The `http_anti_scan` and `http_bot` phases **support only batch updates**. The single-rule operation `UpdateWafRule` does not accept these two values. Conversely, other phases such as `http_custom` and `http_whitelist` can be updated only by using the single-rule operation, not this batch operation.
	//
	// **Required constraint**: Although this parameter is marked as optional (required: false) in the specification, it is **required*	- when you call this batch operation. The server cannot determine the target ruleset without the Phase parameter and returns `InvalidParameter(400)` if it is not provided.
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF ruleset. You can call the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation to obtain the ruleset ID.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The shared configuration for multiple rules. Specifies the common properties shared across multiple rules.
	//
	// **Conditionally required**: Although this parameter is marked as optional (required: false) in the specification, it is **required*	- when `Phase=http_anti_scan`. The server returns `InvalidParameter(400)` if it is not provided.
	//
	// **Subfield requirements**: When the phase is `http_anti_scan`, Shared must include the `Name` (rule name), `Expression` (match expression), and `Action` (rule action) shared fields. For other phases, the required subfields of Shared vary depending on the specific phase.
	Shared *WafBatchRuleShared `json:"Shared,omitempty" xml:"Shared,omitempty"`
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
