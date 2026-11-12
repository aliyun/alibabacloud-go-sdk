// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FailoverDBClusterZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *FailoverDBClusterZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *FailoverDBClusterZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FailoverDBClusterZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FailoverDBClusterZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FailoverDBClusterZonalRequest
	GetResourceOwnerId() *int64
	SetRollBackForDisaster(v bool) *FailoverDBClusterZonalRequest
	GetRollBackForDisaster() *bool
	SetTargetDBNodeId(v string) *FailoverDBClusterZonalRequest
	GetTargetDBNodeId() *string
	SetTargetZoneType(v string) *FailoverDBClusterZonalRequest
	GetTargetZoneType() *string
}

type FailoverDBClusterZonalRequest struct {
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and is case-sensitive. It can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the edge cloud cluster. Only edge cluster IDs returned by DescribeDBClustersZonal are accepted. Standard PolarDB cluster IDs are not supported. Call description: Only edge cluster IDs are accepted. Call DescribeDBClustersZonal to obtain available edge cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to switch back to the original primary zone after a failover. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	RollBackForDisaster *bool `json:"RollBackForDisaster,omitempty" xml:"RollBackForDisaster,omitempty"`
	// The ID of the edge cluster node to be switched to the primary node. If you do not specify this parameter, the system automatically elects a primary node. You can call DescribeDbClusterAttributeZonal to query the node ID. Call description: DescribeDbClusterAttributeZonal.
	//
	// example:
	//
	// pi-***********
	TargetDBNodeId *string `json:"TargetDBNodeId,omitempty" xml:"TargetDBNodeId,omitempty"`
	// Specifies whether to perform a primary/secondary switchover within the primary zone. Valid values:
	//
	// - Primary: Performs a primary/secondary switchover within the primary zone.
	//
	// - Standby: Switches to the hot standby storage cluster.
	//
	// example:
	//
	// Primary
	TargetZoneType *string `json:"TargetZoneType,omitempty" xml:"TargetZoneType,omitempty"`
}

func (s FailoverDBClusterZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterZonalRequest) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FailoverDBClusterZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *FailoverDBClusterZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FailoverDBClusterZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FailoverDBClusterZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FailoverDBClusterZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FailoverDBClusterZonalRequest) GetRollBackForDisaster() *bool {
	return s.RollBackForDisaster
}

func (s *FailoverDBClusterZonalRequest) GetTargetDBNodeId() *string {
	return s.TargetDBNodeId
}

func (s *FailoverDBClusterZonalRequest) GetTargetZoneType() *string {
	return s.TargetZoneType
}

func (s *FailoverDBClusterZonalRequest) SetClientToken(v string) *FailoverDBClusterZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetDBClusterId(v string) *FailoverDBClusterZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetOwnerAccount(v string) *FailoverDBClusterZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetOwnerId(v int64) *FailoverDBClusterZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetResourceOwnerAccount(v string) *FailoverDBClusterZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetResourceOwnerId(v int64) *FailoverDBClusterZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetRollBackForDisaster(v bool) *FailoverDBClusterZonalRequest {
	s.RollBackForDisaster = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetTargetDBNodeId(v string) *FailoverDBClusterZonalRequest {
	s.TargetDBNodeId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetTargetZoneType(v string) *FailoverDBClusterZonalRequest {
	s.TargetZoneType = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) Validate() error {
	return dara.Validate(s)
}
