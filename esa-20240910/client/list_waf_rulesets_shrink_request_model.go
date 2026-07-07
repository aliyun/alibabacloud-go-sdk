// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWafRulesetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesetsShrinkRequest
	GetPageSize() *int32
	SetPhase(v string) *ListWafRulesetsShrinkRequest
	GetPhase() *string
	SetQueryArgsShrink(v string) *ListWafRulesetsShrinkRequest
	GetQueryArgsShrink() *string
	SetSiteId(v int64) *ListWafRulesetsShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListWafRulesetsShrinkRequest
	GetSiteVersion() *int32
}

type ListWafRulesetsShrinkRequest struct {
	// The page number. Specifies the current page number for paging queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Specifies the number of records per page for paging queries.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The WAF rule execution phase. Valid values:
	//
	// - http_whitelist: whitelist rules
	//
	// - http_custom: custom rules
	//
	// - http_managed: managed rules
	//
	// - http_anti_scan: scan protection rules
	//
	// - http_ratelimit: frequency control rules
	//
	// - ip_access_rule: IP access rules
	//
	// - http_bot: advanced mode bots
	//
	// - http_security_level_rule: security rules
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The query parameters, passed in JSON format, including various filter conditions.
	//
	// example:
	//
	// http_bot
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
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

func (s ListWafRulesetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesetsShrinkRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesetsShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListWafRulesetsShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafRulesetsShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListWafRulesetsShrinkRequest) SetPageNumber(v int32) *ListWafRulesetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetPageSize(v int32) *ListWafRulesetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetPhase(v string) *ListWafRulesetsShrinkRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetQueryArgsShrink(v string) *ListWafRulesetsShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetSiteId(v int64) *ListWafRulesetsShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) SetSiteVersion(v int32) *ListWafRulesetsShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListWafRulesetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
