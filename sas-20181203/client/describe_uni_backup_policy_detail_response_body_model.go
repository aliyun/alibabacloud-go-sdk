// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPolicyDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUniBackupPolicyDetailResponseBody
	GetRequestId() *string
	SetUniBackupPolicyDTO(v *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) *DescribeUniBackupPolicyDetailResponseBody
	GetUniBackupPolicyDTO() *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO
}

type DescribeUniBackupPolicyDetailResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the anti-ransomware policy.
	UniBackupPolicyDTO *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO `json:"UniBackupPolicyDTO,omitempty" xml:"UniBackupPolicyDTO,omitempty" type:"Struct"`
}

func (s DescribeUniBackupPolicyDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniBackupPolicyDetailResponseBody) GetUniBackupPolicyDTO() *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	return s.UniBackupPolicyDTO
}

func (s *DescribeUniBackupPolicyDetailResponseBody) SetRequestId(v string) *DescribeUniBackupPolicyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBody) SetUniBackupPolicyDTO(v *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) *DescribeUniBackupPolicyDetailResponseBody {
	s.UniBackupPolicyDTO = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBody) Validate() error {
	if s.UniBackupPolicyDTO != nil {
		if err := s.UniBackupPolicyDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO struct {
	// The name of the database account.
	//
	// example:
	//
	// admin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the database client. Valid values:
	//
	// 	- **UNKNOWN**: unknown
	//
	// 	- **INSTALLED**: installed
	//
	// 	- **INSTALL_FAILED**: installation failed
	//
	// 	- **UNINSTALL_FAILED**: uninstallation failed
	//
	// example:
	//
	// INSTALLED
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **MYSQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// ORACLE
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The details of the policy for full backup.
	FullPlan *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan `json:"FullPlan,omitempty" xml:"FullPlan,omitempty" type:"Struct"`
	// The policy for incremental data backup.
	IncPlan *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan `json:"IncPlan,omitempty" xml:"IncPlan,omitempty" type:"Struct"`
	// The ID of the server.
	//
	// example:
	//
	// i-2zefcy2id5d60m9t****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// auto_test_sql
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **initiating**: initializing
	//
	// 	- **opening**: enabled
	//
	// 	- **closing**: disabled
	//
	// 	- **deleting**: deleting
	//
	// example:
	//
	// opening
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The retention period of the backup snapshot.
	//
	// example:
	//
	// 7
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The maximum network bandwidth that is allowed during data backup. Unit: bytes.
	//
	// example:
	//
	// 5242880
	SpeedLimiter *int64 `json:"SpeedLimiter,omitempty" xml:"SpeedLimiter,omitempty"`
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetFullPlan() *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan {
	return s.FullPlan
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetIncPlan() *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan {
	return s.IncPlan
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetRetention() *int32 {
	return s.Retention
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) GetSpeedLimiter() *int64 {
	return s.SpeedLimiter
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetAccountName(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.AccountName = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetAgentStatus(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.AgentStatus = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetDatabaseType(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.DatabaseType = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetFullPlan(v *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.FullPlan = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetIncPlan(v *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.IncPlan = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetInstanceId(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.InstanceId = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetInstanceName(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.InstanceName = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetPolicyId(v int64) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.PolicyId = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetPolicyName(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.PolicyName = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetPolicyStatus(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetRetention(v int32) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.Retention = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) SetSpeedLimiter(v int64) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO {
	s.SpeedLimiter = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTO) Validate() error {
	if s.FullPlan != nil {
		if err := s.FullPlan.Validate(); err != nil {
			return err
		}
	}
	if s.IncPlan != nil {
		if err := s.IncPlan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan struct {
	// An array that consists of the days of a week on which the backup is performed.
	Days []*string `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The interval of backup tasks.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The unit of the interval. Valid values:
	//
	// 	- **hourly**: hour
	//
	// 	- **daily**: day
	//
	// 	- **weekly**: week
	//
	// example:
	//
	// daily
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The time when the full backup started. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 00:10:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) GetDays() []*string {
	return s.Days
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) SetDays(v []*string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan {
	s.Days = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) SetInterval(v int32) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan {
	s.Interval = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) SetPlanType(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan {
	s.PlanType = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) SetStartTime(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan {
	s.StartTime = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOFullPlan) Validate() error {
	return dara.Validate(s)
}

type DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan struct {
	// An array that consists of the days of a week on which the backup is performed.
	Days []*string `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The interval of backup tasks.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The unit of the interval. Valid values:
	//
	// 	- **hourly**: hour
	//
	// 	- **daily**: day
	//
	// 	- **weekly**: week
	//
	// example:
	//
	// daily
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The time when the incremental data backup starts. The time is in the hh:mm:ss format.
	//
	// example:
	//
	// 00:10:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) GetDays() []*string {
	return s.Days
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) SetDays(v []*string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan {
	s.Days = v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) SetInterval(v int32) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan {
	s.Interval = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) SetPlanType(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan {
	s.PlanType = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) SetStartTime(v string) *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan {
	s.StartTime = &v
	return s
}

func (s *DescribeUniBackupPolicyDetailResponseBodyUniBackupPolicyDTOIncPlan) Validate() error {
	return dara.Validate(s)
}
