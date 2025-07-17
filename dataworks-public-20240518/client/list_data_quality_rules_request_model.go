// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskId(v int64) *ListDataQualityRulesRequest
	GetDataQualityEvaluationTaskId() *int64
	SetName(v string) *ListDataQualityRulesRequest
	GetName() *string
	SetPageNumber(v int32) *ListDataQualityRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityRulesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityRulesRequest
	GetProjectId() *int64
	SetTableGuid(v string) *ListDataQualityRulesRequest
	GetTableGuid() *string
}

type ListDataQualityRulesRequest struct {
	// The ID of the data quality monitoring task that is associated with the rule.
	//
	// example:
	//
	// 10000
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// The name of the rule. Fuzzy match is supported.
	//
	// example:
	//
	// unit_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 200.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the table that is limited by the rule in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s ListDataQualityRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesRequest) GetDataQualityEvaluationTaskId() *int64 {
	return s.DataQualityEvaluationTaskId
}

func (s *ListDataQualityRulesRequest) GetName() *string {
	return s.Name
}

func (s *ListDataQualityRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityRulesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityRulesRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityRulesRequest) SetDataQualityEvaluationTaskId(v int64) *ListDataQualityRulesRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetName(v string) *ListDataQualityRulesRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetPageNumber(v int32) *ListDataQualityRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetPageSize(v int32) *ListDataQualityRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetProjectId(v int64) *ListDataQualityRulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetTableGuid(v string) *ListDataQualityRulesRequest {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityRulesRequest) Validate() error {
	return dara.Validate(s)
}
