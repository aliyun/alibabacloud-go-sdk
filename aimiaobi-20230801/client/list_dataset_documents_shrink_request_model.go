// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetDescription(v string) *ListDatasetDocumentsShrinkRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *ListDatasetDocumentsShrinkRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *ListDatasetDocumentsShrinkRequest
	GetDatasetName() *string
	SetDocType(v string) *ListDatasetDocumentsShrinkRequest
	GetDocType() *string
	SetExcludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetExcludeFieldsShrink() *string
	SetIncludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest
	GetIncludeFieldsShrink() *string
	SetPageNumber(v int32) *ListDatasetDocumentsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetDocumentsShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *ListDatasetDocumentsShrinkRequest
	GetQuery() *string
	SetStatus(v int32) *ListDatasetDocumentsShrinkRequest
	GetStatus() *int32
	SetWorkspaceId(v string) *ListDatasetDocumentsShrinkRequest
	GetWorkspaceId() *string
}

type ListDatasetDocumentsShrinkRequest struct {
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
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// text
	DocType             *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExcludeFieldsShrink *string `json:"ExcludeFields,omitempty" xml:"ExcludeFields,omitempty"`
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
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 100
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetDocumentsShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetDocumentsShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDatasetDocumentsShrinkRequest) GetExcludeFieldsShrink() *string {
	return s.ExcludeFieldsShrink
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

func (s *ListDatasetDocumentsShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDatasetDocumentsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *ListDatasetDocumentsShrinkRequest) SetDocType(v string) *ListDatasetDocumentsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetExcludeFieldsShrink(v string) *ListDatasetDocumentsShrinkRequest {
	s.ExcludeFieldsShrink = &v
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

func (s *ListDatasetDocumentsShrinkRequest) SetStatus(v int32) *ListDatasetDocumentsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) SetWorkspaceId(v string) *ListDatasetDocumentsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
