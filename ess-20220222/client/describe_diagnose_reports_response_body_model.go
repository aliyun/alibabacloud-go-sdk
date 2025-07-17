// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDiagnoseReportsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiagnoseReportsResponseBody
	GetPageSize() *int32
	SetReports(v []*DescribeDiagnoseReportsResponseBodyReports) *DescribeDiagnoseReportsResponseBody
	GetReports() []*DescribeDiagnoseReportsResponseBodyReports
	SetRequestId(v string) *DescribeDiagnoseReportsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDiagnoseReportsResponseBody
	GetTotalCount() *int32
}

type DescribeDiagnoseReportsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The diagnostic reports.
	Reports []*DescribeDiagnoseReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ECA123C6-107B-5F70-A177-740A7224C996
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of diagnostic reports.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnoseReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnoseReportsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiagnoseReportsResponseBody) GetReports() []*DescribeDiagnoseReportsResponseBodyReports {
	return s.Reports
}

func (s *DescribeDiagnoseReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnoseReportsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiagnoseReportsResponseBody) SetPageNumber(v int32) *DescribeDiagnoseReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBody) SetPageSize(v int32) *DescribeDiagnoseReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBody) SetReports(v []*DescribeDiagnoseReportsResponseBodyReports) *DescribeDiagnoseReportsResponseBody {
	s.Reports = v
	return s
}

func (s *DescribeDiagnoseReportsResponseBody) SetRequestId(v string) *DescribeDiagnoseReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBody) SetTotalCount(v int32) *DescribeDiagnoseReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnoseReportsResponseBodyReports struct {
	// The time when the diagnostic report was created.
	//
	// example:
	//
	// 2024-08-23T02:22:30Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The details of the diagnostic report.
	Details []*DescribeDiagnoseReportsResponseBodyReportsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The status of the diagnostic item. Only the severe status is displayed in the diagnostic report. Valid values:
	//
	// 	- Normal: The diagnostic result is normal.
	//
	// 	- Warn: The diagnostic result is warning.
	//
	// 	- Critical: The diagnostic result is critical.
	//
	// example:
	//
	// Normal
	DiagnoseStatus *string `json:"DiagnoseStatus,omitempty" xml:"DiagnoseStatus,omitempty"`
	// The status of the diagnostic report. Valid values:
	//
	// 	- processing: The diagnosis is in progress.
	//
	// 	- Finished: The diagnosis is complete.
	//
	// example:
	//
	// Finished
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-bp14p0cjp7wvjob5l6hk
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp124uve5iph3*****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The user ID of the scaling group.
	//
	// example:
	//
	// 161456884*******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeDiagnoseReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetDetails() []*DescribeDiagnoseReportsResponseBodyReportsDetails {
	return s.Details
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetDiagnoseStatus() *string {
	return s.DiagnoseStatus
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeDiagnoseReportsResponseBodyReports) GetUserId() *string {
	return s.UserId
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetCreationTime(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.CreationTime = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetDetails(v []*DescribeDiagnoseReportsResponseBodyReportsDetails) *DescribeDiagnoseReportsResponseBodyReports {
	s.Details = v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetDiagnoseStatus(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.DiagnoseStatus = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetProcessStatus(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetRegionId(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetReportId(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetScalingGroupId(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) SetUserId(v string) *DescribeDiagnoseReportsResponseBodyReports {
	s.UserId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}

type DescribeDiagnoseReportsResponseBodyReportsDetails struct {
	// The type of the diagnostic item. Valid values:
	//
	// 	- AccountArrearage: Checks whether your Alibaba Cloud account has overdue payments.
	//
	// 	- AccountNotEnoughBalance: Checks whether the balance of your Alibaba Cloud account at the China site (aliyun.com) is greater than or equal to CNY 100.
	//
	// 	- ElasticStrength: Checks whether the instance types that are specified in the scaling configuration are sufficient.
	//
	// 	- VSwitch: Checks whether a specific vSwitch can work as expected. For example, if a vSwitch is deleted, the vSwitch cannot provide services and an exception occurs.
	//
	// 	- SecurityGroup: Checks whether a specific security group can work as expected. For example, if a security group is deleted, the security group cannot provide services and an exception occurs.
	//
	// 	- KeyPair: Checks whether the key pair is available. If the specified key pair is deleted, specify another key pair for the scaling group.
	//
	// 	- SlbBackendServerQuota: Checks whether the number of ECS instances that are added to the default server group and the vServer groups of the SLB instances associated with the scaling group has reached the upper limit.
	//
	// 	- AlbBackendServerQuota: Checks whether the number of ECS instances that are added to the backend server groups of the ALB instances associated with the scaling group has reached the upper limit.
	//
	// 	- NlbBackendServerQuota: Checks whether the number of ECS instances that are added to the backend server groups of the NLB instances associated with the scaling group has reached the upper limit.
	//
	// example:
	//
	// AccountArrearage
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The error code of the diagnostic item. Valid values:
	//
	// 	- VSwitchIdNotFound: The vSwitch does not exist.
	//
	// 	- SecurityGroupNotFound: The security group does not exist.
	//
	// 	- KeyPairNotFound: The key pair does not exist.
	//
	// 	- SlbBackendServerQuotaExceeded: The number of ECS instances that are added to the default server group and the vServer groups of the SLB instances associated with the scaling group has reached the upper limit.
	//
	// 	- AlbBackendServerQuotaExceeded: The number of ECS instances that are attached to the ALB instances of the scaling group has reached the upper limit.
	//
	// 	- NlbBackendServerQuotaExceeded: The number of ECS instances that are attached to the NLB instances of the scaling group has reached the upper limit.
	//
	// 	- AccountArrearage: Your account has overdue payments.
	//
	// 	- AccountNotEnoughBalance: The balance of your Alibaba Cloud account is less than CNY 100.
	//
	// 	- ElasticStrengthAlert: The inventory levels are lower than expected.
	//
	// example:
	//
	// VSwitchIdNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// sg-280ih****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the diagnostic item. Valid values:
	//
	// 	- Normal: The diagnostic result is normal.
	//
	// 	- Warn: The diagnostic result is warning.
	//
	// 	- Critical: The diagnostic result is critical.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiagnoseReportsResponseBodyReportsDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportsResponseBodyReportsDetails) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) SetDiagnoseType(v string) *DescribeDiagnoseReportsResponseBodyReportsDetails {
	s.DiagnoseType = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) SetErrorCode(v string) *DescribeDiagnoseReportsResponseBodyReportsDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) SetResourceId(v string) *DescribeDiagnoseReportsResponseBodyReportsDetails {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) SetStatus(v string) *DescribeDiagnoseReportsResponseBodyReportsDetails {
	s.Status = &v
	return s
}

func (s *DescribeDiagnoseReportsResponseBodyReportsDetails) Validate() error {
	return dara.Validate(s)
}
