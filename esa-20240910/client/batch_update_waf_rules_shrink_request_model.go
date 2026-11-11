// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateWafRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigsShrink(v string) *BatchUpdateWafRulesShrinkRequest
	GetConfigsShrink() *string
	SetPhase(v string) *BatchUpdateWafRulesShrinkRequest
	GetPhase() *string
	SetRulesetId(v int64) *BatchUpdateWafRulesShrinkRequest
	GetRulesetId() *int64
	SetSharedShrink(v string) *BatchUpdateWafRulesShrinkRequest
	GetSharedShrink() *string
	SetSiteId(v int64) *BatchUpdateWafRulesShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *BatchUpdateWafRulesShrinkRequest
	GetSiteVersion() *int32
}

type BatchUpdateWafRulesShrinkRequest struct {
	// The list of rule configurations, which specifies the detailed configuration of each rule.
	//
	// **Required sub-fields for each phase*	- (applicable only to the two phases supported by this batch operation):
	//
	// - `http_anti_scan`: You must provide `Type`, and at least one of `ManagedList` or `RateLimit`.
	//
	// - `http_bot`: You must provide configurations related to advanced bot management. The sub-fields are subject to the definition of the `WafRuleConfig` data structure.
	//
	// > Note: Other phases such as `http_custom` and `http_whitelist` cannot be updated by using this batch operation. Use the single-rule operation `UpdateWafRule` instead. The sub-field constraints are subject to the documentation of the single-rule operation.
	//
	// > Note: If `Configs` is empty or sub-fields are missing, the server returns an `InvalidParameter(400)` or `Rule.Config.Malformed` error.
	//
	// example:
	//
	// 10000001
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The execution phase of WAF rules. This **batch operation supports only*	- the following two phases. For other phases, use the single-rule operation `UpdateWafRule`:
	//
	// - `http_anti_scan`: scan protection rules
	//
	// - `http_bot`: advanced bot management rules
	//
	// > Note: The `http_anti_scan` and `http_bot` phases **support only batch updates**. The single-rule operation `UpdateWafRule` does not accept these two values. Conversely, other phases (such as `http_custom` and `http_whitelist`) can only be updated by using the single-rule operation and cannot be updated by using this batch operation.
	//
	// **Required constraint**: Although this parameter is marked as optional (required: false) in the specification, it **must be specified*	- when you call this batch operation. The server cannot determine the target rule set without the Phase parameter. If this parameter is not provided, an `InvalidParameter(400)` error is returned.
	//
	// example:
	//
	// http_anti_scan
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the WAF rule set. You can obtain the ID by calling the [ListWafRulesets](https://help.aliyun.com/document_detail/2878359.html) operation.
	//
	// example:
	//
	// 10000001
	RulesetId *int64 `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// The shared configuration for multiple rules, which specifies the common properties of multiple rules.
	//
	// **Conditional requirement**: Although this parameter is marked as optional (required: false) in the specification, it is **actually required*	- when `Phase` is set to `http_anti_scan`. If this parameter is not provided, the server returns an `InvalidParameter(400)` error.
	//
	// **Sub-field requirements**: In the `http_anti_scan` phase, the Shared parameter must include shared fields such as `Name` (rule name), `Expression` (match expression), and `Action` (rule action). For other phases, the required sub-fields of Shared depend on the specific phase.
	SharedShrink *string `json:"Shared,omitempty" xml:"Shared,omitempty"`
	// The ID of the site. You can obtain the ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s BatchUpdateWafRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateWafRulesShrinkRequest) GetConfigsShrink() *string {
	return s.ConfigsShrink
}

func (s *BatchUpdateWafRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *BatchUpdateWafRulesShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSharedShrink() *string {
	return s.SharedShrink
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BatchUpdateWafRulesShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *BatchUpdateWafRulesShrinkRequest) SetConfigsShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetPhase(v string) *BatchUpdateWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetRulesetId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSharedShrink(v string) *BatchUpdateWafRulesShrinkRequest {
	s.SharedShrink = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteId(v int64) *BatchUpdateWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) SetSiteVersion(v int32) *BatchUpdateWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *BatchUpdateWafRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
