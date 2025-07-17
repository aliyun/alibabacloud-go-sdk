// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityEvaluationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListDataQualityEvaluationTasksRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataQualityEvaluationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityEvaluationTasksRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityEvaluationTasksRequest
	GetProjectId() *int64
	SetTableGuid(v string) *ListDataQualityEvaluationTasksRequest
	GetTableGuid() *string
}

type ListDataQualityEvaluationTasksRequest struct {
	// The name of the data quality monitoring task. Fuzzy match is supported.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the table in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s ListDataQualityEvaluationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksRequest) GetName() *string {
	return s.Name
}

func (s *ListDataQualityEvaluationTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityEvaluationTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityEvaluationTasksRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityEvaluationTasksRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityEvaluationTasksRequest) SetName(v string) *ListDataQualityEvaluationTasksRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetPageNumber(v int32) *ListDataQualityEvaluationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetPageSize(v int32) *ListDataQualityEvaluationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetProjectId(v int64) *ListDataQualityEvaluationTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetTableGuid(v string) *ListDataQualityEvaluationTasksRequest {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) Validate() error {
	return dara.Validate(s)
}
