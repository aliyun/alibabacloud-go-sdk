// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWafRulesetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesetsRequest
	GetPageSize() *int32
	SetPhase(v string) *ListWafRulesetsRequest
	GetPhase() *string
	SetQueryArgs(v *ListWafRulesetsRequestQueryArgs) *ListWafRulesetsRequest
	GetQueryArgs() *ListWafRulesetsRequestQueryArgs
	SetSiteId(v int64) *ListWafRulesetsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListWafRulesetsRequest
	GetSiteVersion() *int32
}

type ListWafRulesetsRequest struct {
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
	QueryArgs *ListWafRulesetsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
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

func (s ListWafRulesetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesetsRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesetsRequest) GetQueryArgs() *ListWafRulesetsRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListWafRulesetsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafRulesetsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListWafRulesetsRequest) SetPageNumber(v int32) *ListWafRulesetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesetsRequest) SetPageSize(v int32) *ListWafRulesetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesetsRequest) SetPhase(v string) *ListWafRulesetsRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesetsRequest) SetQueryArgs(v *ListWafRulesetsRequestQueryArgs) *ListWafRulesetsRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafRulesetsRequest) SetSiteId(v int64) *ListWafRulesetsRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesetsRequest) SetSiteVersion(v int32) *ListWafRulesetsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListWafRulesetsRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWafRulesetsRequestQueryArgs struct {
	// The fuzzy match string for the ruleset ID, ruleset name, rule ID, or rule name.
	//
	// example:
	//
	// example
	AnyLike *string `json:"AnyLike,omitempty" xml:"AnyLike,omitempty"`
	// Specifies whether to sort the results in descending order.
	Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The fuzzy match string for the ruleset name.
	//
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// The column by which to sort the results.
	//
	// example:
	//
	// id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
}

func (s ListWafRulesetsRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsRequestQueryArgs) GetAnyLike() *string {
	return s.AnyLike
}

func (s *ListWafRulesetsRequestQueryArgs) GetDesc() *bool {
	return s.Desc
}

func (s *ListWafRulesetsRequestQueryArgs) GetNameLike() *string {
	return s.NameLike
}

func (s *ListWafRulesetsRequestQueryArgs) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWafRulesetsRequestQueryArgs) SetAnyLike(v string) *ListWafRulesetsRequestQueryArgs {
	s.AnyLike = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetDesc(v bool) *ListWafRulesetsRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetNameLike(v string) *ListWafRulesetsRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) SetOrderBy(v string) *ListWafRulesetsRequestQueryArgs {
	s.OrderBy = &v
	return s
}

func (s *ListWafRulesetsRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
