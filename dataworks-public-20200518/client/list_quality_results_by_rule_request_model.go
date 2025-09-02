// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *ListQualityResultsByRuleRequest
	GetEndDate() *string
	SetPageNumber(v int32) *ListQualityResultsByRuleRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQualityResultsByRuleRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListQualityResultsByRuleRequest
	GetProjectId() *int64
	SetProjectName(v string) *ListQualityResultsByRuleRequest
	GetProjectName() *string
	SetRuleId(v int64) *ListQualityResultsByRuleRequest
	GetRuleId() *int64
	SetStartDate(v string) *ListQualityResultsByRuleRequest
	GetStartDate() *string
}

type ListQualityResultsByRuleRequest struct {
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// This parameter is used together with the StartDate parameter. The interval between the time specified by this parameter and the time specified by the StartDate parameter cannot exceed 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-22 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source for which data quality is monitored.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The monitoring rule ID. You can use the ID and information such as a partition filter expression to perform a joint query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 152322134
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// This parameter is used together with the EndDate parameter. The interval between the time specified by this parameter and the time specified by the EndDate parameter cannot exceed 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-20 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListQualityResultsByRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByRuleRequest) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByRuleRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListQualityResultsByRuleRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityResultsByRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityResultsByRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListQualityResultsByRuleRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListQualityResultsByRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListQualityResultsByRuleRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListQualityResultsByRuleRequest) SetEndDate(v string) *ListQualityResultsByRuleRequest {
	s.EndDate = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetPageNumber(v int32) *ListQualityResultsByRuleRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetPageSize(v int32) *ListQualityResultsByRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetProjectId(v int64) *ListQualityResultsByRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetProjectName(v string) *ListQualityResultsByRuleRequest {
	s.ProjectName = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetRuleId(v int64) *ListQualityResultsByRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) SetStartDate(v string) *ListQualityResultsByRuleRequest {
	s.StartDate = &v
	return s
}

func (s *ListQualityResultsByRuleRequest) Validate() error {
	return dara.Validate(s)
}
