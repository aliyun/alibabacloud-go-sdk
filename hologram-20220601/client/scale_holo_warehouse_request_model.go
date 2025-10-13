// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v int64) *ScaleHoloWarehouseRequest
	GetCpu() *int64
	SetName(v string) *ScaleHoloWarehouseRequest
	GetName() *string
}

type ScaleHoloWarehouseRequest struct {
	// The specifications of the virtual warehouse. The number of vCPUs must be an integer multiple of 16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 64
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ScaleHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *ScaleHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *ScaleHoloWarehouseRequest) SetCpu(v int64) *ScaleHoloWarehouseRequest {
	s.Cpu = &v
	return s
}

func (s *ScaleHoloWarehouseRequest) SetName(v string) *ScaleHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *ScaleHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
