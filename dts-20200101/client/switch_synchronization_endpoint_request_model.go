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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >  Default value: **Forward**.
	//
	// The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the DescribeSynchronizationJobs operation to query the instance ID.
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
	// The IP address of the database.
	//
	// >  You must specify the IP address only if the **Endpoint.InstanceType*	- parameter is set to **Express**.
	//
	// example:
	//
	// 172.15.185.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the ECS instance or the virtual private cloud (VPC).
	//
	// >
	//
	// 	- If the **Endpoint.InstanceType*	- parameter is set to **ECS**, you must specify the ID of the ECS instance.
	//
	// 	- If the **Endpoint.InstanceType*	- parameter is set to **Express**, you must specify the ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp11haem1kpkhoup****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the database. Valid values:
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **ECS**: self-managed database that is hosted on ECS
	//
	// 	- **Express**: self-managed database that is connected over Express Connect
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The service port number of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Specifies whether to update the connection settings of the source instance or the destination instance. Valid values:
	//
	// 	- **Source**
	//
	// 	- **Destination**
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
	// The ID of the Alibaba Cloud account to which the source instance belongs. You must specify this parameter only if the source instance and the destination instance belong to different Alibaba Cloud accounts.
	//
	// example:
	//
	// 14069264****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The authorized Resource Access Management (RAM) role of the source instance. You must specify the RAM role only if the source instance and the destination instance belong to different Alibaba Cloud accounts. You can use the RAM role to allow the Alibaba Cloud account that owns the destination instance to access the source instance.
	//
	// >  For information about the permissions and authorization methods of the RAM role, see [Configure RAM authorization for cross-account data migration and synchronization](https://help.aliyun.com/document_detail/48468.html).
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
