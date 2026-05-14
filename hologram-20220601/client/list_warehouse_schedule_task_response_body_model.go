// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWarehouseScheduleTaskResponseBody
	GetRequestId() *string
	SetScheduleTaskList(v []*ListWarehouseScheduleTaskResponseBodyScheduleTaskList) *ListWarehouseScheduleTaskResponseBody
	GetScheduleTaskList() []*ListWarehouseScheduleTaskResponseBodyScheduleTaskList
}

type ListWarehouseScheduleTaskResponseBody struct {
	// example:
	//
	// A0A16C46-5B56-1F9B-AA37-4C3EAD95AAA8
	RequestId        *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduleTaskList []*ListWarehouseScheduleTaskResponseBodyScheduleTaskList `json:"ScheduleTaskList,omitempty" xml:"ScheduleTaskList,omitempty" type:"Repeated"`
}

func (s ListWarehouseScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarehouseScheduleTaskResponseBody) GetScheduleTaskList() []*ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	return s.ScheduleTaskList
}

func (s *ListWarehouseScheduleTaskResponseBody) SetRequestId(v string) *ListWarehouseScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBody) SetScheduleTaskList(v []*ListWarehouseScheduleTaskResponseBodyScheduleTaskList) *ListWarehouseScheduleTaskResponseBody {
	s.ScheduleTaskList = v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBody) Validate() error {
	if s.ScheduleTaskList != nil {
		for _, item := range s.ScheduleTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWarehouseScheduleTaskResponseBodyScheduleTaskList struct {
	// example:
	//
	// timed
	ElasticType *string                                                       `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	Plans       []*ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	// example:
	//
	// 32
	ReservedCpu *int64 `json:"ReservedCpu,omitempty" xml:"ReservedCpu,omitempty"`
	// example:
	//
	// kRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	WarehouseId *string `json:"WarehouseId,omitempty" xml:"WarehouseId,omitempty"`
	// example:
	//
	// init_warehouse
	WarehouseName *string `json:"WarehouseName,omitempty" xml:"WarehouseName,omitempty"`
}

func (s ListWarehouseScheduleTaskResponseBodyScheduleTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetElasticType() *string {
	return s.ElasticType
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetPlans() []*ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	return s.Plans
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetReservedCpu() *int64 {
	return s.ReservedCpu
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetWarehouseId() *string {
	return s.WarehouseId
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) GetWarehouseName() *string {
	return s.WarehouseName
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetElasticType(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.ElasticType = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetPlans(v []*ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.Plans = v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetReservedCpu(v int64) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.ReservedCpu = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetStatus(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.Status = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetWarehouseId(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.WarehouseId = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) SetWarehouseName(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskList {
	s.WarehouseName = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskList) Validate() error {
	if s.Plans != nil {
		for _, item := range s.Plans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans struct {
	// example:
	//
	// bill stat
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 32
	ElasticCu *int64 `json:"ElasticCu,omitempty" xml:"ElasticCu,omitempty"`
	// example:
	//
	// 0400
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1802985780260052993
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0100
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GetDescription() *string {
	return s.Description
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GetElasticCu() *int64 {
	return s.ElasticCu
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GetId() *string {
	return s.Id
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) SetDescription(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	s.Description = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) SetElasticCu(v int64) *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	s.ElasticCu = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) SetEndTime(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	s.EndTime = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) SetId(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	s.Id = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) SetStartTime(v string) *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans {
	s.StartTime = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponseBodyScheduleTaskListPlans) Validate() error {
	return dara.Validate(s)
}
