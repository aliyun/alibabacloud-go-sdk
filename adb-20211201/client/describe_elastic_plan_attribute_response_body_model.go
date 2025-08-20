// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticPlan(v *DescribeElasticPlanAttributeResponseBodyElasticPlan) *DescribeElasticPlanAttributeResponseBody
	GetElasticPlan() *DescribeElasticPlanAttributeResponseBodyElasticPlan
	SetRequestId(v string) *DescribeElasticPlanAttributeResponseBody
	GetRequestId() *string
}

type DescribeElasticPlanAttributeResponseBody struct {
	// The queried scaling plan.
	ElasticPlan *DescribeElasticPlanAttributeResponseBodyElasticPlan `json:"ElasticPlan,omitempty" xml:"ElasticPlan,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticPlanAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponseBody) GetElasticPlan() *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	return s.ElasticPlan
}

func (s *DescribeElasticPlanAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlanAttributeResponseBody) SetElasticPlan(v *DescribeElasticPlanAttributeResponseBodyElasticPlan) *DescribeElasticPlanAttributeResponseBody {
	s.ElasticPlan = v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBody) SetRequestId(v string) *DescribeElasticPlanAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticPlanAttributeResponseBodyElasticPlan struct {
	// Indicates whether **Default Proportional Scaling for EIUs*	- is enabled. Valid values: true: Default Proportional Scaling for EIUs is enabled. If you set this parameter to true, storage resources are scaled along with computing resources. false: Default Proportional Scaling for EIUs is not enabled.
	//
	// >  You can enable Default Proportional Scaling for EIUs for only a single scaling plan of a cluster. After you enable a scaling plan of the Default Proportional Scaling for EIUs type, you cannot enable scaling plans of other types.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// A CORN expression that indicates the scaling cycle and time for the scaling plan.
	//
	// example:
	//
	// 0 20 14 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The name of the scaling plan.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the scaling plan is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The end time of the scaling plan.
	//
	// >  The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-01-01T12:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group used by the scaling plan.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time of the scaling plan.
	//
	// >  The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The amount of elastic resources after scaling.
	//
	// example:
	//
	// 32ACU
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the scaling plan.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanAttributeResponseBodyElasticPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanAttributeResponseBodyElasticPlan) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetTargetSize() *string {
	return s.TargetSize
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) GetType() *string {
	return s.Type
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetAutoScale(v bool) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.AutoScale = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetCronExpression(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.CronExpression = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetElasticPlanName(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetEnabled(v bool) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetEndTime(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetResourceGroupName(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetStartTime(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetTargetSize(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) SetType(v string) *DescribeElasticPlanAttributeResponseBodyElasticPlan {
	s.Type = &v
	return s
}

func (s *DescribeElasticPlanAttributeResponseBodyElasticPlan) Validate() error {
	return dara.Validate(s)
}
