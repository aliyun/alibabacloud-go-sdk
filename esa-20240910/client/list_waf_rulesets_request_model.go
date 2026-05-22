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
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// http_bot
	Phase     *string                          `json:"Phase,omitempty" xml:"Phase,omitempty"`
	QueryArgs *ListWafRulesetsRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
	// example:
	//
	// example
	AnyLike *string `json:"AnyLike,omitempty" xml:"AnyLike,omitempty"`
	Desc    *bool   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// example
	NameLike *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
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
