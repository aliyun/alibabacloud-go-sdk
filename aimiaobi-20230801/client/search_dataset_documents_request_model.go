// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryUuids(v []*string) *SearchDatasetDocumentsRequest
	GetCategoryUuids() []*string
	SetCreateTimeEnd(v int64) *SearchDatasetDocumentsRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *SearchDatasetDocumentsRequest
	GetCreateTimeStart() *int64
	SetDatasetId(v int64) *SearchDatasetDocumentsRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *SearchDatasetDocumentsRequest
	GetDatasetName() *string
	SetDocIds(v []*string) *SearchDatasetDocumentsRequest
	GetDocIds() []*string
	SetDocTypes(v []*string) *SearchDatasetDocumentsRequest
	GetDocTypes() []*string
	SetDocUuids(v []*string) *SearchDatasetDocumentsRequest
	GetDocUuids() []*string
	SetEndTime(v int64) *SearchDatasetDocumentsRequest
	GetEndTime() *int64
	SetExtend1(v string) *SearchDatasetDocumentsRequest
	GetExtend1() *string
	SetExtend2(v string) *SearchDatasetDocumentsRequest
	GetExtend2() *string
	SetExtend3(v string) *SearchDatasetDocumentsRequest
	GetExtend3() *string
	SetIncludeContent(v bool) *SearchDatasetDocumentsRequest
	GetIncludeContent() *bool
	SetPageSize(v string) *SearchDatasetDocumentsRequest
	GetPageSize() *string
	SetQuery(v string) *SearchDatasetDocumentsRequest
	GetQuery() *string
	SetSearchMode(v string) *SearchDatasetDocumentsRequest
	GetSearchMode() *string
	SetStartTime(v int64) *SearchDatasetDocumentsRequest
	GetStartTime() *int64
	SetTags(v []*string) *SearchDatasetDocumentsRequest
	GetTags() []*string
	SetWorkspaceId(v string) *SearchDatasetDocumentsRequest
	GetWorkspaceId() *string
}

type SearchDatasetDocumentsRequest struct {
	// A list of category UUIDs for filtering.
	CategoryUuids []*string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty" type:"Repeated"`
	// The end time for document creation, specified as a Unix timestamp.
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time for document creation, specified as a Unix timestamp.
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The ID of the dataset. You must specify either `DatasetId` or `DatasetName`.
	//
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the dataset.
	//
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// A list of document IDs for filtering.
	DocIds []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// A list of document types for filtering.
	DocTypes []*string `json:"DocTypes,omitempty" xml:"DocTypes,omitempty" type:"Repeated"`
	// A list of document UUIDs for filtering.
	DocUuids []*string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty" type:"Repeated"`
	// The end of the time range for filtering, specified as a Unix timestamp.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A business parameter.
	//
	// example:
	//
	// 业务参数
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// A reserved business parameter.
	Extend2 *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	// A reserved business parameter.
	Extend3 *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	// Specifies whether to include the document content in the search results. The default value is `false`.
	//
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// The number of documents to return per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 搜索内容
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The search mode.
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start of the time range for filtering, specified as a Unix timestamp.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of tags for filtering.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Model Studio workspace. For more information, see [Get workspaceId](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchDatasetDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsRequest) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsRequest) GetCategoryUuids() []*string {
	return s.CategoryUuids
}

func (s *SearchDatasetDocumentsRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *SearchDatasetDocumentsRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *SearchDatasetDocumentsRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SearchDatasetDocumentsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SearchDatasetDocumentsRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *SearchDatasetDocumentsRequest) GetDocTypes() []*string {
	return s.DocTypes
}

func (s *SearchDatasetDocumentsRequest) GetDocUuids() []*string {
	return s.DocUuids
}

func (s *SearchDatasetDocumentsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchDatasetDocumentsRequest) GetExtend1() *string {
	return s.Extend1
}

func (s *SearchDatasetDocumentsRequest) GetExtend2() *string {
	return s.Extend2
}

func (s *SearchDatasetDocumentsRequest) GetExtend3() *string {
	return s.Extend3
}

func (s *SearchDatasetDocumentsRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *SearchDatasetDocumentsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *SearchDatasetDocumentsRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchDatasetDocumentsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *SearchDatasetDocumentsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchDatasetDocumentsRequest) GetTags() []*string {
	return s.Tags
}

func (s *SearchDatasetDocumentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchDatasetDocumentsRequest) SetCategoryUuids(v []*string) *SearchDatasetDocumentsRequest {
	s.CategoryUuids = v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetCreateTimeEnd(v int64) *SearchDatasetDocumentsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetCreateTimeStart(v int64) *SearchDatasetDocumentsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDatasetId(v int64) *SearchDatasetDocumentsRequest {
	s.DatasetId = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDatasetName(v string) *SearchDatasetDocumentsRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDocIds(v []*string) *SearchDatasetDocumentsRequest {
	s.DocIds = v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDocTypes(v []*string) *SearchDatasetDocumentsRequest {
	s.DocTypes = v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDocUuids(v []*string) *SearchDatasetDocumentsRequest {
	s.DocUuids = v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetEndTime(v int64) *SearchDatasetDocumentsRequest {
	s.EndTime = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetExtend1(v string) *SearchDatasetDocumentsRequest {
	s.Extend1 = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetExtend2(v string) *SearchDatasetDocumentsRequest {
	s.Extend2 = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetExtend3(v string) *SearchDatasetDocumentsRequest {
	s.Extend3 = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetIncludeContent(v bool) *SearchDatasetDocumentsRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetPageSize(v string) *SearchDatasetDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetQuery(v string) *SearchDatasetDocumentsRequest {
	s.Query = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetSearchMode(v string) *SearchDatasetDocumentsRequest {
	s.SearchMode = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetStartTime(v int64) *SearchDatasetDocumentsRequest {
	s.StartTime = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetTags(v []*string) *SearchDatasetDocumentsRequest {
	s.Tags = v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetWorkspaceId(v string) *SearchDatasetDocumentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
