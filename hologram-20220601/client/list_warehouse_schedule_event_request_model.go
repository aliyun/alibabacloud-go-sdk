// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticType(v string) *ListWarehouseScheduleEventRequest
	GetElasticType() *string
	SetEndTime(v int64) *ListWarehouseScheduleEventRequest
	GetEndTime() *int64
	SetPageNumber(v int64) *ListWarehouseScheduleEventRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListWarehouseScheduleEventRequest
	GetPageSize() *int64
	SetStartTime(v int64) *ListWarehouseScheduleEventRequest
	GetStartTime() *int64
}

type ListWarehouseScheduleEventRequest struct {
	// example:
	//
	// timed
	ElasticType *string `json:"elasticType,omitempty" xml:"elasticType,omitempty"`
	// example:
	//
	// 1777516201
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1777257001
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListWarehouseScheduleEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleEventRequest) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleEventRequest) GetElasticType() *string {
	return s.ElasticType
}

func (s *ListWarehouseScheduleEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListWarehouseScheduleEventRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListWarehouseScheduleEventRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWarehouseScheduleEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListWarehouseScheduleEventRequest) SetElasticType(v string) *ListWarehouseScheduleEventRequest {
	s.ElasticType = &v
	return s
}

func (s *ListWarehouseScheduleEventRequest) SetEndTime(v int64) *ListWarehouseScheduleEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListWarehouseScheduleEventRequest) SetPageNumber(v int64) *ListWarehouseScheduleEventRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWarehouseScheduleEventRequest) SetPageSize(v int64) *ListWarehouseScheduleEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListWarehouseScheduleEventRequest) SetStartTime(v int64) *ListWarehouseScheduleEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListWarehouseScheduleEventRequest) Validate() error {
	return dara.Validate(s)
}
