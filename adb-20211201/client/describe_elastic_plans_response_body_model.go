// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticPlans(v []*DescribeElasticPlansResponseBodyElasticPlans) *DescribeElasticPlansResponseBody
	GetElasticPlans() []*DescribeElasticPlansResponseBodyElasticPlans
	SetPageNumber(v int32) *DescribeElasticPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeElasticPlansResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeElasticPlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeElasticPlansResponseBody
	GetTotalCount() *int32
}

type DescribeElasticPlansResponseBody struct {
	// The queried scaling plans.
	ElasticPlans []*DescribeElasticPlansResponseBodyElasticPlans `json:"ElasticPlans,omitempty" xml:"ElasticPlans,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponseBody) GetElasticPlans() []*DescribeElasticPlansResponseBodyElasticPlans {
	return s.ElasticPlans
}

func (s *DescribeElasticPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeElasticPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeElasticPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeElasticPlansResponseBody) SetElasticPlans(v []*DescribeElasticPlansResponseBodyElasticPlans) *DescribeElasticPlansResponseBody {
	s.ElasticPlans = v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetPageNumber(v int32) *DescribeElasticPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetPageSize(v int32) *DescribeElasticPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetRequestId(v string) *DescribeElasticPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) SetTotalCount(v int32) *DescribeElasticPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticPlansResponseBody) Validate() error {
	if s.ElasticPlans != nil {
		for _, item := range s.ElasticPlans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticPlansResponseBodyElasticPlans struct {
	// Indicates whether **Proportional Default Scaling for EIUs*	- is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The name of the scaling plan.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Indicates whether the scaling plan is immediately enabled after the plan is created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when the next scheduling is performed.
	//
	// > The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	NextScheduleTime *string `json:"NextScheduleTime,omitempty" xml:"NextScheduleTime,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the name of a resource group within a cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The amount of elastic resources after scaling.
	//
	// example:
	//
	// 32ACU
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- **EXECUTOR**: the interactive resource group type, which specifies the computing resource type.
	//
	// 	- **WORKER**: the EIU type.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlansResponseBodyElasticPlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlansResponseBodyElasticPlans) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetNextScheduleTime() *string {
	return s.NextScheduleTime
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetTargetSize() *string {
	return s.TargetSize
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) GetType() *string {
	return s.Type
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetAutoScale(v bool) *DescribeElasticPlansResponseBodyElasticPlans {
	s.AutoScale = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetElasticPlanName(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetEnabled(v bool) *DescribeElasticPlansResponseBodyElasticPlans {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetNextScheduleTime(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.NextScheduleTime = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetResourceGroupName(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetTargetSize(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) SetType(v string) *DescribeElasticPlansResponseBodyElasticPlans {
	s.Type = &v
	return s
}

func (s *DescribeElasticPlansResponseBodyElasticPlans) Validate() error {
	return dara.Validate(s)
}
