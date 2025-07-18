// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDBInstancePlanStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SetDBInstancePlanStatusRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *SetDBInstancePlanStatusRequest
	GetOwnerId() *int64
	SetPlanId(v string) *SetDBInstancePlanStatusRequest
	GetPlanId() *string
	SetPlanStatus(v string) *SetDBInstancePlanStatusRequest
	GetPlanStatus() *string
}

type SetDBInstancePlanStatusRequest struct {
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
	// Specifies whether to enable or disable the plan. Valid values:
	//
	// 	- **disable**: disables the plan.
	//
	// 	- **enable**: enables the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	PlanStatus *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
}

func (s SetDBInstancePlanStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDBInstancePlanStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDBInstancePlanStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SetDBInstancePlanStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDBInstancePlanStatusRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *SetDBInstancePlanStatusRequest) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *SetDBInstancePlanStatusRequest) SetDBInstanceId(v string) *SetDBInstancePlanStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetOwnerId(v int64) *SetDBInstancePlanStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanId(v string) *SetDBInstancePlanStatusRequest {
	s.PlanId = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) SetPlanStatus(v string) *SetDBInstancePlanStatusRequest {
	s.PlanStatus = &v
	return s
}

func (s *SetDBInstancePlanStatusRequest) Validate() error {
	return dara.Validate(s)
}
