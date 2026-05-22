// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWafRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafRulesRequest
	GetPageSize() *int32
	SetPhase(v string) *ListWafRulesRequest
	GetPhase() *string
	SetQueryArgs(v *ListWafRulesRequestQueryArgs) *ListWafRulesRequest
	GetQueryArgs() *ListWafRulesRequestQueryArgs
	SetRulesetId(v int64) *ListWafRulesRequest
	GetRulesetId() *int64
	SetSiteId(v int64) *ListWafRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListWafRulesRequest
	GetSiteVersion() *int32
}

type ListWafRulesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http_custom
	Phase     *string                       `json:"Phase,omitempty" xml:"Phase,omitempty"`
	QueryArgs *ListWafRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	RulesetId *int64                        `json:"RulesetId,omitempty" xml:"RulesetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafRulesRequest) GetQueryArgs() *ListWafRulesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListWafRulesRequest) GetRulesetId() *int64 {
	return s.RulesetId
}

func (s *ListWafRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListWafRulesRequest) SetPageNumber(v int32) *ListWafRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafRulesRequest) SetPageSize(v int32) *ListWafRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafRulesRequest) SetPhase(v string) *ListWafRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafRulesRequest) SetQueryArgs(v *ListWafRulesRequestQueryArgs) *ListWafRulesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafRulesRequest) SetRulesetId(v int64) *ListWafRulesRequest {
	s.RulesetId = &v
	return s
}

func (s *ListWafRulesRequest) SetSiteId(v int64) *ListWafRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafRulesRequest) SetSiteVersion(v int32) *ListWafRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListWafRulesRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWafRulesRequestQueryArgs struct {
	ConfigValueLike *string `json:"ConfigValueLike,omitempty" xml:"ConfigValueLike,omitempty"`
	Desc            *bool   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// 20000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// example
	IdNameLike *string `json:"IdNameLike,omitempty" xml:"IdNameLike,omitempty"`
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	// example:
	//
	// position
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafRulesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafRulesRequestQueryArgs) GetConfigValueLike() *string {
	return s.ConfigValueLike
}

func (s *ListWafRulesRequestQueryArgs) GetDesc() *bool {
	return s.Desc
}

func (s *ListWafRulesRequestQueryArgs) GetId() *int64 {
	return s.Id
}

func (s *ListWafRulesRequestQueryArgs) GetIdNameLike() *string {
	return s.IdNameLike
}

func (s *ListWafRulesRequestQueryArgs) GetNameLike() *string {
	return s.NameLike
}

func (s *ListWafRulesRequestQueryArgs) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWafRulesRequestQueryArgs) GetStatus() *string {
	return s.Status
}

func (s *ListWafRulesRequestQueryArgs) SetConfigValueLike(v string) *ListWafRulesRequestQueryArgs {
	s.ConfigValueLike = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetDesc(v bool) *ListWafRulesRequestQueryArgs {
	s.Desc = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetId(v int64) *ListWafRulesRequestQueryArgs {
	s.Id = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetIdNameLike(v string) *ListWafRulesRequestQueryArgs {
	s.IdNameLike = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetNameLike(v string) *ListWafRulesRequestQueryArgs {
	s.NameLike = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetOrderBy(v string) *ListWafRulesRequestQueryArgs {
	s.OrderBy = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) SetStatus(v string) *ListWafRulesRequestQueryArgs {
	s.Status = &v
	return s
}

func (s *ListWafRulesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
