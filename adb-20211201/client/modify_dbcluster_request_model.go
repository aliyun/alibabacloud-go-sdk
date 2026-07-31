// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAINodeNumber(v int32) *ModifyDBClusterRequest
	GetAINodeNumber() *int32
	SetAINodeSpec(v string) *ModifyDBClusterRequest
	GetAINodeSpec() *string
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
	AINodeNumber *int32 `json:"AINodeNumber,omitempty" xml:"AINodeNumber,omitempty"`
	// example:
	//
	// ADB.MLPlus.4
	AINodeSpec *string `json:"AINodeSpec,omitempty" xml:"AINodeSpec,omitempty"`
	// The compute reserved resources. Valid values: 0 ACU to 4096 ACU, in increments of 16. 1 ACU is approximately equivalent to 1 core and 4 GB of memory.
	//
	// > Include the unit when you specify this parameter.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID of a Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to allocate all compute reserved resources to the default resource group (user_default). Valid values:
	//
	// - true (default): All compute reserved resources are allocated to the default resource group.
	//
	// - false: Not all compute reserved resources are allocated to the default resource group.
	//
	// example:
	//
	// true
	EnableDefaultResourcePool *bool `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	// The product form. Valid values:
	//
	// - **IntegrationForm**: integrated form.
	//
	// - **LegacyForm**: Data Lakehouse Edition.
	//
	// example:
	//
	// LegacyForm
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of reserved nodes.
	//
	// - Enterprise Edition: The default value is 3. The value increases in increments of 3.
	//
	// - Basic Edition: The default value is 1.
	//
	// > This parameter is required only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// 3
	ReservedNodeCount *int32 `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// The node specifications of storage reserved resources. Valid values: 8ACU, 12ACU, and 16ACU.
	//
	// > Include the unit when you specify this parameter. This parameter is required only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// 8ACU
	ReservedNodeSize *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	// The storage reserved resources. Valid values: 0 ACU to 2064 ACU, in increments of 24. 1 ACU is approximately equivalent to 1 core and 4 GB of memory.
	//
	// > Include the unit when you specify this parameter.
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

func (s *ModifyDBClusterRequest) GetAINodeNumber() *int32 {
	return s.AINodeNumber
}

func (s *ModifyDBClusterRequest) GetAINodeSpec() *string {
	return s.AINodeSpec
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

func (s *ModifyDBClusterRequest) SetAINodeNumber(v int32) *ModifyDBClusterRequest {
	s.AINodeNumber = &v
	return s
}

func (s *ModifyDBClusterRequest) SetAINodeSpec(v string) *ModifyDBClusterRequest {
	s.AINodeSpec = &v
	return s
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
