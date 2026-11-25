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
	// The list of diagnostic reports.
	Reports []*DescribeDiagnoseReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// Id of the request
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

type DescribeDiagnoseReportsResponseBodyReports struct {
	// The time when the diagnostic report was created. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2024-08-23T02:22:30Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The details of the diagnostic report.
	Details []*DescribeDiagnoseReportsResponseBodyReportsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The diagnostic status. The value is the most severe status among the diagnostic details. Valid values:
	//
	// - Normal: the diagnostic result is normal.
	//
	// - Warn: the diagnostic result is a warning.
	//
	// - Critical: the diagnostic result is critical.
	//
	// example:
	//
	// Normal
	DiagnoseStatus *string `json:"DiagnoseStatus,omitempty" xml:"DiagnoseStatus,omitempty"`
	// The status of the diagnostic report. Valid values:
	//
	// - Processing: in progress.
	//
	// - Finished: completed.
	//
	// example:
	//
	// Finished
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The region ID.
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
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnoseReportsResponseBodyReportsDetails struct {
	// The type of the diagnostic item. Valid values:
	//
	// - AccountArrearage: account overdue payment. Checks whether the user account has overdue payments.
	//
	// - AccountNotEnoughBalance: insufficient account balance. Checks whether the balance of a user account on the China site is greater than or equal to CNY 100.
	//
	// - ElasticStrength: inventory health. Checks the inventory health of the instance type configured for the scaling group.
	//
	// - VSwitch: vSwitch status. Checks whether the vSwitch is active, such as whether it is deleted.
	//
	// - SecurityGroup: security group status. Checks whether the security group is active, such as whether it is deleted.
	//
	// - KeyPair: key pair status. Checks whether the key pair is active, such as whether it is deleted.
	//
	// - SlbBackendServerQuota: number of instances mounted to the SLB instance in the scaling group. Checks whether the number of ECS instances mounted to the default server group and virtual service groups of the SLB instance exceeds the quota.
	//
	// - AlbBackendServerQuota: number of instances mounted to the ALB instance in the scaling group. Checks whether the number of ECS instances mounted to the server groups of the ALB instance exceeds the quota.
	//
	// - NlbBackendServerQuota: number of instances mounted to the NLB instance in the scaling group. Checks whether the number of ECS instances mounted to the server groups of the NLB instance exceeds the quota.
	//
	// example:
	//
	// AccountArrearage
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The error code of the diagnostic item. Valid values:
	//
	// - VSwitchIdNotFound: the vSwitch does not exist.
	//
	// - SecurityGroupNotFound: the security group does not exist.
	//
	// - KeyPairNotFound: the key pair does not exist.
	//
	// - SlbBackendServerQuotaExceeded: the number of ECS instances mounted to the default server group and virtual service groups of the SLB instance exceeds the quota.
	//
	// - AlbBackendServerQuotaExceeded: the number of ECS instances mounted to the server groups of the ALB instance exceeds the quota.
	//
	// - NlbBackendServerQuotaExceeded: the number of ECS instances mounted to the server groups of the NLB instance exceeds the quota.
	//
	// - AccountArrearage: the account has overdue payments.
	//
	// - AccountNotEnoughBalance: the account balance is insufficient (less than CNY 100).
	//
	// - ElasticStrengthAlert: the inventory health is poor.
	//
	// example:
	//
	// VSwitchIdNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// sg-280ih****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The diagnostic status of the diagnostic item. Valid values:
	//
	// - Normal: the diagnostic result is normal.
	//
	// - Warn: the diagnostic result is a warning.
	//
	// - Critical: the diagnostic result is critical.
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
