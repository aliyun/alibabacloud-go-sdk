// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetDescription(v string) *ListDatasetDocumentsRequest
	GetDatasetDescription() *string
	SetDatasetId(v int64) *ListDatasetDocumentsRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *ListDatasetDocumentsRequest
	GetDatasetName() *string
	SetDocType(v string) *ListDatasetDocumentsRequest
	GetDocType() *string
	SetExcludeFields(v []*string) *ListDatasetDocumentsRequest
	GetExcludeFields() []*string
	SetIncludeFields(v []*string) *ListDatasetDocumentsRequest
	GetIncludeFields() []*string
	SetPageNumber(v int32) *ListDatasetDocumentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetDocumentsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListDatasetDocumentsRequest
	GetQuery() *string
	SetStatus(v int32) *ListDatasetDocumentsRequest
	GetStatus() *int32
	SetWorkspaceId(v string) *ListDatasetDocumentsRequest
	GetWorkspaceId() *string
}

type ListDatasetDocumentsRequest struct {
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
	DocType       *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExcludeFields []*string `json:"ExcludeFields,omitempty" xml:"ExcludeFields,omitempty" type:"Repeated"`
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

func (s ListDatasetDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetDocumentsRequest) GoString() string {
	return s.String()
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

func (s *ListDatasetDocumentsRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListDatasetDocumentsRequest) GetExcludeFields() []*string {
	return s.ExcludeFields
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

func (s *ListDatasetDocumentsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDatasetDocumentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *ListDatasetDocumentsRequest) SetDocType(v string) *ListDatasetDocumentsRequest {
	s.DocType = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetExcludeFields(v []*string) *ListDatasetDocumentsRequest {
	s.ExcludeFields = v
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

func (s *ListDatasetDocumentsRequest) SetStatus(v int32) *ListDatasetDocumentsRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetDocumentsRequest) SetWorkspaceId(v string) *ListDatasetDocumentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
