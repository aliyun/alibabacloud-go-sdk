// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterCount(v int64) *CreateHoloWarehouseRequest
	GetClusterCount() *int64
	SetCpu(v string) *CreateHoloWarehouseRequest
	GetCpu() *string
	SetName(v string) *CreateHoloWarehouseRequest
	GetName() *string
}

type CreateHoloWarehouseRequest struct {
	// The number of clusters.
	//
	// example:
	//
	// 2
	ClusterCount *int64 `json:"clusterCount,omitempty" xml:"clusterCount,omitempty"`
	// The specifications of the compute group. The value must be a multiple of 16, and the minimum value is 16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the compute group.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseRequest) GetClusterCount() *int64 {
	return s.ClusterCount
}

func (s *CreateHoloWarehouseRequest) GetCpu() *string {
	return s.Cpu
}

func (s *CreateHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *CreateHoloWarehouseRequest) SetClusterCount(v int64) *CreateHoloWarehouseRequest {
	s.ClusterCount = &v
	return s
}

func (s *CreateHoloWarehouseRequest) SetCpu(v string) *CreateHoloWarehouseRequest {
	s.Cpu = &v
	return s
}

func (s *CreateHoloWarehouseRequest) SetName(v string) *CreateHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *CreateHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
