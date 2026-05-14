// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*ListWarehouseScheduleEventResponseBodyEventList) *ListWarehouseScheduleEventResponseBody
	GetEventList() []*ListWarehouseScheduleEventResponseBodyEventList
	SetPageNumber(v int64) *ListWarehouseScheduleEventResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListWarehouseScheduleEventResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWarehouseScheduleEventResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListWarehouseScheduleEventResponseBody
	GetTotalCount() *int64
}

type ListWarehouseScheduleEventResponseBody struct {
	EventList []*ListWarehouseScheduleEventResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// RequestId
	//
	// example:
	//
	// E16D32D4-DF86-1180-8220-0D39770A5AF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 120
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWarehouseScheduleEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleEventResponseBody) GetEventList() []*ListWarehouseScheduleEventResponseBodyEventList {
	return s.EventList
}

func (s *ListWarehouseScheduleEventResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListWarehouseScheduleEventResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWarehouseScheduleEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarehouseScheduleEventResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWarehouseScheduleEventResponseBody) SetEventList(v []*ListWarehouseScheduleEventResponseBodyEventList) *ListWarehouseScheduleEventResponseBody {
	s.EventList = v
	return s
}

func (s *ListWarehouseScheduleEventResponseBody) SetPageNumber(v int64) *ListWarehouseScheduleEventResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBody) SetPageSize(v int64) *ListWarehouseScheduleEventResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBody) SetRequestId(v string) *ListWarehouseScheduleEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBody) SetTotalCount(v int64) *ListWarehouseScheduleEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBody) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWarehouseScheduleEventResponseBodyEventList struct {
	// example:
	//
	// 2
	ClusterCount *int64 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// example:
	//
	// 32
	ClusterCpu *int64 `json:"ClusterCpu,omitempty" xml:"ClusterCpu,omitempty"`
	// example:
	//
	// 48
	ElasticCpu *int64 `json:"ElasticCpu,omitempty" xml:"ElasticCpu,omitempty"`
	// example:
	//
	// ScaleUp
	ElasticType *string `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	// example:
	//
	// AlterWarehouse
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 2024-07-22T09:43:02.638Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// insufficient resource
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// example:
	//
	// 2
	InitClusterCount *int64 `json:"InitClusterCount,omitempty" xml:"InitClusterCount,omitempty"`
	// example:
	//
	// 0
	OriginalElasticCpu *int64 `json:"OriginalElasticCpu,omitempty" xml:"OriginalElasticCpu,omitempty"`
	// example:
	//
	// 64
	ReservedCpu *int64 `json:"ReservedCpu,omitempty" xml:"ReservedCpu,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// init_warehouse
	WarehouseName *string `json:"WarehouseName,omitempty" xml:"WarehouseName,omitempty"`
}

func (s ListWarehouseScheduleEventResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleEventResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetClusterCount() *int64 {
	return s.ClusterCount
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetClusterCpu() *int64 {
	return s.ClusterCpu
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetElasticCpu() *int64 {
	return s.ElasticCpu
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetElasticType() *string {
	return s.ElasticType
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetEventName() *string {
	return s.EventName
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetEventTime() *string {
	return s.EventTime
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetFailedReason() *string {
	return s.FailedReason
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetInitClusterCount() *int64 {
	return s.InitClusterCount
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetOriginalElasticCpu() *int64 {
	return s.OriginalElasticCpu
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetReservedCpu() *int64 {
	return s.ReservedCpu
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetResult() *string {
	return s.Result
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) GetWarehouseName() *string {
	return s.WarehouseName
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetClusterCount(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.ClusterCount = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetClusterCpu(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.ClusterCpu = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetElasticCpu(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.ElasticCpu = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetElasticType(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.ElasticType = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetEventName(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.EventName = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetEventTime(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.EventTime = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetFailedReason(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.FailedReason = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetInitClusterCount(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.InitClusterCount = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetOriginalElasticCpu(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.OriginalElasticCpu = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetReservedCpu(v int64) *ListWarehouseScheduleEventResponseBodyEventList {
	s.ReservedCpu = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetResult(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.Result = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) SetWarehouseName(v string) *ListWarehouseScheduleEventResponseBodyEventList {
	s.WarehouseName = &v
	return s
}

func (s *ListWarehouseScheduleEventResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}
