// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateWarehouseRequest
	GetRegionId() *string
	SetConfig(v string) *CreateWarehouseRequest
	GetConfig() *string
	SetCpu(v string) *CreateWarehouseRequest
	GetCpu() *string
	SetWarehouseName(v string) *CreateWarehouseRequest
	GetWarehouseName() *string
}

type CreateWarehouseRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration information.
	//
	// example:
	//
	// ""
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 32
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the virtual warehouse.
	//
	// example:
	//
	// warehouse-test
	WarehouseName *string `json:"warehouseName,omitempty" xml:"warehouseName,omitempty"`
}

func (s CreateWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseRequest) GoString() string {
	return s.String()
}

func (s *CreateWarehouseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWarehouseRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateWarehouseRequest) GetCpu() *string {
	return s.Cpu
}

func (s *CreateWarehouseRequest) GetWarehouseName() *string {
	return s.WarehouseName
}

func (s *CreateWarehouseRequest) SetRegionId(v string) *CreateWarehouseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWarehouseRequest) SetConfig(v string) *CreateWarehouseRequest {
	s.Config = &v
	return s
}

func (s *CreateWarehouseRequest) SetCpu(v string) *CreateWarehouseRequest {
	s.Cpu = &v
	return s
}

func (s *CreateWarehouseRequest) SetWarehouseName(v string) *CreateWarehouseRequest {
	s.WarehouseName = &v
	return s
}

func (s *CreateWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
