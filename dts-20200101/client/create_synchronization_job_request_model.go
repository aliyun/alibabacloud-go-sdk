// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationEndpoint(v *CreateSynchronizationJobRequestDestinationEndpoint) *CreateSynchronizationJobRequest
	GetDestinationEndpoint() *CreateSynchronizationJobRequestDestinationEndpoint
	SetSourceEndpoint(v *CreateSynchronizationJobRequestSourceEndpoint) *CreateSynchronizationJobRequest
	GetSourceEndpoint() *CreateSynchronizationJobRequestSourceEndpoint
	SetAccountId(v string) *CreateSynchronizationJobRequest
	GetAccountId() *string
	SetClientToken(v string) *CreateSynchronizationJobRequest
	GetClientToken() *string
	SetDBInstanceCount(v int32) *CreateSynchronizationJobRequest
	GetDBInstanceCount() *int32
	SetDestRegion(v string) *CreateSynchronizationJobRequest
	GetDestRegion() *string
	SetOwnerId(v string) *CreateSynchronizationJobRequest
	GetOwnerId() *string
	SetPayType(v string) *CreateSynchronizationJobRequest
	GetPayType() *string
	SetPeriod(v string) *CreateSynchronizationJobRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSynchronizationJobRequest
	GetResourceGroupId() *string
	SetSourceRegion(v string) *CreateSynchronizationJobRequest
	GetSourceRegion() *string
	SetSynchronizationJobClass(v string) *CreateSynchronizationJobRequest
	GetSynchronizationJobClass() *string
	SetTopology(v string) *CreateSynchronizationJobRequest
	GetTopology() *string
	SetUsedTime(v int32) *CreateSynchronizationJobRequest
	GetUsedTime() *int32
	SetNetworkType(v string) *CreateSynchronizationJobRequest
	GetNetworkType() *string
}

type CreateSynchronizationJobRequest struct {
	DestinationEndpoint *CreateSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	SourceEndpoint      *CreateSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a value from your client to ensure uniqueness across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of private custom ApsaraDB RDS instances attached to the source PolarDB-X instance. This parameter is required when **SourceEndpoint.InstanceType*	- is set to **DRDS**. Default value: **1**.
	//
	// example:
	//
	// 1
	DBInstanceCount *int32 `json:"DBInstanceCount,omitempty" xml:"DBInstanceCount,omitempty"`
	// The region ID of the destination database for data synchronization. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the region specified by the **SourceRegion*	- parameter is Hong Kong (China) or a region outside China, set this parameter to the same region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing method of the subscription instance. Valid values:
	//
	// - **Year**: annual subscription.
	//
	// - **Month**: monthly subscription.
	//
	// > This parameter is valid and required only when **PayType*	- is set to **PrePaid*	- (subscription).
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the data synchronization instance. Set this parameter to the same value as the **DestRegion*	- parameter.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The region ID of the source database for data synchronization. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The specification of the data synchronization link. Valid values: **micro**, **small**, **medium**, **large**.
	//
	// > For more information about the description and performance test results of each specification, see [Specifications of data synchronization links](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// small
	SynchronizationJobClass *string `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	// The synchronization topology. Valid values:
	//
	// - **oneway**: one-way synchronization.
	//
	// - **bidirectional**: two-way synchronization.
	//
	// > - Default value: **oneway**.
	//
	// - You can set this parameter to **bidirectional*	- only when both **SourceEndpoint.InstanceType*	- and **DestinationEndpoint.InstanceType*	- are set to **MySQL**, **PolarDB**, or **Redis**.
	//
	// example:
	//
	// oneway
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The subscription duration of the subscription instance.
	//
	// - If the billing method is set to **Year**, valid values are **1 to 5**.
	//
	// - If the billing method is set to **Month**, valid values are **1 to 60**.
	//
	// > This parameter is valid and required only when **PayType*	- is set to **PrePaid*	- (subscription).
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The network type for Data Transmission Service. Set the value to **Intranet*	- (Express Connect).
	//
	// example:
	//
	// Intranet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s CreateSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequest) GetDestinationEndpoint() *CreateSynchronizationJobRequestDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *CreateSynchronizationJobRequest) GetSourceEndpoint() *CreateSynchronizationJobRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *CreateSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateSynchronizationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSynchronizationJobRequest) GetDBInstanceCount() *int32 {
	return s.DBInstanceCount
}

func (s *CreateSynchronizationJobRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *CreateSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateSynchronizationJobRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateSynchronizationJobRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSynchronizationJobRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CreateSynchronizationJobRequest) GetSynchronizationJobClass() *string {
	return s.SynchronizationJobClass
}

func (s *CreateSynchronizationJobRequest) GetTopology() *string {
	return s.Topology
}

func (s *CreateSynchronizationJobRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateSynchronizationJobRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateSynchronizationJobRequest) SetDestinationEndpoint(v *CreateSynchronizationJobRequestDestinationEndpoint) *CreateSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSourceEndpoint(v *CreateSynchronizationJobRequestSourceEndpoint) *CreateSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetAccountId(v string) *CreateSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetClientToken(v string) *CreateSynchronizationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDBInstanceCount(v int32) *CreateSynchronizationJobRequest {
	s.DBInstanceCount = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDestRegion(v string) *CreateSynchronizationJobRequest {
	s.DestRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetOwnerId(v string) *CreateSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetPayType(v string) *CreateSynchronizationJobRequest {
	s.PayType = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetPeriod(v string) *CreateSynchronizationJobRequest {
	s.Period = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetRegionId(v string) *CreateSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetResourceGroupId(v string) *CreateSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSourceRegion(v string) *CreateSynchronizationJobRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSynchronizationJobClass(v string) *CreateSynchronizationJobRequest {
	s.SynchronizationJobClass = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetTopology(v string) *CreateSynchronizationJobRequest {
	s.Topology = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetUsedTime(v int32) *CreateSynchronizationJobRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetNetworkType(v string) *CreateSynchronizationJobRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateSynchronizationJobRequest) Validate() error {
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSynchronizationJobRequestDestinationEndpoint struct {
	// 目标库的实例类型，取值：
	//
	// - **MySQL**：MySQL数据库（包括RDS MySQL和自建MySQL）。
	//
	// - **PolarDB**：PolarDB集群（仅支持MySQL或兼容Oracle语法的引擎）。
	//
	// - **Redis**：Redis数据库。
	//
	// - **MaxCompute**：MaxCompute实例。
	//
	// >- 默认取值为**MySQL**。
	//
	// - 关于支持的源库和目标库对应情况，请参见支持的[数据库、同步初始化类型和同步拓扑](https://help.aliyun.com/document_detail/130744.html)。
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSynchronizationJobRequestDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequestDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *CreateSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *CreateSynchronizationJobRequestDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type CreateSynchronizationJobRequestSourceEndpoint struct {
	// 源库的实例类型，取值：
	//
	// - **MySQL**：MySQL数据库（包括RDS MySQL和自建MySQL）。
	//
	// - **PolarDB**：PolarDB集群（仅支持MySQL或兼容Oracle语法的引擎）。
	//
	// - **Redis**：Redis数据库。
	//
	// - **DRDS**：云原生分布式数据库PolarDB-X 1.0。
	//
	// > - 默认取值为**MySQL**。
	//
	// - 关于支持的源库和目标库对应情况，请参见支持的[数据库、同步初始化类型和同步拓扑](https://help.aliyun.com/document_detail/130744.html)。
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSynchronizationJobRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequestSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *CreateSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *CreateSynchronizationJobRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
