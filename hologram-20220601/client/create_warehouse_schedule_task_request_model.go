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
	// example:
	//
	// user stat
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 32
	ElasticCu *int64 `json:"elasticCu,omitempty" xml:"elasticCu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0200
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0100
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
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
