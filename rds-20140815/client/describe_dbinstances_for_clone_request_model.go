// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesForCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstancesForCloneRequest
	GetClientToken() *string
	SetConnectionMode(v string) *DescribeDBInstancesForCloneRequest
	GetConnectionMode() *string
	SetCurrentInstanceId(v string) *DescribeDBInstancesForCloneRequest
	GetCurrentInstanceId() *string
	SetDBInstanceClass(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceId() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceStatus() *string
	SetDBInstanceType(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceType() *string
	SetEngine(v string) *DescribeDBInstancesForCloneRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBInstancesForCloneRequest
	GetEngineVersion() *string
	SetExpired(v string) *DescribeDBInstancesForCloneRequest
	GetExpired() *string
	SetInstanceNetworkType(v string) *DescribeDBInstancesForCloneRequest
	GetInstanceNetworkType() *string
	SetNodeType(v string) *DescribeDBInstancesForCloneRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *DescribeDBInstancesForCloneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesForCloneRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesForCloneRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesForCloneRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeDBInstancesForCloneRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeDBInstancesForCloneRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesForCloneRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesForCloneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesForCloneRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeDBInstancesForCloneRequest
	GetSearchKey() *string
	SetVSwitchId(v string) *DescribeDBInstancesForCloneRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeDBInstancesForCloneRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstancesForCloneRequest
	GetZoneId() *string
	SetProxyId(v string) *DescribeDBInstancesForCloneRequest
	GetProxyId() *string
}

type DescribeDBInstancesForCloneRequest struct {
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionMode      *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	CurrentInstanceId   *string `json:"CurrentInstanceId,omitempty" xml:"CurrentInstanceId,omitempty"`
	DBInstanceClass     *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus    *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceType      *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	Engine              *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	Expired             *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	NodeType            *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PayType             *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchKey            *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProxyId              *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstancesForCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstancesForCloneRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesForCloneRequest) GetCurrentInstanceId() *string {
	return s.CurrentInstanceId
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesForCloneRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesForCloneRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesForCloneRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstancesForCloneRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesForCloneRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBInstancesForCloneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesForCloneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesForCloneRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesForCloneRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesForCloneRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesForCloneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesForCloneRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeDBInstancesForCloneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesForCloneRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesForCloneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesForCloneRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstancesForCloneRequest) SetClientToken(v string) *DescribeDBInstancesForCloneRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetConnectionMode(v string) *DescribeDBInstancesForCloneRequest {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetCurrentInstanceId(v string) *DescribeDBInstancesForCloneRequest {
	s.CurrentInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceClass(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceId(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceType(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetEngine(v string) *DescribeDBInstancesForCloneRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetEngineVersion(v string) *DescribeDBInstancesForCloneRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetExpired(v string) *DescribeDBInstancesForCloneRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesForCloneRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetNodeType(v string) *DescribeDBInstancesForCloneRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetOwnerAccount(v string) *DescribeDBInstancesForCloneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetOwnerId(v int64) *DescribeDBInstancesForCloneRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPageNumber(v int32) *DescribeDBInstancesForCloneRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPageSize(v int32) *DescribeDBInstancesForCloneRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPayType(v string) *DescribeDBInstancesForCloneRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetRegionId(v string) *DescribeDBInstancesForCloneRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceGroupId(v string) *DescribeDBInstancesForCloneRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesForCloneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesForCloneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetSearchKey(v string) *DescribeDBInstancesForCloneRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetVSwitchId(v string) *DescribeDBInstancesForCloneRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetVpcId(v string) *DescribeDBInstancesForCloneRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetZoneId(v string) *DescribeDBInstancesForCloneRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetProxyId(v string) *DescribeDBInstancesForCloneRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) Validate() error {
	return dara.Validate(s)
}
