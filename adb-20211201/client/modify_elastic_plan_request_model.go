// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyElasticPlanRequest
	GetCronExpression() *string
	SetDBClusterId(v string) *ModifyElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *ModifyElasticPlanRequest
	GetElasticPlanName() *string
	SetEndTime(v string) *ModifyElasticPlanRequest
	GetEndTime() *string
	SetStartTime(v string) *ModifyElasticPlanRequest
	GetStartTime() *string
	SetTargetSize(v string) *ModifyElasticPlanRequest
	GetTargetSize() *string
}

type ModifyElasticPlanRequest struct {
	// A CORN expression that specifies the scaling cycle and time for the scaling plan.
	//
	// example:
	//
	// 0 20 14 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](https://help.aliyun.com/document_detail/601334.html) operation to query the names of scaling plans.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The end time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-01-01T12:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The desired specifications of elastic resources after scaling.
	//
	// >
	//
	// 	- If the scaling plan uses **EIUs*	- and **Default Proportional Scaling for EIUs*	- is enabled, you do not need to specify this parameter. In other cases, you must specify this parameter.
	//
	// 	- You can call the [DescribeElasticPlanSpecifications](https://help.aliyun.com/document_detail/601278.html) operation to query the specifications that are supported for scaling plans.
	//
	// example:
	//
	// 32ACU
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *ModifyElasticPlanRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyElasticPlanRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyElasticPlanRequest) GetTargetSize() *string {
	return s.TargetSize
}

func (s *ModifyElasticPlanRequest) SetCronExpression(v string) *ModifyElasticPlanRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetEndTime(v string) *ModifyElasticPlanRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetStartTime(v string) *ModifyElasticPlanRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetTargetSize(v string) *ModifyElasticPlanRequest {
	s.TargetSize = &v
	return s
}

func (s *ModifyElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
