// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileSummaryInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialFileSummaryInfoRequest
	GetBizId() *string
	SetDirectoryId(v string) *QueryMaterialFileSummaryInfoRequest
	GetDirectoryId() *string
	SetName(v string) *QueryMaterialFileSummaryInfoRequest
	GetName() *string
	SetOrderColumn(v string) *QueryMaterialFileSummaryInfoRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialFileSummaryInfoRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialFileSummaryInfoRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialFileSummaryInfoRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *QueryMaterialFileSummaryInfoRequest
	GetStatusList() []*string
	SetTypeList(v []*string) *QueryMaterialFileSummaryInfoRequest
	GetTypeList() []*string
}

type QueryMaterialFileSummaryInfoRequest struct {
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
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TypeList   []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s QueryMaterialFileSummaryInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileSummaryInfoRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialFileSummaryInfoRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *QueryMaterialFileSummaryInfoRequest) GetName() *string {
	return s.Name
}

func (s *QueryMaterialFileSummaryInfoRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialFileSummaryInfoRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialFileSummaryInfoRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialFileSummaryInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialFileSummaryInfoRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryMaterialFileSummaryInfoRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *QueryMaterialFileSummaryInfoRequest) SetBizId(v string) *QueryMaterialFileSummaryInfoRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetDirectoryId(v string) *QueryMaterialFileSummaryInfoRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetName(v string) *QueryMaterialFileSummaryInfoRequest {
	s.Name = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetOrderColumn(v string) *QueryMaterialFileSummaryInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetOrderType(v string) *QueryMaterialFileSummaryInfoRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetPageNum(v int32) *QueryMaterialFileSummaryInfoRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetPageSize(v int32) *QueryMaterialFileSummaryInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetStatusList(v []*string) *QueryMaterialFileSummaryInfoRequest {
	s.StatusList = v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) SetTypeList(v []*string) *QueryMaterialFileSummaryInfoRequest {
	s.TypeList = v
	return s
}

func (s *QueryMaterialFileSummaryInfoRequest) Validate() error {
	return dara.Validate(s)
}
