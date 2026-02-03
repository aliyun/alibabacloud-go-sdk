// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryUuidsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetCategoryUuidsShrink() *string
	SetCreateTimeEnd(v int64) *ListDatasetDocumentsShrinkRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *ListDatasetDocumentsShrinkRequest
	GetCreateTimeStart() *int64
	SetDatasetDescription(v string) *ListDatasetDocumentsShrinkRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *ListDatasetDocumentsShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *ListDatasetDocumentsShrinkRequest
	GetDatasetName() *string
	SetDocIdsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetDocIdsShrink() *string
	SetDocType(v string) *ListDatasetDocumentsShrinkRequest
	GetDocType() *string
	SetDocUuidsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetDocUuidsShrink() *string
	SetEndTime(v int64) *ListDatasetDocumentsShrinkRequest
	GetEndTime() *int64
	SetExcludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetExcludeFieldsShrink() *string
	SetExtend1(v string) *ListDatasetDocumentsShrinkRequest
	GetExtend1() *string
	SetExtend2(v string) *ListDatasetDocumentsShrinkRequest
	GetExtend2() *string
	SetExtend3(v string) *ListDatasetDocumentsShrinkRequest
	GetExtend3() *string
	SetIncludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetIncludeFieldsShrink() *string
	SetPageNumber(v int32) *ListDatasetDocumentsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetDocumentsShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *ListDatasetDocumentsShrinkRequest
	GetQuery() *string
	SetStartTime(v int64) *ListDatasetDocumentsShrinkRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListDatasetDocumentsShrinkRequest
	GetStatus() *int32
	SetTagsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetTagsShrink() *string
	SetTitle(v string) *ListDatasetDocumentsShrinkRequest
	GetTitle() *string
	SetWorkspaceId(v string) *ListDatasetDocumentsShrinkRequest
	GetWorkspaceId() *string
}

type ListDatasetDocumentsShrinkRequest struct {
	CategoryUuidsShrink *string `json:"CategoryUuids,omitempty" xml:"CategoryUuids,omitempty"`
	CreateTimeEnd       *int64  `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart     *int64  `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
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
	DatasetName  *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DocIdsShrink *string `json:"DocIds,omitempty" xml:"DocIds,omitempty"`
	// example:
	//
	// text
	DocType             *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocUuidsShrink      *string `json:"DocUuids,omitempty" xml:"DocUuids,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExcludeFieldsShrink *string `json:"ExcludeFields,omitempty" xml:"ExcludeFields,omitempty"`
	Extend1             *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	Extend2             *string `json:"Extend2,omitempty" xml:"Extend2,omitempty"`
	Extend3             *string `json:"Extend3,omitempty" xml:"Extend3,omitempty"`
	IncludeFieldsShrink *string `json:"IncludeFields,omitempty" xml:"IncludeFields,omitempty"`
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
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetDocumentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetDocumentsShrinkRequest) GetCategoryUuidsShrink() *string {
	return s.CategoryUuidsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *ListDatasetDocumentsShrinkRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetDocumentsShrinkRequest) GetDocIdsShrink() *string {
	return s.DocIdsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDatasetDocumentsShrinkRequest) GetDocUuidsShrink() *string {
	return s.DocUuidsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDatasetDocumentsShrinkRequest) GetExcludeFieldsShrink() *string {
	return s.ExcludeFieldsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetExtend1() *string {
	return s.Extend1
}

func (s *ListDatasetDocumentsShrinkRequest) GetExtend2() *string {
	return s.Extend2
}

func (s *ListDatasetDocumentsShrinkRequest) GetExtend3() *string {
	return s.Extend3
}

func (s *ListDatasetDocumentsShrinkRequest) GetIncludeFieldsShrink() *string {
	return s.IncludeFieldsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetDocumentsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetDocumentsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListDatasetDocumentsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDatasetDocumentsShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDatasetDocumentsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListDatasetDocumentsShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *ListDatasetDocumentsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetDocumentsShrinkRequest) SetCategoryUuidsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.CategoryUuidsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetCreateTimeEnd(v int64) *ListDatasetDocumentsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetCreateTimeStart(v int64) *ListDatasetDocumentsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDatasetDescription(v string) *ListDatasetDocumentsShrinkRequest {
	s.DatasetDescription = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDatasetId(v int64) *ListDatasetDocumentsShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDatasetName(v string) *ListDatasetDocumentsShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDocIdsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.DocIdsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDocType(v string) *ListDatasetDocumentsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetDocUuidsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.DocUuidsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetEndTime(v int64) *ListDatasetDocumentsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetExcludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.ExcludeFieldsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetExtend1(v string) *ListDatasetDocumentsShrinkRequest {
	s.Extend1 = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetExtend2(v string) *ListDatasetDocumentsShrinkRequest {
	s.Extend2 = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetExtend3(v string) *ListDatasetDocumentsShrinkRequest {
	s.Extend3 = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetIncludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.IncludeFieldsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetPageNumber(v int32) *ListDatasetDocumentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetPageSize(v int32) *ListDatasetDocumentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetQuery(v string) *ListDatasetDocumentsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetStartTime(v int64) *ListDatasetDocumentsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetStatus(v int32) *ListDatasetDocumentsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetTagsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetTitle(v string) *ListDatasetDocumentsShrinkRequest {
	s.Title = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetWorkspaceId(v string) *ListDatasetDocumentsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
