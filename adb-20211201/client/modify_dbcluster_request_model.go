// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v string) *ModifyDBClusterRequest
	GetComputeResource() *string
	SetDBClusterId(v string) *ModifyDBClusterRequest
	GetDBClusterId() *string
	SetEnableDefaultResourcePool(v bool) *ModifyDBClusterRequest
	GetEnableDefaultResourcePool() *bool
	SetProductForm(v string) *ModifyDBClusterRequest
	GetProductForm() *string
	SetRegionId(v string) *ModifyDBClusterRequest
	GetRegionId() *string
	SetReservedNodeCount(v int32) *ModifyDBClusterRequest
	GetReservedNodeCount() *int32
	SetReservedNodeSize(v string) *ModifyDBClusterRequest
	GetReservedNodeSize() *string
	SetStorageResource(v string) *ModifyDBClusterRequest
	GetStorageResource() *string
}

type ModifyDBClusterRequest struct {
	// The reserved computing resources. Valid values: 0ACU to 4096ACU. The value must be in increments of 16ACU. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to allocate all reserved computing resources to the user_default resource group. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableDefaultResourcePool *bool `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	// example:
	//
	// LegacyForm
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservedNodeCount *int32  `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// example:
	//
	// LegacyForm
	ReservedNodeSize *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	// The reserved storage resources. Valid values: 0ACU to 2064ACU. The value must be in increments of 24ACU. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *ModifyDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterRequest) GetEnableDefaultResourcePool() *bool {
	return s.EnableDefaultResourcePool
}

func (s *ModifyDBClusterRequest) GetProductForm() *string {
	return s.ProductForm
}

func (s *ModifyDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterRequest) GetReservedNodeCount() *int32 {
	return s.ReservedNodeCount
}

func (s *ModifyDBClusterRequest) GetReservedNodeSize() *string {
	return s.ReservedNodeSize
}

func (s *ModifyDBClusterRequest) GetStorageResource() *string {
	return s.StorageResource
}

func (s *ModifyDBClusterRequest) SetComputeResource(v string) *ModifyDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetEnableDefaultResourcePool(v bool) *ModifyDBClusterRequest {
	s.EnableDefaultResourcePool = &v
	return s
}

func (s *ModifyDBClusterRequest) SetProductForm(v string) *ModifyDBClusterRequest {
	s.ProductForm = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetReservedNodeCount(v int32) *ModifyDBClusterRequest {
	s.ReservedNodeCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetReservedNodeSize(v string) *ModifyDBClusterRequest {
	s.ReservedNodeSize = &v
	return s
}

func (s *ModifyDBClusterRequest) SetStorageResource(v string) *ModifyDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *ModifyDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
