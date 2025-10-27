// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEIURangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v string) *DescribeEIURangeRequest
	GetComputeResource() *string
	SetDBClusterId(v string) *DescribeEIURangeRequest
	GetDBClusterId() *string
	SetDBClusterVersion(v string) *DescribeEIURangeRequest
	GetDBClusterVersion() *string
	SetOperation(v string) *DescribeEIURangeRequest
	GetOperation() *string
	SetOwnerAccount(v string) *DescribeEIURangeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEIURangeRequest
	GetOwnerId() *int64
	SetProductVersion(v string) *DescribeEIURangeRequest
	GetProductVersion() *string
	SetRegionId(v string) *DescribeEIURangeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEIURangeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeEIURangeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEIURangeRequest
	GetResourceOwnerId() *int64
	SetStorageSize(v string) *DescribeEIURangeRequest
	GetStorageSize() *string
	SetSubOperation(v string) *DescribeEIURangeRequest
	GetSubOperation() *string
	SetZoneId(v string) *DescribeEIURangeRequest
	GetZoneId() *string
}

type DescribeEIURangeRequest struct {
	// The specifications of computing resources.
	//
	// >  You can call the [DescribeComputeResource](https://help.aliyun.com/document_detail/469002.html) operation to query the specifications of computing resources.
	//
	// example:
	//
	// {
	//
	//       "RealValue": "32Core128GBNEW",
	//
	//       "DisplayValue": "32Core128GB"
	//
	//     }
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// 	- This parameter can be left empty when **Operation*	- is set to **Buy**.
	//
	// 	- This parameter must be specified when **Operation*	- is set to **Upgrade*	- or **Downgrade**.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// example:
	//
	// am-bp16t5ci7r74s****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The version of the AnalyticDB for MySQL Data Warehouse Edition cluster. Set the value to **3.0**.
	//
	// example:
	//
	// 3.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- **Buy**: purchases a cluster.
	//
	// 	- **Modify**: changes configurations of a cluster.
	//
	// example:
	//
	// Buy
	Operation      *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specifications of storage resources. Default value: **8ACU**. Valid values:
	//
	// 	- **8ACU**
	//
	// 	- **12ACU**
	//
	// 	- **16ACU**
	//
	// example:
	//
	// 8ACU
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The type of the sub-operation. Valid values:
	//
	// 	- **Upgrade**: upgrades a cluster.
	//
	// 	- **Downgrade**: downgrades a cluster.
	//
	// example:
	//
	// Upgrade
	SubOperation *string `json:"SubOperation,omitempty" xml:"SubOperation,omitempty"`
	// The zone ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612293.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeEIURangeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEIURangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeRequest) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeEIURangeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEIURangeRequest) GetDBClusterVersion() *string {
	return s.DBClusterVersion
}

func (s *DescribeEIURangeRequest) GetOperation() *string {
	return s.Operation
}

func (s *DescribeEIURangeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEIURangeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEIURangeRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeEIURangeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEIURangeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEIURangeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEIURangeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEIURangeRequest) GetStorageSize() *string {
	return s.StorageSize
}

func (s *DescribeEIURangeRequest) GetSubOperation() *string {
	return s.SubOperation
}

func (s *DescribeEIURangeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeEIURangeRequest) SetComputeResource(v string) *DescribeEIURangeRequest {
	s.ComputeResource = &v
	return s
}

func (s *DescribeEIURangeRequest) SetDBClusterId(v string) *DescribeEIURangeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetDBClusterVersion(v string) *DescribeEIURangeRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOperation(v string) *DescribeEIURangeRequest {
	s.Operation = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOwnerAccount(v string) *DescribeEIURangeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEIURangeRequest) SetOwnerId(v int64) *DescribeEIURangeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetProductVersion(v string) *DescribeEIURangeRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeEIURangeRequest) SetRegionId(v string) *DescribeEIURangeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceGroupId(v string) *DescribeEIURangeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceOwnerAccount(v string) *DescribeEIURangeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEIURangeRequest) SetResourceOwnerId(v int64) *DescribeEIURangeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEIURangeRequest) SetStorageSize(v string) *DescribeEIURangeRequest {
	s.StorageSize = &v
	return s
}

func (s *DescribeEIURangeRequest) SetSubOperation(v string) *DescribeEIURangeRequest {
	s.SubOperation = &v
	return s
}

func (s *DescribeEIURangeRequest) SetZoneId(v string) *DescribeEIURangeRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeEIURangeRequest) Validate() error {
	return dara.Validate(s)
}
