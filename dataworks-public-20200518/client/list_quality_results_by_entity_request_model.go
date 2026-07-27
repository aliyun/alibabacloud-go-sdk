// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *ListQualityResultsByEntityRequest
	GetEndDate() *string
	SetEntityId(v int64) *ListQualityResultsByEntityRequest
	GetEntityId() *int64
	SetPageNumber(v int32) *ListQualityResultsByEntityRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQualityResultsByEntityRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListQualityResultsByEntityRequest
	GetProjectId() *int64
	SetProjectName(v string) *ListQualityResultsByEntityRequest
	GetProjectName() *string
	SetStartDate(v string) *ListQualityResultsByEntityRequest
	GetStartDate() *string
}

type ListQualityResultsByEntityRequest struct {
	// The end of the business date range. The value must be in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-21 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the partition expression. You can call the [GetQualityEntity](https://help.aliyun.com/document_detail/174003.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15232
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The page number to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. The default is 10, and the maximum is 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the engine or data source. You can find this name on the Data Source Configuration page.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The start of the business date range. The value must be in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-20 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListQualityResultsByEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityRequest) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListQualityResultsByEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListQualityResultsByEntityRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityResultsByEntityRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityResultsByEntityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListQualityResultsByEntityRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityResultsByEntityRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListQualityResultsByEntityRequest) SetEndDate(v string) *ListQualityResultsByEntityRequest {
	s.EndDate = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetEntityId(v int64) *ListQualityResultsByEntityRequest {
	s.EntityId = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetPageNumber(v int32) *ListQualityResultsByEntityRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetPageSize(v int32) *ListQualityResultsByEntityRequest {
	s.PageSize = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetProjectId(v int64) *ListQualityResultsByEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetProjectName(v string) *ListQualityResultsByEntityRequest {
	s.ProjectName = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) SetStartDate(v string) *ListQualityResultsByEntityRequest {
	s.StartDate = &v
	return s
}

func (s *ListQualityResultsByEntityRequest) Validate() error {
	return dara.Validate(s)
}
