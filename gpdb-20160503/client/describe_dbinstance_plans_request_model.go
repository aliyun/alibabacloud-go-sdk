// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePlansRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBInstancePlansRequest
	GetOwnerId() *int64
	SetPlanCreateDate(v string) *DescribeDBInstancePlansRequest
	GetPlanCreateDate() *string
	SetPlanDesc(v string) *DescribeDBInstancePlansRequest
	GetPlanDesc() *string
	SetPlanId(v string) *DescribeDBInstancePlansRequest
	GetPlanId() *string
	SetPlanScheduleType(v string) *DescribeDBInstancePlansRequest
	GetPlanScheduleType() *string
	SetPlanType(v string) *DescribeDBInstancePlansRequest
	GetPlanType() *string
}

type DescribeDBInstancePlansRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time that is used to filter plans. If you specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format, the plans that are created before this time are returned. The time must be in UTC. If you do not specify this parameter, all plans are returned.
	//
	// example:
	//
	// 2022-04-17T23:00Z
	PlanCreateDate *string `json:"PlanCreateDate,omitempty" xml:"PlanCreateDate,omitempty"`
	// The description of the plan.
	//
	// example:
	//
	// this is a test plan
	PlanDesc *string `json:"PlanDesc,omitempty" xml:"PlanDesc,omitempty"`
	// The plan ID.
	//
	// > You can call the [DescribeDBInstancePlans](https://help.aliyun.com/document_detail/449398.html) operation to query the information about plans, including plan IDs.
	//
	// example:
	//
	// 1234
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The execution mode of the plan. Valid values:
	//
	// 	- **Postpone**: The plan is executed later.
	//
	// 	- **Regular**: The plan is executed periodically.
	//
	// example:
	//
	// Regular
	PlanScheduleType *string `json:"PlanScheduleType,omitempty" xml:"PlanScheduleType,omitempty"`
	// The type of the plan. Valid values:
	//
	// 	- **PauseResume**: pauses and resumes an instance.
	//
	// 	- **Resize**: scales an instance.
	//
	// 	- **ModifySpec**: changes compute node specifications.
	//
	// example:
	//
	// PauseResume
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
}

func (s DescribeDBInstancePlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePlansRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePlansRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancePlansRequest) GetPlanCreateDate() *string {
	return s.PlanCreateDate
}

func (s *DescribeDBInstancePlansRequest) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *DescribeDBInstancePlansRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeDBInstancePlansRequest) GetPlanScheduleType() *string {
	return s.PlanScheduleType
}

func (s *DescribeDBInstancePlansRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *DescribeDBInstancePlansRequest) SetDBInstanceId(v string) *DescribeDBInstancePlansRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetOwnerId(v int64) *DescribeDBInstancePlansRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanCreateDate(v string) *DescribeDBInstancePlansRequest {
	s.PlanCreateDate = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanDesc(v string) *DescribeDBInstancePlansRequest {
	s.PlanDesc = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanId(v string) *DescribeDBInstancePlansRequest {
	s.PlanId = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanScheduleType(v string) *DescribeDBInstancePlansRequest {
	s.PlanScheduleType = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) SetPlanType(v string) *DescribeDBInstancePlansRequest {
	s.PlanType = &v
	return s
}

func (s *DescribeDBInstancePlansRequest) Validate() error {
	return dara.Validate(s)
}
