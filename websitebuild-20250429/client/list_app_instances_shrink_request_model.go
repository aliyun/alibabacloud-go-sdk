// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppInstancesShrinkRequest
	GetBizId() *string
	SetEndTimeBegin(v string) *ListAppInstancesShrinkRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *ListAppInstancesShrinkRequest
	GetEndTimeEnd() *string
	SetExtend(v string) *ListAppInstancesShrinkRequest
	GetExtend() *string
	SetMaxResults(v int32) *ListAppInstancesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppInstancesShrinkRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListAppInstancesShrinkRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListAppInstancesShrinkRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListAppInstancesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppInstancesShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *ListAppInstancesShrinkRequest
	GetQuery() *string
	SetStatusListShrink(v string) *ListAppInstancesShrinkRequest
	GetStatusListShrink() *string
}

type ListAppInstancesShrinkRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 2025-12-18T14:30:00Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2025-12-31T14:30:00Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// DOWNGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {\\"RequestId\\":\\"\\"}
	Query            *string `json:"Query,omitempty" xml:"Query,omitempty"`
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
}

func (s ListAppInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstancesShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppInstancesShrinkRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *ListAppInstancesShrinkRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *ListAppInstancesShrinkRequest) GetExtend() *string {
	return s.Extend
}

func (s *ListAppInstancesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppInstancesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppInstancesShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListAppInstancesShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAppInstancesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstancesShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAppInstancesShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *ListAppInstancesShrinkRequest) SetBizId(v string) *ListAppInstancesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetEndTimeBegin(v string) *ListAppInstancesShrinkRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetEndTimeEnd(v string) *ListAppInstancesShrinkRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetExtend(v string) *ListAppInstancesShrinkRequest {
	s.Extend = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetMaxResults(v int32) *ListAppInstancesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetNextToken(v string) *ListAppInstancesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetOrderColumn(v string) *ListAppInstancesShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetOrderType(v string) *ListAppInstancesShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetPageNum(v int32) *ListAppInstancesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetPageSize(v int32) *ListAppInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetQuery(v string) *ListAppInstancesShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) SetStatusListShrink(v string) *ListAppInstancesShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *ListAppInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
