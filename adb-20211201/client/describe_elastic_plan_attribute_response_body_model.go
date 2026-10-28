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
	// The elastic plan.
	ElasticPlan *DescribeElasticPlanAttributeResponseBodyElasticPlan `json:"ElasticPlan,omitempty" xml:"ElasticPlan,omitempty" type:"Struct"`
	// The ID of the request.
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
	if s.ElasticPlan != nil {
		if err := s.ElasticPlan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticPlanAttributeResponseBodyElasticPlan struct {
	// Indicates whether **proportional default elastic scaling for the storage acceleration service (EIU)*	- is enabled.
	//
	// Valid values:
	//
	// - true: Enabled. After this feature is enabled, storage resources scale in or out along with the elastic scaling of compute resources.
	//
	// - false: Disabled.
	//
	// > Only one storage elastic plan with proportional default elastic scaling for the storage acceleration service (EIU) can be enabled for the entire cluster. After this type of elastic plan is enabled, no other storage elastic plans can be enabled.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The cron expression, which indicates the scaling cycle and scaling time of the elastic plan.
	//
	// example:
	//
	// 0 20 14 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The name of the elastic plan.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the elastic plan is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The end time.
	//
	// > The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-01-01T12:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group to which the elastic plan belongs.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time.
	//
	// > The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The target specifications.
	//
	// example:
	//
	// 32ACU
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the elastic plan.
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
