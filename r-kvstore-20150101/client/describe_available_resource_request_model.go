// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeAvailableResourceRequest
	GetAcceptLanguage() *string
	SetEngine(v string) *DescribeAvailableResourceRequest
	GetEngine() *string
	SetInstanceChargeType(v string) *DescribeAvailableResourceRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *DescribeAvailableResourceRequest
	GetInstanceId() *string
	SetInstanceScene(v string) *DescribeAvailableResourceRequest
	GetInstanceScene() *string
	SetNodeId(v string) *DescribeAvailableResourceRequest
	GetNodeId() *string
	SetOrderType(v string) *DescribeAvailableResourceRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableResourceRequest
	GetOwnerId() *int64
	SetProductType(v string) *DescribeAvailableResourceRequest
	GetProductType() *string
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAvailableResourceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeAvailableResourceRequest
	GetSecurityToken() *string
	SetZoneId(v string) *DescribeAvailableResourceRequest
	GetZoneId() *string
}

type DescribeAvailableResourceRequest struct {
	// The display language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese. This is the default value.
	//
	// 	- **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **Redis**
	//
	// 	- **Memcache**
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid*	- (default): subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the instance.
	//
	// > This parameter is available and required only if the **OrderType*	- parameter is set to **UPGRADE*	- or **DOWNGRADE**.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **professional**: Standard Edition. This edition supports the standalone, master-replica, read /write splitting, and cluster architectures and provides high scalability.
	//
	// example:
	//
	// professional
	InstanceScene *string `json:"InstanceScene,omitempty" xml:"InstanceScene,omitempty"`
	// The ID of the data node for which you want to query available resources that can be created. You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query the ID of the data node. Remove the number sign (`#`) and the content that follows the number sign. For example, retain only r-bp10noxlhcoim2\\*\\*\\*\\*-db-0.
	//
	// > Before you specify this parameter, you must set the **InstanceId*	- parameter to the ID of an instance that uses the cluster or read/write splitting architecture.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order type. Valid values:
	//
	// 	- **BUY*	- (default): orders that are used to create instances
	//
	// 	- **UPGRADE**: orders that are used to upgrade instances
	//
	// 	- **DOWNGRADE**: orders that are used to downgrade instances
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance type. Default value: Local. Valid values:
	//
	// 	- **Local**: classic Redis Open-Source Edition instance or classic DRAM-based instance
	//
	// 	- **Tair_rdb**: cloud-native DRAM-based instance
	//
	// 	- **Tair_scm**: persistent memory-optimized instance
	//
	// 	- **Tair_essd**: ESSD/SSD-based instance
	//
	// 	- **OnECS**: cloud-native Redis Open-Source Edition instance
	//
	// >  The default value of this parameter is Local. To query disk resources, you must specify the instance type that provides the required disk resources.
	//
	// example:
	//
	// Local
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the IDs of resource groups.
	//
	// >  You can also query the IDs of resource groups in the Resource Management console. For more information, see [View basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4e******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The zone ID of the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeAvailableResourceRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableResourceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAvailableResourceRequest) GetInstanceScene() *string {
	return s.InstanceScene
}

func (s *DescribeAvailableResourceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAvailableResourceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeAvailableResourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableResourceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableResourceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceRequest) SetAcceptLanguage(v string) *DescribeAvailableResourceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngine(v string) *DescribeAvailableResourceRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceId(v string) *DescribeAvailableResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetInstanceScene(v string) *DescribeAvailableResourceRequest {
	s.InstanceScene = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetNodeId(v string) *DescribeAvailableResourceRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOrderType(v string) *DescribeAvailableResourceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetProductType(v string) *DescribeAvailableResourceRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceGroupId(v string) *DescribeAvailableResourceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSecurityToken(v string) *DescribeAvailableResourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
