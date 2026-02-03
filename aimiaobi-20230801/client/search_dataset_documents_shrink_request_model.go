// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryUuidsShrink(v string) *SearchDatasetDocumentsShrinkRequest
	GetCategoryUuidsShrink() *string
	SetCreateTimeEnd(v int64) *SearchDatasetDocumentsShrinkRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *SearchDatasetDocumentsShrinkRequest
	GetCreateTimeStart() *int64
	SetDatasetId(v int64) *SearchDatasetDocumentsShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *SearchDatasetDocumentsShrinkRequest
	GetDatasetName() *string
	SetDocIdsShrink(v string) *SearchDatasetDocumentsShrinkRequest
	GetDocIdsShrink() *string
	SetDocTypesShrink(v string) *SearchDatasetDocumentsShrinkRequest
	GetDocTypesShrink() *string
	SetDocUuidsShrink(v string) *SearchDatasetDocumentsShrinkRequest
	GetDocUuidsShrink() *string
	SetEndTime(v int64) *SearchDatasetDocumentsShrinkRequest
	GetEndTime() *int64
	SetExtend1(v string) *SearchDatasetDocumentsShrinkRequest
	GetExtend1() *string
	SetExtend2(v string) *SearchDatasetDocumentsShrinkRequest
	GetExtend2() *string
	SetExtend3(v string) *SearchDatasetDocumentsShrinkRequest
	GetExtend3() *string
	SetIncludeContent(v bool) *SearchDatasetDocumentsShrinkRequest
	GetIncludeContent() *bool
	SetPageSize(v string) *SearchDatasetDocumentsShrinkRequest
	GetPageSize() *string
	SetQuery(v string) *SearchDatasetDocumentsShrinkRequest
	GetQuery() *string
	SetSearchMode(v string) *SearchDatasetDocumentsShrinkRequest
	GetSearchMode() *string
	SetStartTime(v int64) *SearchDatasetDocumentsShrinkRequest
	GetStartTime() *int64
	SetTagsShrink(v string) *SearchDatasetDocumentsShrinkRequest
	GetTagsShrink() *string
	SetWorkspaceId(v string) *SearchDatasetDocumentsShrinkRequest
	GetWorkspaceId() *string
}

type SearchDatasetDocumentsShrinkRequest struct {
	CategoryUuidsShrink *string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty"`
	CreateTimeEnd       *int64  `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart     *int64  `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName    *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DocIdsShrink   *string `json:"DocIds,omitempty" xml:"DocIds,omitempty"`
	DocTypesShrink *string `json:"DocTypes,omitempty" xml:"DocTypes,omitempty"`
	DocUuidsShrink *string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 业务参数
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 搜索内容
	Query      *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchDatasetDocumentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsShrinkRequest) GetCategoryUuidsShrink() *string {
	return s.CategoryUuidsShrink
}

func (s *SearchDatasetDocumentsShrinkRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *SearchDatasetDocumentsShrinkRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *SearchDatasetDocumentsShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SearchDatasetDocumentsShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SearchDatasetDocumentsShrinkRequest) GetDocIdsShrink() *string {
	return s.DocIdsShrink
}

func (s *SearchDatasetDocumentsShrinkRequest) GetDocTypesShrink() *string {
	return s.DocTypesShrink
}

func (s *SearchDatasetDocumentsShrinkRequest) GetDocUuidsShrink() *string {
	return s.DocUuidsShrink
}

func (s *SearchDatasetDocumentsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchDatasetDocumentsShrinkRequest) GetExtend1() *string {
	return s.Extend1
}

func (s *SearchDatasetDocumentsShrinkRequest) GetExtend2() *string {
	return s.Extend2
}

func (s *SearchDatasetDocumentsShrinkRequest) GetExtend3() *string {
	return s.Extend3
}

func (s *SearchDatasetDocumentsShrinkRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *SearchDatasetDocumentsShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *SearchDatasetDocumentsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchDatasetDocumentsShrinkRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *SearchDatasetDocumentsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchDatasetDocumentsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *SearchDatasetDocumentsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchDatasetDocumentsShrinkRequest) SetCategoryUuidsShrink(v string) *SearchDatasetDocumentsShrinkRequest {
	s.CategoryUuidsShrink = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetCreateTimeEnd(v int64) *SearchDatasetDocumentsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetCreateTimeStart(v int64) *SearchDatasetDocumentsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetDatasetId(v int64) *SearchDatasetDocumentsShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetDatasetName(v string) *SearchDatasetDocumentsShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetDocIdsShrink(v string) *SearchDatasetDocumentsShrinkRequest {
	s.DocIdsShrink = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetDocTypesShrink(v string) *SearchDatasetDocumentsShrinkRequest {
	s.DocTypesShrink = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetDocUuidsShrink(v string) *SearchDatasetDocumentsShrinkRequest {
	s.DocUuidsShrink = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetEndTime(v int64) *SearchDatasetDocumentsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetExtend1(v string) *SearchDatasetDocumentsShrinkRequest {
	s.Extend1 = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetExtend2(v string) *SearchDatasetDocumentsShrinkRequest {
	s.Extend2 = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetExtend3(v string) *SearchDatasetDocumentsShrinkRequest {
	s.Extend3 = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetIncludeContent(v bool) *SearchDatasetDocumentsShrinkRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetPageSize(v string) *SearchDatasetDocumentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetQuery(v string) *SearchDatasetDocumentsShrinkRequest {
	s.Query = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetSearchMode(v string) *SearchDatasetDocumentsShrinkRequest {
	s.SearchMode = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetStartTime(v int64) *SearchDatasetDocumentsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetTagsShrink(v string) *SearchDatasetDocumentsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) SetWorkspaceId(v string) *SearchDatasetDocumentsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchDatasetDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
