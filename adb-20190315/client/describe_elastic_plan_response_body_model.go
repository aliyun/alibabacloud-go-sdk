// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticPlanList(v []*DescribeElasticPlanResponseBodyElasticPlanList) *DescribeElasticPlanResponseBody
	GetElasticPlanList() []*DescribeElasticPlanResponseBodyElasticPlanList
	SetRequestId(v string) *DescribeElasticPlanResponseBody
	GetRequestId() *string
}

type DescribeElasticPlanResponseBody struct {
	// The queried scaling plans.
	ElasticPlanList []*DescribeElasticPlanResponseBodyElasticPlanList `json:"ElasticPlanList,omitempty" xml:"ElasticPlanList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBody) GetElasticPlanList() []*DescribeElasticPlanResponseBodyElasticPlanList {
	return s.ElasticPlanList
}

func (s *DescribeElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlanResponseBody) SetElasticPlanList(v []*DescribeElasticPlanResponseBodyElasticPlanList) *DescribeElasticPlanResponseBody {
	s.ElasticPlanList = v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetRequestId(v string) *DescribeElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) Validate() error {
	if s.ElasticPlanList != nil {
		for _, item := range s.ElasticPlanList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticPlanResponseBodyElasticPlanList struct {
	// The number of nodes that are involved in the scaling plan.
	//
	// 	- If ElasticPlanType is set to **worker**, a value of 0 or null is returned.
	//
	// 	- If ElasticPlanType is set to **executorcombineworker*	- or **executor**, a value greater than 0 is returned.
	//
	// example:
	//
	// 0
	ElasticNodeNum *int32 `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- **worker**: scales only elastic I/O resources.
	//
	// 	- **executor**: scales only computing resources.
	//
	// 	- **executorcombineworker*	- (default): scales both elastic I/O resources and computing resources by proportion.
	//
	// example:
	//
	// worker
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Valid values:
	//
	// 	- 8 Core 64 GB (default)
	//
	// 	- 16 Core 64 GB
	//
	// 	- 32 Core 64 GB
	//
	// 	- 64 Core 128 GB
	//
	// 	- 12 Core 96 GB
	//
	// 	- 24 Core 96 GB
	//
	// 	- 52 Core 86 GB
	//
	// example:
	//
	// 16 Core 64 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	// Indicates whether the scaling plan takes effect. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end date of the scaling plan. This parameter is returned only if the end date of the scaling plan is set. The date is in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-09
	EndDay *string `json:"EndDay,omitempty" xml:"EndDay,omitempty"`
	// The restoration time of the scaling plan. The interval between the scale-up time and the restoration time cannot be more than 24 hours. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The days of the month when the scaling plan was executed. A value indicates a day of the month. Multiple values are separated by commas (,).
	//
	// example:
	//
	// 1,15,25
	MonthlyRepeat *string `json:"MonthlyRepeat,omitempty" xml:"MonthlyRepeat,omitempty"`
	// The name of the scaling plan.
	//
	// example:
	//
	// realtime
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// USER_DEFAULT
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The start date of the scaling plan. This parameter is returned only if the start date of the scaling plan is set. The date is in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-02
	StartDay *string `json:"StartDay,omitempty" xml:"StartDay,omitempty"`
	// The scale-up time of the scaling plan. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The days of the week when the scaling plan was executed. Valid values: 0 to 6, which indicate Sunday to Saturday. Multiple values are separated by commas (,).
	//
	// example:
	//
	// 3,4,5,6
	WeeklyRepeat *string `json:"WeeklyRepeat,omitempty" xml:"WeeklyRepeat,omitempty"`
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanResponseBodyElasticPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetElasticNodeNum() *int32 {
	return s.ElasticNodeNum
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetElasticPlanType() *string {
	return s.ElasticPlanType
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetElasticPlanWorkerSpec() *string {
	return s.ElasticPlanWorkerSpec
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetEndDay() *string {
	return s.EndDay
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetMonthlyRepeat() *string {
	return s.MonthlyRepeat
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetStartDay() *string {
	return s.StartDay
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) GetWeeklyRepeat() *string {
	return s.WeeklyRepeat
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticNodeNum(v int32) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticPlanType(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticPlanType = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetElasticPlanWorkerSpec(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEnable(v bool) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.Enable = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetEndTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetMonthlyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.MonthlyRepeat = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetPlanName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetResourcePoolName(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartDay(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartDay = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetStartTime(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) SetWeeklyRepeat(v string) *DescribeElasticPlanResponseBodyElasticPlanList {
	s.WeeklyRepeat = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyElasticPlanList) Validate() error {
	return dara.Validate(s)
}
