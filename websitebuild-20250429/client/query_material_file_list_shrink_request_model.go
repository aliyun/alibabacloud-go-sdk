// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialFileListShrinkRequest
	GetBizId() *string
	SetDirectoryId(v string) *QueryMaterialFileListShrinkRequest
	GetDirectoryId() *string
	SetMaxFileSize(v int64) *QueryMaterialFileListShrinkRequest
	GetMaxFileSize() *int64
	SetMaxResults(v int32) *QueryMaterialFileListShrinkRequest
	GetMaxResults() *int32
	SetMinFileSize(v int64) *QueryMaterialFileListShrinkRequest
	GetMinFileSize() *int64
	SetName(v string) *QueryMaterialFileListShrinkRequest
	GetName() *string
	SetNextToken(v string) *QueryMaterialFileListShrinkRequest
	GetNextToken() *string
	SetOrderColumn(v string) *QueryMaterialFileListShrinkRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryMaterialFileListShrinkRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryMaterialFileListShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryMaterialFileListShrinkRequest
	GetPageSize() *int32
	SetStatusListShrink(v string) *QueryMaterialFileListShrinkRequest
	GetStatusListShrink() *string
	SetSuffixListShrink(v string) *QueryMaterialFileListShrinkRequest
	GetSuffixListShrink() *string
	SetTypeListShrink(v string) *QueryMaterialFileListShrinkRequest
	GetTypeListShrink() *string
}

type QueryMaterialFileListShrinkRequest struct {
	// The business ID of the application instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum file size.
	//
	// example:
	//
	// 1024
	MaxFileSize *int64 `json:"MaxFileSize,omitempty" xml:"MaxFileSize,omitempty"`
	// The maximum number of entries to return per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The minimum file size.
	//
	// example:
	//
	// 0
	MinFileSize *int64 `json:"MinFileSize,omitempty" xml:"MinFileSize,omitempty"`
	// The website name.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of statuses.
	StatusListShrink *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The list of file suffixes.
	SuffixListShrink *string `json:"SuffixList,omitempty" xml:"SuffixList,omitempty"`
	// The list of file types.
	TypeListShrink *string `json:"TypeList,omitempty" xml:"TypeList,omitempty"`
}

func (s QueryMaterialFileListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileListShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialFileListShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *QueryMaterialFileListShrinkRequest) GetMaxFileSize() *int64 {
	return s.MaxFileSize
}

func (s *QueryMaterialFileListShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryMaterialFileListShrinkRequest) GetMinFileSize() *int64 {
	return s.MinFileSize
}

func (s *QueryMaterialFileListShrinkRequest) GetName() *string {
	return s.Name
}

func (s *QueryMaterialFileListShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryMaterialFileListShrinkRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryMaterialFileListShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryMaterialFileListShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryMaterialFileListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMaterialFileListShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *QueryMaterialFileListShrinkRequest) GetSuffixListShrink() *string {
	return s.SuffixListShrink
}

func (s *QueryMaterialFileListShrinkRequest) GetTypeListShrink() *string {
	return s.TypeListShrink
}

func (s *QueryMaterialFileListShrinkRequest) SetBizId(v string) *QueryMaterialFileListShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetDirectoryId(v string) *QueryMaterialFileListShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetMaxFileSize(v int64) *QueryMaterialFileListShrinkRequest {
	s.MaxFileSize = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetMaxResults(v int32) *QueryMaterialFileListShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetMinFileSize(v int64) *QueryMaterialFileListShrinkRequest {
	s.MinFileSize = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetName(v string) *QueryMaterialFileListShrinkRequest {
	s.Name = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetNextToken(v string) *QueryMaterialFileListShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetOrderColumn(v string) *QueryMaterialFileListShrinkRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetOrderType(v string) *QueryMaterialFileListShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetPageNum(v int32) *QueryMaterialFileListShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetPageSize(v int32) *QueryMaterialFileListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetStatusListShrink(v string) *QueryMaterialFileListShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetSuffixListShrink(v string) *QueryMaterialFileListShrinkRequest {
	s.SuffixListShrink = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) SetTypeListShrink(v string) *QueryMaterialFileListShrinkRequest {
	s.TypeListShrink = &v
	return s
}

func (s *QueryMaterialFileListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
