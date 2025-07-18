// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstancePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpdateDBInstancePlanRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UpdateDBInstancePlanRequest
	GetOwnerId() *int64
	SetPlanConfig(v string) *UpdateDBInstancePlanRequest
	GetPlanConfig() *string
	SetPlanDesc(v string) *UpdateDBInstancePlanRequest
	GetPlanDesc() *string
	SetPlanEndDate(v string) *UpdateDBInstancePlanRequest
	GetPlanEndDate() *string
	SetPlanId(v string) *UpdateDBInstancePlanRequest
	GetPlanId() *string
	SetPlanName(v string) *UpdateDBInstancePlanRequest
	GetPlanName() *string
	SetPlanStartDate(v string) *UpdateDBInstancePlanRequest
	GetPlanStartDate() *string
}

type UpdateDBInstancePlanRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
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
	// example:
	//
	// {"pause":{"executeTime":"2022-08-30T16:00:00Z"}}
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
	// 	- This parameter must be specified only for **periodically executed*	- plans.
	//
	// 	- If you do not specify this parameter, the plan stops until the plan is deleted.
	//
	// example:
	//
	// 2023-04-17T23:00Z
	PlanEndDate *string `json:"PlanEndDate,omitempty" xml:"PlanEndDate,omitempty"`
	// The ID of the plan.
	//
	// >  You can call the [DescribeDBInstancePlans](https://help.aliyun.com/document_detail/449398.html) operation to query the details of plans, including plan IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	//
	// example:
	//
	// test-plan
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The start time of the plan. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >
	//
	// 	- This parameter must be specified only for **periodically executed*	- plans.
	//
	// 	- If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 2022-04-17T23:00Z
	PlanStartDate *string `json:"PlanStartDate,omitempty" xml:"PlanStartDate,omitempty"`
}

func (s UpdateDBInstancePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstancePlanRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdateDBInstancePlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateDBInstancePlanRequest) GetPlanConfig() *string {
	return s.PlanConfig
}

func (s *UpdateDBInstancePlanRequest) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *UpdateDBInstancePlanRequest) GetPlanEndDate() *string {
	return s.PlanEndDate
}

func (s *UpdateDBInstancePlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateDBInstancePlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *UpdateDBInstancePlanRequest) GetPlanStartDate() *string {
	return s.PlanStartDate
}

func (s *UpdateDBInstancePlanRequest) SetDBInstanceId(v string) *UpdateDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetOwnerId(v int64) *UpdateDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanConfig(v string) *UpdateDBInstancePlanRequest {
	s.PlanConfig = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanDesc(v string) *UpdateDBInstancePlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanEndDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanEndDate = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanId(v string) *UpdateDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanName(v string) *UpdateDBInstancePlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) SetPlanStartDate(v string) *UpdateDBInstancePlanRequest {
	s.PlanStartDate = &v
	return s
}

func (s *UpdateDBInstancePlanRequest) Validate() error {
	return dara.Validate(s)
}
