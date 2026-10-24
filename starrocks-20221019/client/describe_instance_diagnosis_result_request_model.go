// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDiagnosisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v string) *DescribeInstanceDiagnosisResultRequest
	GetDimension() *string
	SetInstanceId(v string) *DescribeInstanceDiagnosisResultRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeInstanceDiagnosisResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceDiagnosisResultRequest
	GetPageSize() *int32
	SetReportDate(v string) *DescribeInstanceDiagnosisResultRequest
	GetReportDate() *string
	SetStatuses(v string) *DescribeInstanceDiagnosisResultRequest
	GetStatuses() *string
}

type DescribeInstanceDiagnosisResultRequest struct {
	// The evaluation dimension.
	//
	// example:
	//
	// table_analysis
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for paging. Default value: 1.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of authorization entries returned per page when paging. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The report date in the YYYY-MM-DD format.
	//
	// example:
	//
	// 2026-03-08
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// The status collection.
	//
	// example:
	//
	// healthy
	Statuses *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s DescribeInstanceDiagnosisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDiagnosisResultRequest) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeInstanceDiagnosisResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDiagnosisResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceDiagnosisResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceDiagnosisResultRequest) GetReportDate() *string {
	return s.ReportDate
}

func (s *DescribeInstanceDiagnosisResultRequest) GetStatuses() *string {
	return s.Statuses
}

func (s *DescribeInstanceDiagnosisResultRequest) SetDimension(v string) *DescribeInstanceDiagnosisResultRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) SetInstanceId(v string) *DescribeInstanceDiagnosisResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) SetPageNumber(v int32) *DescribeInstanceDiagnosisResultRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) SetPageSize(v int32) *DescribeInstanceDiagnosisResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) SetReportDate(v string) *DescribeInstanceDiagnosisResultRequest {
	s.ReportDate = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) SetStatuses(v string) *DescribeInstanceDiagnosisResultRequest {
	s.Statuses = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultRequest) Validate() error {
	return dara.Validate(s)
}
