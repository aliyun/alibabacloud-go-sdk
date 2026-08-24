// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDiagnosticReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDiagnosticReportsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiagnosticReportsRequest
	GetRegionId() *string
	SetReportIds(v []*string) *DescribeDiagnosticReportsRequest
	GetReportIds() []*string
	SetResourceIds(v []*string) *DescribeDiagnosticReportsRequest
	GetResourceIds() []*string
	SetSeverity(v string) *DescribeDiagnosticReportsRequest
	GetSeverity() *string
	SetStatus(v string) *DescribeDiagnosticReportsRequest
	GetStatus() *string
}

type DescribeDiagnosticReportsRequest struct {
	// The maximum number of entries per page for paging. Maximum value: 100.
	//
	// Default value:
	//
	// - If this parameter is not set, the default value is 10.
	//
	// - If the value you set is greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the `NextToken` value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of diagnostic report IDs.
	ReportIds []*string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty" type:"Repeated"`
	// The list of resource IDs. A maximum of 100 IDs are supported.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The severity level. Valid values:
	//
	// - Unknown: The initial state, which indicates that the diagnosis has not started or the diagnosis process exited abnormally. No diagnostic conclusion is available.
	//
	// - Normal: The diagnosis is normal and no issues are found.
	//
	// - Info: Related information is available and may be associated with an exception.
	//
	// - Warn: Related information is available and may cause an exception.
	//
	// - Critical: A critical exception exists.
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The report status. Valid values:
	//
	// - InProgress: The diagnosis is in progress.
	//
	// - Failed: The diagnosis failed.
	//
	// - Finished: The diagnosis is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnosticReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiagnosticReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosticReportsRequest) GetReportIds() []*string {
	return s.ReportIds
}

func (s *DescribeDiagnosticReportsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeDiagnosticReportsRequest) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnosticReportsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosticReportsRequest) SetMaxResults(v int32) *DescribeDiagnosticReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetNextToken(v string) *DescribeDiagnosticReportsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetRegionId(v string) *DescribeDiagnosticReportsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetReportIds(v []*string) *DescribeDiagnosticReportsRequest {
	s.ReportIds = v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetResourceIds(v []*string) *DescribeDiagnosticReportsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetSeverity(v string) *DescribeDiagnosticReportsRequest {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) SetStatus(v string) *DescribeDiagnosticReportsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosticReportsRequest) Validate() error {
	return dara.Validate(s)
}
