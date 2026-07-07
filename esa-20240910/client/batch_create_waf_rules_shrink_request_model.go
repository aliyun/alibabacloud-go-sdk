// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateWafRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *BatchCreateWafRulesShrinkRequest
	GetConfigsShrink() *string
	SetPhase(v string) *BatchCreateWafRulesShrinkRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchCreateWafRulesShrinkRequest
	GetRulesetId() *int64
	SetSharedShrink(v string) *BatchCreateWafRulesShrinkRequest
	GetSharedShrink() *string
	SetSiteId(v int64) *BatchCreateWafRulesShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchCreateWafRulesShrinkRequest
	GetSiteVersion() *int32
}

type BatchCreateWafRulesShrinkRequest struct {
	// The list of rule configurations. Specifies the detailed configuration for each rule.
	//
	// **Required subfields for each phase*	- (applicable only to the two phases supported by this batch operation):
	//
	// - `http_anti_scan`: You must specify `Type` and at least one of `ManagedList` or `RateLimit`.
	//
	// - `http_bot`: You must specify the advanced mode bots configuration. The subfields are defined in the `WafRuleConfig` data structure.
	//
	// > Note: Other phases such as `http_custom` and `http_whitelist` cannot use this batch operation. Use the single-rule operation `CreateWafRule` instead. The subfield constraints are described in the single-rule operation documentation.
	//
	// > If `Configs` is not specified or required subfields are missing, the service returns `InvalidParameter(400)` or `Rule.Config.Malformed`.
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The WAF rule execution phase. This **batch operation supports only*	- the following two phases. For other phases, use the single-rule operations `CreateWafRule` or `UpdateWafRule`:
	//
	// - `http_anti_scan`: scan protection rules
	//
	// - `http_bot`: advanced mode bots
	//
	// > Note: The `http_anti_scan` and `http_bot` phases **support only batch creation**. The single-rule operation `CreateWafRule` does not accept these two values. Conversely, other phases such as `http_custom` and `http_whitelist` can be created only by using single-rule operations and cannot use this batch operation.
	//
	// **Required constraint**: Although this parameter is marked as optional (required: false) in the specification, it is **required*	- when you call this batch operation. If this parameter is not specified, the service returns `InvalidParameter(400)`.
	//
	// **Plan prerequisite**: `http_anti_scan` requires the site to have a **high or higher plan**. Calling this operation with a basic plan returns `Phase.HttpAntiScan.NotSupport`. Verify the site plan before calling this operation.
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
	// The shared configuration for multiple rules. Specifies the common properties of multiple rules.
	//
	// **Conditionally required**: Although this parameter is marked as optional (required: false) in the specification, it is **required*	- when `Phase=http_anti_scan`. If this parameter is not specified, the service returns `InvalidParameter(400)`.
	//
	// **Subfield requirements**: In the `http_anti_scan` phase, Shared must include shared fields such as `Name` (rule name) and `Action` (rule action). For other phases, the required subfields of Shared vary depending on the specific phase.
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
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

func (s BatchCreateWafRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateWafRulesShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *BatchCreateWafRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchCreateWafRulesShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchCreateWafRulesShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *BatchCreateWafRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchCreateWafRulesShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchCreateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetPhase(v string) *BatchCreateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetRulesetId(v int64) *BatchCreateWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchCreateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteId(v int64) *BatchCreateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchCreateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchCreateWafRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
