// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryUuids(v []*string) *ListDatasetDocumentsRequest
	GetCategoryUuids() []*string
	SetCreateTimeEnd(v int64) *ListDatasetDocumentsRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *ListDatasetDocumentsRequest
	GetCreateTimeStart() *int64
	SetDatasetDescription(v string) *ListDatasetDocumentsRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *ListDatasetDocumentsRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *ListDatasetDocumentsRequest
	GetDatasetName() *string
	SetDocIds(v []*string) *ListDatasetDocumentsRequest
	GetDocIds() []*string
	SetDocType(v string) *ListDatasetDocumentsRequest
	GetDocType() *string
	SetDocUuids(v []*string) *ListDatasetDocumentsRequest
	GetDocUuids() []*string
	SetEndTime(v int64) *ListDatasetDocumentsRequest
	GetEndTime() *int64
	SetExcludeFields(v []*string) *ListDatasetDocumentsRequest
	GetExcludeFields() []*string
	SetExtend1(v string) *ListDatasetDocumentsRequest
	GetExtend1() *string
	SetExtend2(v string) *ListDatasetDocumentsRequest
	GetExtend2() *string
	SetExtend3(v string) *ListDatasetDocumentsRequest
	GetExtend3() *string
	SetIncludeFields(v []*string) *ListDatasetDocumentsRequest
	GetIncludeFields() []*string
	SetPageNumber(v int32) *ListDatasetDocumentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetDocumentsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListDatasetDocumentsRequest
	GetQuery() *string
	SetStartTime(v int64) *ListDatasetDocumentsRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListDatasetDocumentsRequest
	GetStatus() *int32
	SetTags(v []*string) *ListDatasetDocumentsRequest
	GetTags() []*string
	SetTitle(v string) *ListDatasetDocumentsRequest
	GetTitle() *string
	SetWorkspaceId(v string) *ListDatasetDocumentsRequest
	GetWorkspaceId() *string
}

type ListDatasetDocumentsRequest struct {
	CategoryUuids   []*string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty" type:"Repeated"`
	CreateTimeEnd   *int64    `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *int64    `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// xx
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string   `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DocIds      []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// example:
	//
	// text
	DocType       *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocUuids      []*string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty" type:"Repeated"`
	EndTime       *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExcludeFields []*string `json:"ExcludeFields,omitempty" xml:"ExcludeFields,omitempty" type:"Repeated"`
	Extend1       *string   `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2       *string   `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3       *string   `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	IncludeFields []*string `json:"IncludeFields,omitempty" xml:"IncludeFields,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 搜索条件
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 100
	Status *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Title  *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsRequest) GetCategoryUuids() []*string {
	return s.CategoryUuids
}

func (s *ListDatasetDocumentsRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *ListDatasetDocumentsRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *ListDatasetDocumentsRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *ListDatasetDocumentsRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetDocumentsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetDocumentsRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *ListDatasetDocumentsRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDatasetDocumentsRequest) GetDocUuids() []*string {
	return s.DocUuids
}

func (s *ListDatasetDocumentsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDatasetDocumentsRequest) GetExcludeFields() []*string {
	return s.ExcludeFields
}

func (s *ListDatasetDocumentsRequest) GetExtend1() *string {
	return s.Extend1
}

func (s *ListDatasetDocumentsRequest) GetExtend2() *string {
	return s.Extend2
}

func (s *ListDatasetDocumentsRequest) GetExtend3() *string {
	return s.Extend3
}

func (s *ListDatasetDocumentsRequest) GetIncludeFields() []*string {
	return s.IncludeFields
}

func (s *ListDatasetDocumentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetDocumentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetDocumentsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListDatasetDocumentsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDatasetDocumentsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDatasetDocumentsRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListDatasetDocumentsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListDatasetDocumentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetDocumentsRequest) SetCategoryUuids(v []*string) *ListDatasetDocumentsRequest {
	s.CategoryUuids = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetCreateTimeEnd(v int64) *ListDatasetDocumentsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetCreateTimeStart(v int64) *ListDatasetDocumentsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDatasetDescription(v string) *ListDatasetDocumentsRequest {
	s.DatasetDescription = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDatasetId(v int64) *ListDatasetDocumentsRequest {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDatasetName(v string) *ListDatasetDocumentsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDocIds(v []*string) *ListDatasetDocumentsRequest {
	s.DocIds = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDocType(v string) *ListDatasetDocumentsRequest {
	s.DocType = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetDocUuids(v []*string) *ListDatasetDocumentsRequest {
	s.DocUuids = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetEndTime(v int64) *ListDatasetDocumentsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetExcludeFields(v []*string) *ListDatasetDocumentsRequest {
	s.ExcludeFields = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetExtend1(v string) *ListDatasetDocumentsRequest {
	s.Extend1 = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetExtend2(v string) *ListDatasetDocumentsRequest {
	s.Extend2 = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetExtend3(v string) *ListDatasetDocumentsRequest {
	s.Extend3 = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetIncludeFields(v []*string) *ListDatasetDocumentsRequest {
	s.IncludeFields = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetPageNumber(v int32) *ListDatasetDocumentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetPageSize(v int32) *ListDatasetDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetQuery(v string) *ListDatasetDocumentsRequest {
	s.Query = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetStartTime(v int64) *ListDatasetDocumentsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetStatus(v int32) *ListDatasetDocumentsRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetTags(v []*string) *ListDatasetDocumentsRequest {
	s.Tags = v
	return s
}

func (s *ListDatasetDocumentsRequest) SetTitle(v string) *ListDatasetDocumentsRequest {
	s.Title = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetWorkspaceId(v string) *ListDatasetDocumentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
