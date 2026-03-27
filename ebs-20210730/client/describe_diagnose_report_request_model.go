// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDiagnoseReportRequest
	GetClientToken() *string
	SetDiagnoseType(v string) *DescribeDiagnoseReportRequest
	GetDiagnoseType() *string
	SetEndTime(v string) *DescribeDiagnoseReportRequest
	GetEndTime() *string
	SetMaxResults(v int32) *DescribeDiagnoseReportRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDiagnoseReportRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiagnoseReportRequest
	GetRegionId() *string
	SetReportIds(v []*string) *DescribeDiagnoseReportRequest
	GetReportIds() []*string
	SetResourceIds(v []*string) *DescribeDiagnoseReportRequest
	GetResourceIds() []*string
	SetStartTime(v string) *DescribeDiagnoseReportRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeDiagnoseReportRequest
	GetStatus() *string
}

type DescribeDiagnoseReportRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// Performance
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// example:
	//
	// 2024-08-26T02:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// adf8a2534f1c0a23e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// report-asb1s8***
	ReportIds []*string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty" type:"Repeated"`
	// example:
	//
	// d-asb1s8***
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-08-26T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDiagnoseReportRequest) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *DescribeDiagnoseReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnoseReportRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiagnoseReportRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnoseReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnoseReportRequest) GetReportIds() []*string {
	return s.ReportIds
}

func (s *DescribeDiagnoseReportRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeDiagnoseReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiagnoseReportRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnoseReportRequest) SetClientToken(v string) *DescribeDiagnoseReportRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetDiagnoseType(v string) *DescribeDiagnoseReportRequest {
	s.DiagnoseType = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetEndTime(v string) *DescribeDiagnoseReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetMaxResults(v int32) *DescribeDiagnoseReportRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetNextToken(v string) *DescribeDiagnoseReportRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetRegionId(v string) *DescribeDiagnoseReportRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetReportIds(v []*string) *DescribeDiagnoseReportRequest {
	s.ReportIds = v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetResourceIds(v []*string) *DescribeDiagnoseReportRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetStartTime(v string) *DescribeDiagnoseReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) SetStatus(v string) *DescribeDiagnoseReportRequest {
	s.Status = &v
	return s
}

func (s *DescribeDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
