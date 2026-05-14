// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteWarehouseScheduleTaskRequest
	GetId() *string
	SetWarehouseId(v string) *DeleteWarehouseScheduleTaskRequest
	GetWarehouseId() *string
}

type DeleteWarehouseScheduleTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2034449120420339713
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	WarehouseId *string `json:"warehouseId,omitempty" xml:"warehouseId,omitempty"`
}

func (s DeleteWarehouseScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseScheduleTaskRequest) GetId() *string {
	return s.Id
}

func (s *DeleteWarehouseScheduleTaskRequest) GetWarehouseId() *string {
	return s.WarehouseId
}

func (s *DeleteWarehouseScheduleTaskRequest) SetId(v string) *DeleteWarehouseScheduleTaskRequest {
	s.Id = &v
	return s
}

func (s *DeleteWarehouseScheduleTaskRequest) SetWarehouseId(v string) *DeleteWarehouseScheduleTaskRequest {
	s.WarehouseId = &v
	return s
}

func (s *DeleteWarehouseScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
