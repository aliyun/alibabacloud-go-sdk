// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWarehouseScheduleTaskRequest
	GetDescription() *string
	SetElasticCu(v int64) *CreateWarehouseScheduleTaskRequest
	GetElasticCu() *int64
	SetEndTime(v string) *CreateWarehouseScheduleTaskRequest
	GetEndTime() *string
	SetStartTime(v string) *CreateWarehouseScheduleTaskRequest
	GetStartTime() *string
	SetWarehouseId(v int64) *CreateWarehouseScheduleTaskRequest
	GetWarehouseId() *int64
}

type CreateWarehouseScheduleTaskRequest struct {
	// The description of the time-based elastic plan.
	//
	// example:
	//
	// user stat
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of elastic CPU cores. The value must be a multiple of 16, greater than or equal to 16, and less than or equal to the specifications of the current compute group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32
	ElasticCu *int64 `json:"elasticCu,omitempty" xml:"elasticCu,omitempty"`
	// The end time. The format is HHmm (four digits, 24-hour clock). The time must be on the hour or half hour, such as 0030, 0100, or 0400.
	//
	// The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0200
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time. The format is HHmm (four digits, 24-hour clock). The time must be on the hour or half hour, such as 0030, 0100, or 0400.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0100
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the compute group. You can call ListWarehouses to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WarehouseId *int64 `json:"warehouseId,omitempty" xml:"warehouseId,omitempty"`
}

func (s CreateWarehouseScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWarehouseScheduleTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWarehouseScheduleTaskRequest) GetElasticCu() *int64 {
	return s.ElasticCu
}

func (s *CreateWarehouseScheduleTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateWarehouseScheduleTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateWarehouseScheduleTaskRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
}

func (s *CreateWarehouseScheduleTaskRequest) SetDescription(v string) *CreateWarehouseScheduleTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateWarehouseScheduleTaskRequest) SetElasticCu(v int64) *CreateWarehouseScheduleTaskRequest {
	s.ElasticCu = &v
	return s
}

func (s *CreateWarehouseScheduleTaskRequest) SetEndTime(v string) *CreateWarehouseScheduleTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateWarehouseScheduleTaskRequest) SetStartTime(v string) *CreateWarehouseScheduleTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateWarehouseScheduleTaskRequest) SetWarehouseId(v int64) *CreateWarehouseScheduleTaskRequest {
	s.WarehouseId = &v
	return s
}

func (s *CreateWarehouseScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
