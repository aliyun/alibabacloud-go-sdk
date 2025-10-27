// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticDailyPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticDailyPlanList(v []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) *DescribeElasticDailyPlanResponseBody
	GetElasticDailyPlanList() []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList
	SetRequestId(v string) *DescribeElasticDailyPlanResponseBody
	GetRequestId() *string
}

type DescribeElasticDailyPlanResponseBody struct {
	// Details of the current-day scaling plans.
	ElasticDailyPlanList []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList `json:"ElasticDailyPlanList,omitempty" xml:"ElasticDailyPlanList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBody) GetElasticDailyPlanList() []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	return s.ElasticDailyPlanList
}

func (s *DescribeElasticDailyPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticDailyPlanResponseBody) SetElasticDailyPlanList(v []*DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) *DescribeElasticDailyPlanResponseBody {
	s.ElasticDailyPlanList = v
	return s
}

func (s *DescribeElasticDailyPlanResponseBody) SetRequestId(v string) *DescribeElasticDailyPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBody) Validate() error {
	if s.ElasticDailyPlanList != nil {
		for _, item := range s.ElasticDailyPlanList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticDailyPlanResponseBodyElasticDailyPlanList struct {
	// The start date of the current-day scaling plan. The date is in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-02
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// The number of nodes involved in the scaling plan.
	//
	// 	- If ElasticPlanType is set to **worker**, a value of 0 or null is returned.
	//
	// 	- If ElasticPlanType is set to **executorcombineworker*	- or **executor**, a value greater than 0 is returned.
	//
	// example:
	//
	// 0
	ElasticNodeNum *int32 `json:"ElasticNodeNum,omitempty" xml:"ElasticNodeNum,omitempty"`
	// The type of the scaling plan. Default value: executorcombineworker. Valid values:
	//
	// 	- **worker**: scales only elastic I/O resources.
	//
	// 	- **executor**: scales only computing resources.
	//
	// 	- **executorcombineworker**: scales both elastic I/O resources and computing resources by proportion.
	//
	// example:
	//
	// worker
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Default value: 8 Core 64 GB. Valid values:
	//
	// 	- 8 Core 64 GB
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
	// The actual restoration time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-02 16:00:00
	EndTs *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The scheduled restoration time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-02 16:00:00
	PlanEndTs *string `json:"PlanEndTs,omitempty" xml:"PlanEndTs,omitempty"`
	// The name of the scaling plan.
	//
	// example:
	//
	// realtimeplan
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The scheduled scale-up time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-02 15:00:00
	PlanStartTs *string `json:"PlanStartTs,omitempty" xml:"PlanStartTs,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
	// The actual scale-up time. The time is in the yyyy-MM-dd hh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-02 16:00:00
	StartTs *string `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The execution state of the current-day scaling plan. Multiple values are separated by commas (,). Valid values:
	//
	// 	- **1**: The scaling plan is not executed.
	//
	// 	- **2**: The scaling plan is being executed.
	//
	// 	- **3**: The scaling plan is executed.
	//
	// 	- **4**: The scaling plan fails to be executed.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetDay() *string {
	return s.Day
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetElasticNodeNum() *int32 {
	return s.ElasticNodeNum
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetElasticPlanType() *string {
	return s.ElasticPlanType
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetElasticPlanWorkerSpec() *string {
	return s.ElasticPlanWorkerSpec
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetEndTs() *string {
	return s.EndTs
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetPlanEndTs() *string {
	return s.PlanEndTs
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetPlanStartTs() *string {
	return s.PlanStartTs
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetStartTs() *string {
	return s.StartTs
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetDay(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Day = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticNodeNum(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticNodeNum = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticPlanType(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticPlanType = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetElasticPlanWorkerSpec(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.EndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanEndTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanEndTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetPlanStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.PlanStartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetResourcePoolName(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStartTs(v string) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.StartTs = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) SetStatus(v int32) *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList {
	s.Status = &v
	return s
}

func (s *DescribeElasticDailyPlanResponseBodyElasticDailyPlanList) Validate() error {
	return dara.Validate(s)
}
