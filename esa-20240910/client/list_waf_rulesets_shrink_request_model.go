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
	// Page number, specifying the current page number for paginated queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, specifying the number of records per page for paginated queries.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// WAF operation phase, specifying the rule set phase to query.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// Query parameters, passed in JSON format, containing various filtering conditions.
	//
	// example:
	//
	// http_bot
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
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
