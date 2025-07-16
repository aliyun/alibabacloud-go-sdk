// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *ListMultiDimTableRecordsShrinkRequest
	GetBaseId() *string
	SetFilterShrink(v string) *ListMultiDimTableRecordsShrinkRequest
	GetFilterShrink() *string
	SetMaxResults(v int32) *ListMultiDimTableRecordsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiDimTableRecordsShrinkRequest
	GetNextToken() *string
	SetSheetIdOrName(v string) *ListMultiDimTableRecordsShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *ListMultiDimTableRecordsShrinkRequest
	GetTenantContextShrink() *string
}

type ListMultiDimTableRecordsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101114
	BaseId       *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUUg5QSTWwHyeElt8z5z4Qo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ListMultiDimTableRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *ListMultiDimTableRecordsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetBaseId(v string) *ListMultiDimTableRecordsShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetFilterShrink(v string) *ListMultiDimTableRecordsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetMaxResults(v int32) *ListMultiDimTableRecordsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetNextToken(v string) *ListMultiDimTableRecordsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetSheetIdOrName(v string) *ListMultiDimTableRecordsShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) SetTenantContextShrink(v string) *ListMultiDimTableRecordsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListMultiDimTableRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
