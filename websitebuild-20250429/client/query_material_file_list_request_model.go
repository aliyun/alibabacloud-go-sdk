// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialFileListRequest
	GetBizId() *string
	SetDirectoryId(v string) *QueryMaterialFileListRequest
	GetDirectoryId() *string
	SetMaxFileSize(v int64) *QueryMaterialFileListRequest
	GetMaxFileSize() *int64
	SetMaxResults(v int32) *QueryMaterialFileListRequest
	GetMaxResults() *int32
	SetMinFileSize(v int64) *QueryMaterialFileListRequest
	GetMinFileSize() *int64
	SetName(v string) *QueryMaterialFileListRequest
	GetName() *string
	SetNextToken(v string) *QueryMaterialFileListRequest
	GetNextToken() *string
	SetOrderColumn(v string) *QueryMaterialFileListRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialFileListRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialFileListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialFileListRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *QueryMaterialFileListRequest
	GetStatusList() []*string
	SetSuffixList(v []*string) *QueryMaterialFileListRequest
	GetSuffixList() []*string
	SetTypeList(v []*string) *QueryMaterialFileListRequest
	GetTypeList() []*string
}

type QueryMaterialFileListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 1024
	MaxFileSize *int64 `json:"MaxFileSize,omitempty" xml:"MaxFileSize,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0
	MinFileSize *int64  `json:"MinFileSize,omitempty" xml:"MinFileSize,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	SuffixList []*string `json:"SuffixList,omitempty" xml:"SuffixList,omitempty" type:"Repeated"`
	TypeList   []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s QueryMaterialFileListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileListRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileListRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialFileListRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *QueryMaterialFileListRequest) GetMaxFileSize() *int64 {
	return s.MaxFileSize
}

func (s *QueryMaterialFileListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialFileListRequest) GetMinFileSize() *int64 {
	return s.MinFileSize
}

func (s *QueryMaterialFileListRequest) GetName() *string {
	return s.Name
}

func (s *QueryMaterialFileListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialFileListRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialFileListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialFileListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialFileListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialFileListRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QueryMaterialFileListRequest) GetSuffixList() []*string {
	return s.SuffixList
}

func (s *QueryMaterialFileListRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *QueryMaterialFileListRequest) SetBizId(v string) *QueryMaterialFileListRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetDirectoryId(v string) *QueryMaterialFileListRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetMaxFileSize(v int64) *QueryMaterialFileListRequest {
	s.MaxFileSize = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetMaxResults(v int32) *QueryMaterialFileListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetMinFileSize(v int64) *QueryMaterialFileListRequest {
	s.MinFileSize = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetName(v string) *QueryMaterialFileListRequest {
	s.Name = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetNextToken(v string) *QueryMaterialFileListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetOrderColumn(v string) *QueryMaterialFileListRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetOrderType(v string) *QueryMaterialFileListRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetPageNum(v int32) *QueryMaterialFileListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetPageSize(v int32) *QueryMaterialFileListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialFileListRequest) SetStatusList(v []*string) *QueryMaterialFileListRequest {
	s.StatusList = v
	return s
}

func (s *QueryMaterialFileListRequest) SetSuffixList(v []*string) *QueryMaterialFileListRequest {
	s.SuffixList = v
	return s
}

func (s *QueryMaterialFileListRequest) SetTypeList(v []*string) *QueryMaterialFileListRequest {
	s.TypeList = v
	return s
}

func (s *QueryMaterialFileListRequest) Validate() error {
	return dara.Validate(s)
}
