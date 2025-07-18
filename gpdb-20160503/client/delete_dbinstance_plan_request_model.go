// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstancePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstancePlanRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DeleteDBInstancePlanRequest
	GetOwnerId() *int64
	SetPlanId(v string) *DeleteDBInstancePlanRequest
	GetPlanId() *string
}

type DeleteDBInstancePlanRequest struct {
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
}

func (s DeleteDBInstancePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstancePlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstancePlanRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstancePlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBInstancePlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteDBInstancePlanRequest) SetDBInstanceId(v string) *DeleteDBInstancePlanRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetOwnerId(v int64) *DeleteDBInstancePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) SetPlanId(v string) *DeleteDBInstancePlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteDBInstancePlanRequest) Validate() error {
	return dara.Validate(s)
}
