// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWafRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesShrinkRequest
	GetPageSize() *int32
	SetPhase(v string) *ListWafRulesShrinkRequest
	GetPhase() *string
	SetQueryArgsShrink(v string) *ListWafRulesShrinkRequest
	GetQueryArgsShrink() *string
	SetRulesetId(v int64) *ListWafRulesShrinkRequest
	GetRulesetId() *int64
	SetSiteId(v int64) *ListWafRulesShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListWafRulesShrinkRequest
	GetSiteVersion() *int32
}

type ListWafRulesShrinkRequest struct {
	// The page number for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size for pagination.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The WAF rule execution phase. Valid values:
	//
	// - http_whitelist: whitelist rule
	//
	// - http_custom: custom rule
	//
	// - http_managed: managed rule
	//
	// - http_anti_scan: scan protection rule
	//
	// - http_ratelimit: frequency control rule
	//
	// - ip_access_rule: IP access rule
	//
	// - http_bot: advanced mode bots
	//
	// - http_security_level_rule: security rule
	//
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The query filter conditions.
	//
	// example:
	//
	// http_custom
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
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
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListWafRulesShrinkRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *ListWafRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafRulesShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListWafRulesShrinkRequest) SetPageNumber(v int32) *ListWafRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetPageSize(v int32) *ListWafRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetPhase(v string) *ListWafRulesShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetQueryArgsShrink(v string) *ListWafRulesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetRulesetId(v int64) *ListWafRulesShrinkRequest {
	s.RulesetId = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetSiteId(v int64) *ListWafRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesShrinkRequest) SetSiteVersion(v int32) *ListWafRulesShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListWafRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
