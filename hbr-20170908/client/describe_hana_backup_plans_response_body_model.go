// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaBackupPlansResponseBody
	GetCode() *string
	SetHanaBackupPlans(v *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) *DescribeHanaBackupPlansResponseBody
	GetHanaBackupPlans() *DescribeHanaBackupPlansResponseBodyHanaBackupPlans
	SetMessage(v string) *DescribeHanaBackupPlansResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHanaBackupPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaBackupPlansResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHanaBackupPlansResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaBackupPlansResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeHanaBackupPlansResponseBody
	GetTotalCount() *int64
}

type DescribeHanaBackupPlansResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the backup plan.
	HanaBackupPlans *DescribeHanaBackupPlansResponseBodyHanaBackupPlans `json:"HanaBackupPlans,omitempty" xml:"HanaBackupPlans,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F029C1C7-26B6-5ADD-A73E-D85CCD7C73A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaBackupPlansResponseBody) GetHanaBackupPlans() *DescribeHanaBackupPlansResponseBodyHanaBackupPlans {
	return s.HanaBackupPlans
}

func (s *DescribeHanaBackupPlansResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaBackupPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaBackupPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaBackupPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaBackupPlansResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaBackupPlansResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHanaBackupPlansResponseBody) SetCode(v string) *DescribeHanaBackupPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetHanaBackupPlans(v *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) *DescribeHanaBackupPlansResponseBody {
	s.HanaBackupPlans = v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetMessage(v string) *DescribeHanaBackupPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageNumber(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageSize(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetRequestId(v string) *DescribeHanaBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetSuccess(v bool) *DescribeHanaBackupPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetTotalCount(v int64) *DescribeHanaBackupPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) Validate() error {
	if s.HanaBackupPlans != nil {
		if err := s.HanaBackupPlans.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlans struct {
	HanaBackupPlan []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan `json:"HanaBackupPlan,omitempty" xml:"HanaBackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) GetHanaBackupPlan() []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	return s.HanaBackupPlan
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) SetHanaBackupPlan(v []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) *DescribeHanaBackupPlansResponseBodyHanaBackupPlans {
	s.HanaBackupPlan = v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) Validate() error {
	if s.HanaBackupPlan != nil {
		for _, item := range s.HanaBackupPlan {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan struct {
	// The backup prefix.
	//
	// example:
	//
	// COMPLETE_DATA_BACKUP
	BackupPrefix *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	// The backup type. Valid values:
	//
	// 	- COMPLETE: full backup
	//
	// 	- INCREMENTAL: incremental backup
	//
	// 	- DIFFERENTIAL: differential backup
	//
	// example:
	//
	// COMPLETE
	BackupType     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0002scknka*****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Indicates whether the backup plan is disabled. Valid values:
	//
	// 	- true: The backup plan is disabled.
	//
	// 	- false: The backup plan is enabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The ID of the backup plan.
	//
	// example:
	//
	// pl-0000tnyndg3ne5m4ubeu
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the backup plan.
	//
	// example:
	//
	// plan-20220118-141153
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000csihw82pqkd7hcjws
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupPrefix(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupPrefix = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupType(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupType = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBusinessStatus(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetClusterId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDatabaseName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDisabled(v bool) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Disabled = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetSchedule(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Schedule = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetVaultId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) Validate() error {
	return dara.Validate(s)
}
