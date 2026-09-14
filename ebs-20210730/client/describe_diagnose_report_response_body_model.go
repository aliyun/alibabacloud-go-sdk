// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDiagnoseReportResponseBody
	GetNextToken() *string
	SetReports(v []*DescribeDiagnoseReportResponseBodyReports) *DescribeDiagnoseReportResponseBody
	GetReports() []*DescribeDiagnoseReportResponseBodyReports
	SetRequestId(v string) *DescribeDiagnoseReportResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDiagnoseReportResponseBody
	GetTotalCount() *int32
}

type DescribeDiagnoseReportResponseBody struct {
	// The pagination token returned in this call.
	//
	// example:
	//
	// f07b150eadfa1d7a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of diagnostic reports.
	Reports []*DescribeDiagnoseReportResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AAA478A0-BEE6-1D42-BEB6-A9CFEAD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnoseReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnoseReportResponseBody) GetReports() []*DescribeDiagnoseReportResponseBodyReports {
	return s.Reports
}

func (s *DescribeDiagnoseReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnoseReportResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiagnoseReportResponseBody) SetNextToken(v string) *DescribeDiagnoseReportResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) SetReports(v []*DescribeDiagnoseReportResponseBodyReports) *DescribeDiagnoseReportResponseBody {
	s.Reports = v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) SetRequestId(v string) *DescribeDiagnoseReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) SetTotalCount(v int32) *DescribeDiagnoseReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBody) Validate() error {
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnoseReportResponseBodyReports struct {
	// The user ID.
	//
	// example:
	//
	// 196380451****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the diagnostic report was created, in Unix/POSIX timestamp (seconds).
	//
	// example:
	//
	// 1727239294
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The end timestamp of the resource diagnosis.
	//
	// example:
	//
	// 1727239294
	DiagnoseEndTime *int64 `json:"DiagnoseEndTime,omitempty" xml:"DiagnoseEndTime,omitempty"`
	// The start timestamp of the resource diagnosis.
	//
	// example:
	//
	// 1727229294
	DiagnoseStartTime *int64 `json:"DiagnoseStartTime,omitempty" xml:"DiagnoseStartTime,omitempty"`
	// The type of diagnosis.
	//
	// example:
	//
	// Performance
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The list of diagnosed issues.
	Events []*DescribeDiagnoseReportResponseBodyReportsEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The time when the diagnostic report was completed, in Unix/POSIX timestamp (seconds).
	//
	// example:
	//
	// 1727239295
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The diagnostic report ID.
	//
	// example:
	//
	// report-sag8d****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// d-wz95ycu****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - Disk
	//
	// example:
	//
	// Disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The severity level of the diagnosis. The severity levels in ascending order are:
	//
	// - Info: Associated information that may be related to an anomaly.
	//
	// - Warn: Associated information that may cause an anomaly.
	//
	// - Critical: A critical anomaly exists.
	//
	// example:
	//
	// Warn
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The status of the diagnostic report. Valid values:
	//
	// - Running
	//
	// - Success
	//
	// - TimeOut
	//
	// - Fail
	//
	// The Severity and Events fields are valid only when Status is set to Success.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyReports) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetDiagnoseEndTime() *int64 {
	return s.DiagnoseEndTime
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetDiagnoseStartTime() *int64 {
	return s.DiagnoseStartTime
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetEvents() []*DescribeDiagnoseReportResponseBodyReportsEvents {
	return s.Events
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnoseReportResponseBodyReports) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetAliUid(v int64) *DescribeDiagnoseReportResponseBodyReports {
	s.AliUid = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetCreationTime(v int64) *DescribeDiagnoseReportResponseBodyReports {
	s.CreationTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetDiagnoseEndTime(v int64) *DescribeDiagnoseReportResponseBodyReports {
	s.DiagnoseEndTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetDiagnoseStartTime(v int64) *DescribeDiagnoseReportResponseBodyReports {
	s.DiagnoseStartTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetDiagnoseType(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.DiagnoseType = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetEvents(v []*DescribeDiagnoseReportResponseBodyReportsEvents) *DescribeDiagnoseReportResponseBodyReports {
	s.Events = v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetFinishedTime(v int64) *DescribeDiagnoseReportResponseBodyReports {
	s.FinishedTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetRegionId(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetReportId(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetResourceId(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetResourceType(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetSeverity(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) SetStatus(v string) *DescribeDiagnoseReportResponseBodyReports {
	s.Status = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReports) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnoseReportResponseBodyReportsEvents struct {
	// The event description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The issue name. Valid values:
	//
	// - NoSnapshot: data protection
	//
	// - BurstIOTriggered: I/O burst
	//
	// - CostOptimizationNeeded: cost optimization
	//
	// - DiskSpecNotMatchedWithInstance: instance and cloud disk specification mismatch
	//
	// - DiskIONo4kAligned: non-4K-aligned read/write
	//
	// - DiskIOHang: IOHang occurred on the cloud disk
	//
	// - InstanceIOPSExceedInstanceMaxLimit: instance IOPS reached the upper limit
	//
	// - InstanceBPSExceedInstanceMaxLimit: instance BPS reached the upper limit
	//
	// - DiskIOPSExceedInstanceMaxLimit: cloud disk IOPS reached the instance upper limit
	//
	// - DiskBPSExceedInstanceMaxLimit: cloud disk BPS reached the instance upper limit
	//
	// - DiskIOPSExceedDiskMaxLimit: cloud disk IOPS reached the cloud disk upper limit
	//
	// - DiskBPSExceedDiskMaxLimit: cloud disk BPS reached the cloud disk upper limit
	//
	// example:
	//
	// DiskIOPSExceedDiskMaxLimit
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The recommended action after the event occurs. Valid values:
	//
	// - ModifyDiskSpec: change cloud disk specifications
	//
	// - CreateSnapshot: create a snapshot
	//
	// - ResizeDisk: expand the cloud disk
	//
	// - AdjustProvision: adjust provisioned performance
	//
	// - ModifyInstanceSpec: change instance specifications
	//
	// example:
	//
	// ResizeDisk
	RecommendAction *string `json:"RecommendAction,omitempty" xml:"RecommendAction,omitempty"`
	// The parameters for the recommended action after the event occurs.
	//
	// example:
	//
	// 4096
	RecommendParams *string `json:"RecommendParams,omitempty" xml:"RecommendParams,omitempty"`
	// The severity level of the diagnosed issue. The severity levels in ascending order are:
	//
	// - Info: Associated information that may be related to an anomaly.
	//
	// - Warn: Associated information that may cause an anomaly.
	//
	// - Critical: A critical anomaly exists.
	//
	// example:
	//
	// Warn
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The start timestamp of the event, in milliseconds.
	//
	// example:
	//
	// 1755756214000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiagnoseReportResponseBodyReportsEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponseBodyReportsEvents) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetEventName() *string {
	return s.EventName
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetRecommendAction() *string {
	return s.RecommendAction
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetRecommendParams() *string {
	return s.RecommendParams
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetDescription(v string) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.Description = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetEventName(v string) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.EventName = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetRecommendAction(v string) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.RecommendAction = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetRecommendParams(v string) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.RecommendParams = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetSeverity(v string) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.Severity = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) SetStartTime(v int64) *DescribeDiagnoseReportResponseBodyReportsEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnoseReportResponseBodyReportsEvents) Validate() error {
	return dara.Validate(s)
}
