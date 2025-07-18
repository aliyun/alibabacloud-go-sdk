// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstancePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstancePlanRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *CreateDBInstancePlanRequest
	GetOwnerId() *int64
	SetPlanConfig(v string) *CreateDBInstancePlanRequest
	GetPlanConfig() *string
	SetPlanDesc(v string) *CreateDBInstancePlanRequest
	GetPlanDesc() *string
	SetPlanEndDate(v string) *CreateDBInstancePlanRequest
	GetPlanEndDate() *string
	SetPlanName(v string) *CreateDBInstancePlanRequest
	GetPlanName() *string
	SetPlanScheduleType(v string) *CreateDBInstancePlanRequest
	GetPlanScheduleType() *string
	SetPlanStartDate(v string) *CreateDBInstancePlanRequest
	GetPlanStartDate() *string
	SetPlanType(v string) *CreateDBInstancePlanRequest
	GetPlanType() *string
}

type CreateDBInstancePlanRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The execution information of the plan. Specify the parameter in the JSON format. The parameter value varies based on the values of **PlanType*	- and **PlanScheduleType**. The following section describes the PlanConfig parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"pause":{"planCronTime":"0 0 12 1/1 	- ? "},"resume":{"planCronTime":"0 0 0 1/1 	- ? "}}
	PlanConfig *string `json:"PlanConfig,omitempty" xml:"PlanConfig,omitempty"`
	// The description of the plan.
	//
	// example:
	//
	// this is a test plan
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The end time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// >
	//
	// 	- This parameter must be specified only when **PlanScheduleType*	- is set to **Regular**.
	//
	// 	- If you do not specify this parameter, the plan stops until the plan is deleted.
	//
	// example:
	//
	// 2023-04-17T23:00Z
	PlanEndDate *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	// The name of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-plan
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The execution mode of the plan. Valid values:
	//
	// 	- **Postpone**: The plan is executed later.
	//
	// 	- **Regular**: The plan is executed periodically.
	//
	// This parameter is required.
	//
	// example:
	//
	// Regular
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	// The start time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >
	//
	// 	- This parameter must be specified only when **PlanScheduleType*	- is set to **Regular**.
	//
	// 	- If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 2022-04-17T23:00Z
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
	// The type of the plan. Valid values:
	//
	// 	- **PauseResume**: pauses and resumes an instance.
	//
	// 	- **Resize**: changes the number of compute nodes.
	//
	// 	- **ModifySpec**: changes compute node specifications.
	//
	// > - You can specify the value to ModifySpec only for instances in elastic storage mode.
	//
	// >- You can specify the value to ModifySpec only for instances in elastic storage mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// PauseResume
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s CreateDBInstancePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstancePlanRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstancePlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstancePlanRequest) GetPlanConfig() *string {
	return s.PlanConfig
}

func (s *CreateDBInstancePlanRequest) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *CreateDBInstancePlanRequest) GetPlanEndDate() *string {
	return s.PlanEndDate
}

func (s *CreateDBInstancePlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateDBInstancePlanRequest) GetPlanScheduleType() *string {
	return s.PlanScheduleType
}

func (s *CreateDBInstancePlanRequest) GetPlanStartDate() *string {
	return s.PlanStartDate
}

func (s *CreateDBInstancePlanRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *CreateDBInstancePlanRequest) SetDBInstanceId(v string) *CreateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetOwnerId(v int64) *CreateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanConfig(v string) *CreateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanDesc(v string) *CreateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanEndDate(v string) *CreateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanName(v string) *CreateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanScheduleType(v string) *CreateDBInstancePlanRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanStartDate(v string) *CreateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

func (s *CreateDBInstancePlanRequest) SetPlanType(v string) *CreateDBInstancePlanRequest {
	s.PlanType = &v
	return s
}

func (s *CreateDBInstancePlanRequest) Validate() error {
	return dara.Validate(s)
}
