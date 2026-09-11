// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSynchronizationEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v *SwitchSynchronizationEndpointRequestEndpoint) *SwitchSynchronizationEndpointRequest
	GetEndpoint() *SwitchSynchronizationEndpointRequestEndpoint
	SetSourceEndpoint(v *SwitchSynchronizationEndpointRequestSourceEndpoint) *SwitchSynchronizationEndpointRequest
	GetSourceEndpoint() *SwitchSynchronizationEndpointRequestSourceEndpoint
	SetAccountId(v string) *SwitchSynchronizationEndpointRequest
	GetAccountId() *string
	SetOwnerId(v string) *SwitchSynchronizationEndpointRequest
	GetOwnerId() *string
	SetRegionId(v string) *SwitchSynchronizationEndpointRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SwitchSynchronizationEndpointRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *SwitchSynchronizationEndpointRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *SwitchSynchronizationEndpointRequest
	GetSynchronizationJobId() *string
}

type SwitchSynchronizationEndpointRequest struct {
	Endpoint       *SwitchSynchronizationEndpointRequestEndpoint       `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	SourceEndpoint *SwitchSynchronizationEndpointRequestSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The Alibaba Cloud account ID. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > Default value: **Forward**. The value **Reverse*	- takes effect only when the synchronization topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Instance ID of the data synchronization instance. You can call the DescribeSynchronizationJobs operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s SwitchSynchronizationEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequest) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequest) GetEndpoint() *SwitchSynchronizationEndpointRequestEndpoint {
	return s.Endpoint
}

func (s *SwitchSynchronizationEndpointRequest) GetSourceEndpoint() *SwitchSynchronizationEndpointRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *SwitchSynchronizationEndpointRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *SwitchSynchronizationEndpointRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SwitchSynchronizationEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchSynchronizationEndpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SwitchSynchronizationEndpointRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *SwitchSynchronizationEndpointRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *SwitchSynchronizationEndpointRequest) SetEndpoint(v *SwitchSynchronizationEndpointRequestEndpoint) *SwitchSynchronizationEndpointRequest {
	s.Endpoint = v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSourceEndpoint(v *SwitchSynchronizationEndpointRequestSourceEndpoint) *SwitchSynchronizationEndpointRequest {
	s.SourceEndpoint = v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetAccountId(v string) *SwitchSynchronizationEndpointRequest {
	s.AccountId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetOwnerId(v string) *SwitchSynchronizationEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetRegionId(v string) *SwitchSynchronizationEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetResourceGroupId(v string) *SwitchSynchronizationEndpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationDirection(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationJobId(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) Validate() error {
	if s.Endpoint != nil {
		if err := s.Endpoint.Validate(); err != nil {
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

type SwitchSynchronizationEndpointRequestEndpoint struct {
	// 新数据库的IP地址。
	//
	// > 当**Endpoint.InstanceType**取值为**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 172.15.185.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// ECS或专有网络的实例ID。
	//
	// > - 当**Endpoint.InstanceType**取值为**ECS**时，本参数需传入ECS实例的ID。
	//
	// - 当**Endpoint.InstanceType**取值为**Express**时，本参数需传入专有网络ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp11haem1kpkhoup****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 新数据库所属的实例类型，取值：
	//
	// - **LocalInstance**：有公网IP的自建数据库；
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **Express**：通过专线接入的自建数据库。
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 新的数据库服务端口。
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 待调整连接信息的实例，取值：
	//
	// - **Source**：源实例。
	//
	// - **Destination**：目标实例。
	//
	// This parameter is required.
	//
	// example:
	//
	// Destination
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SwitchSynchronizationEndpointRequestEndpoint) String() string {
	return dara.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequestEndpoint) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) GetIP() *string {
	return s.IP
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) GetPort() *string {
	return s.Port
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) GetType() *string {
	return s.Type
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetIP(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.IP = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceId(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceType = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetPort(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Port = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Type = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) Validate() error {
	return dara.Validate(s)
}

type SwitchSynchronizationEndpointRequestSourceEndpoint struct {
	// 当源实例与目标实例所属阿里云账号不同时，您需要传入该参数指定源实例的所属阿里云账号的ID。
	//
	// example:
	//
	// 14069264****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// 当源实例与目标实例所属阿里云账号不同时，需传入该参数，来指定源实例的授权角色，以允许目标实例阿里云账号访问源实例的实例信息。
	//
	// > 角色所需的权限及授权方式，请参见[跨阿里云账号数据迁移或同步时如何配置RAM授权](https://help.aliyun.com/document_detail/48468.html)。
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s SwitchSynchronizationEndpointRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) GetOwnerID() *string {
	return s.OwnerID
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) GetRole() *string {
	return s.Role
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) SetOwnerID(v string) *SwitchSynchronizationEndpointRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) SetRole(v string) *SwitchSynchronizationEndpointRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
