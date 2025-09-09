// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionZonesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *SearchRecursionZonesShrinkRequest
	GetDirection() *string
	SetEffectiveScopesShrink(v string) *SearchRecursionZonesShrinkRequest
	GetEffectiveScopesShrink() *string
	SetMaxResults(v int32) *SearchRecursionZonesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionZonesShrinkRequest
	GetNextToken() *string
	SetOrderBy(v string) *SearchRecursionZonesShrinkRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *SearchRecursionZonesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionZonesShrinkRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchRecursionZonesShrinkRequest
	GetRemark() *string
	SetZoneName(v string) *SearchRecursionZonesShrinkRequest
	GetZoneName() *string
}

type SearchRecursionZonesShrinkRequest struct {
	// example:
	//
	// asc
	Direction             *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EffectiveScopesShrink *string `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// cheng.suow.cc
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s SearchRecursionZonesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *SearchRecursionZonesShrinkRequest) GetEffectiveScopesShrink() *string {
	return s.EffectiveScopesShrink
}

func (s *SearchRecursionZonesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionZonesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionZonesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchRecursionZonesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionZonesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionZonesShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionZonesShrinkRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *SearchRecursionZonesShrinkRequest) SetDirection(v string) *SearchRecursionZonesShrinkRequest {
	s.Direction = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetEffectiveScopesShrink(v string) *SearchRecursionZonesShrinkRequest {
	s.EffectiveScopesShrink = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetMaxResults(v int32) *SearchRecursionZonesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetNextToken(v string) *SearchRecursionZonesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetOrderBy(v string) *SearchRecursionZonesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetPageNumber(v int32) *SearchRecursionZonesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetPageSize(v int32) *SearchRecursionZonesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetRemark(v string) *SearchRecursionZonesShrinkRequest {
	s.Remark = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) SetZoneName(v string) *SearchRecursionZonesShrinkRequest {
	s.ZoneName = &v
	return s
}

func (s *SearchRecursionZonesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
