// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarehouseScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticCu(v int64) *UpdateWarehouseScheduleTaskRequest
	GetElasticCu() *int64
	SetEndTime(v string) *UpdateWarehouseScheduleTaskRequest
	GetEndTime() *string
	SetId(v string) *UpdateWarehouseScheduleTaskRequest
	GetId() *string
	SetStartTime(v string) *UpdateWarehouseScheduleTaskRequest
	GetStartTime() *string
	SetWarehouseId(v int64) *UpdateWarehouseScheduleTaskRequest
	GetWarehouseId() *int64
}

type UpdateWarehouseScheduleTaskRequest struct {
	// example:
	//
	// 32
	ElasticCu *int64 `json:"elasticCu,omitempty" xml:"elasticCu,omitempty"`
	// example:
	//
	// 0400
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1980869072412614657
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0200
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WarehouseId *int64 `json:"warehouseId,omitempty" xml:"warehouseId,omitempty"`
}

func (s UpdateWarehouseScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarehouseScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarehouseScheduleTaskRequest) GetElasticCu() *int64 {
	return s.ElasticCu
}

func (s *UpdateWarehouseScheduleTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateWarehouseScheduleTaskRequest) GetId() *string {
	return s.Id
}

func (s *UpdateWarehouseScheduleTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateWarehouseScheduleTaskRequest) GetWarehouseId() *int64 {
	return s.WarehouseId
}

func (s *UpdateWarehouseScheduleTaskRequest) SetElasticCu(v int64) *UpdateWarehouseScheduleTaskRequest {
	s.ElasticCu = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskRequest) SetEndTime(v string) *UpdateWarehouseScheduleTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskRequest) SetId(v string) *UpdateWarehouseScheduleTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskRequest) SetStartTime(v string) *UpdateWarehouseScheduleTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskRequest) SetWarehouseId(v int64) *UpdateWarehouseScheduleTaskRequest {
	s.WarehouseId = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
