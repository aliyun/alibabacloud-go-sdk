// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScale(v bool) *CreateElasticPlanRequest
	GetAutoScale() *bool
	SetCronExpression(v string) *CreateElasticPlanRequest
	GetCronExpression() *string
	SetDBClusterId(v string) *CreateElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *CreateElasticPlanRequest
	GetElasticPlanName() *string
	SetEnabled(v bool) *CreateElasticPlanRequest
	GetEnabled() *bool
	SetEndTime(v string) *CreateElasticPlanRequest
	GetEndTime() *string
	SetResourceGroupName(v string) *CreateElasticPlanRequest
	GetResourceGroupName() *string
	SetStartTime(v string) *CreateElasticPlanRequest
	GetStartTime() *string
	SetTargetSize(v string) *CreateElasticPlanRequest
	GetTargetSize() *string
	SetType(v string) *CreateElasticPlanRequest
	GetType() *string
}

type CreateElasticPlanRequest struct {
	// Specifies whether to enable **Default Proportional Scaling for EIUs**. Valid values:
	//
	// 	- true. In this case, storage resources are scaled along with computing resources, and the TargetSize and CronExpression parameters are not supported.
	//
	// 	- false
	//
	// >
	//
	// 	- This parameter must be specified when Type is set to WORKER. This parameter is not required when Type is set to EXECUTOR.
	//
	// 	- You can enable Default Proportional Scaling for EIUs for only a single scaling plan of a cluster.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// A CORN expression that specifies the scaling cycle and time for the scaling plan.
	//
	// example:
	//
	// 0 20 14 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  The name must be 2 to 30 characters in length and can contain letters, digits, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Specifies whether to immediately enable the scaling plan after the plan is created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The end time of the scaling plan.
	//
	// >  Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-01-01T12:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	//
	// >
	//
	// 	- If you want to create a scaling plan that uses interactive resource groups, you must specify this parameter. If you want to create a scaling plan that uses elastic I/O units (EIUs), you do not need to specify this parameter.
	//
	// 	- You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name for a cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
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
	// The type of the scaling plan. Valid values:
	//
	// 	- EXECUTOR: the interactive resource group type, which indicates the computing resource type.
	//
	// 	- WORKER: the EIU type.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *CreateElasticPlanRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *CreateElasticPlanRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateElasticPlanRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateElasticPlanRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *CreateElasticPlanRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateElasticPlanRequest) GetTargetSize() *string {
	return s.TargetSize
}

func (s *CreateElasticPlanRequest) GetType() *string {
	return s.Type
}

func (s *CreateElasticPlanRequest) SetAutoScale(v bool) *CreateElasticPlanRequest {
	s.AutoScale = &v
	return s
}

func (s *CreateElasticPlanRequest) SetCronExpression(v string) *CreateElasticPlanRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetEnabled(v bool) *CreateElasticPlanRequest {
	s.Enabled = &v
	return s
}

func (s *CreateElasticPlanRequest) SetEndTime(v string) *CreateElasticPlanRequest {
	s.EndTime = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceGroupName(v string) *CreateElasticPlanRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetStartTime(v string) *CreateElasticPlanRequest {
	s.StartTime = &v
	return s
}

func (s *CreateElasticPlanRequest) SetTargetSize(v string) *CreateElasticPlanRequest {
	s.TargetSize = &v
	return s
}

func (s *CreateElasticPlanRequest) SetType(v string) *CreateElasticPlanRequest {
	s.Type = &v
	return s
}

func (s *CreateElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
