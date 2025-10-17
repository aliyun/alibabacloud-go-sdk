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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// If you set the **SourceEndpoint.InstanceType*	- parameter to **DRDS**, you must specify the DBInstanceCount parameter. This parameter specifies the number of private RDS instances attached to the source PolarDB-X instance. Default value: **1**.
	//
	// example:
	//
	// 3
	DBInstanceCount *int32 `json:"DBInstanceCount,omitempty" xml:"DBInstanceCount,omitempty"`
	// The ID of the region where the destination database resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  If the **SourceRegion*	- parameter is set to the China (Hong Kong) region or a region outside the Chinese mainland, you must set the DestRegion parameter to the same region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the data synchronization instance.
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid*	- (default value): pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle of the subscription instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// >  You must specify this parameter only if you set the PayType parameter to **PrePaid**.
	//
	// example:
	//
	// Year
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource GroupId
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the region where the source database resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The specification of the data synchronization instance. Valid values: **micro**, **small**, **medium**, and **large**.
	//
	// >  For more information about the test performance of each specification, see [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// small
	SynchronizationJobClass *string `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	// The synchronization topology. Valid values:
	//
	// 	- **oneway**: one-way synchronization
	//
	// 	- **bidirectional**: two-way synchronization
	//
	// >
	//
	// 	- The default value is **oneway**.
	//
	// 	- This parameter can be set to **bidirectional*	- only when the **SourceEndpoint.InstanceType*	- and **DestinationEndpoint.InstanceType*	- parameters are set to **MySQL**, **PolarDB**, or **Redis**.
	//
	// example:
	//
	// oneway
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The subscription length.
	//
	// 	- If the billing cycle is **Year**, the value range is **1 to 5**.
	//
	// 	- If the billing cycle is **Month**, the value range is **1 to 60**.
	//
	// >  You must specify this parameter only if you set the PayType parameter to **PrePaid**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The network type. Valid value: **Intranet**, which indicates virtual private cloud (VPC).
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
	// The instance type of the destination database. Valid values:
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster or PolarDB O Edition cluster
	//
	// 	- **Redis**: Redis database
	//
	// 	- **MaxCompute**: MaxCompute project
	//
	// >
	//
	// 	- Default value: **MySQL**.
	//
	// 	- For more information about the supported source and destination databases, see [Database types, initial synchronization types, and synchronization topologies](https://help.aliyun.com/document_detail/130744.html).
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
	// The instance type of the source database. Valid values:
	//
	// 	- **MySQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster or PolarDB O Edition cluster
	//
	// 	- **Redis**: Redis database
	//
	// 	- **DRDS**: PolarDB-X instance V1.0
	//
	// >
	//
	// 	- Default value: **MySQL**.
	//
	// 	- For more information about the supported source and destination databases, see [Database types, initial synchronization types, and synchronization topologies](https://help.aliyun.com/document_detail/130744.html).
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
