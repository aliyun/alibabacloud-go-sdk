// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileSummaryInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetBizId() *string
	SetDirectoryId(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetDirectoryId() *string
	SetName(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetName() *string
	SetOrderColumn(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialFileSummaryInfoShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialFileSummaryInfoShrinkRequest
	GetPageSize() *int32
	SetStatusListShrink(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetStatusListShrink() *string
	SetTypeListShrink(v string) *QueryMaterialFileSummaryInfoShrinkRequest
	GetTypeListShrink() *string
}

type QueryMaterialFileSummaryInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	TypeListShrink   *string `json:"TypeList,omitempty" xml:"TypeList,omitempty"`
}

func (s QueryMaterialFileSummaryInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileSummaryInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetName() *string {
	return s.Name
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) GetTypeListShrink() *string {
	return s.TypeListShrink
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetBizId(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetDirectoryId(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetName(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.Name = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetOrderColumn(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetOrderType(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetPageNum(v int32) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetPageSize(v int32) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetStatusListShrink(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) SetTypeListShrink(v string) *QueryMaterialFileSummaryInfoShrinkRequest {
	s.TypeListShrink = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
