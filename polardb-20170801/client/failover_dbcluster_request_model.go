// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FailoverDBClusterRequest
	GetClientToken() *string
	SetDBClusterId(v string) *FailoverDBClusterRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *FailoverDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FailoverDBClusterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FailoverDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FailoverDBClusterRequest
	GetResourceOwnerId() *int64
	SetRollBackForDisaster(v bool) *FailoverDBClusterRequest
	GetRollBackForDisaster() *bool
	SetTargetDBNodeId(v string) *FailoverDBClusterRequest
	GetTargetDBNodeId() *string
	SetTargetZoneType(v string) *FailoverDBClusterRequest
	GetTargetZoneType() *string
}

type FailoverDBClusterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must make sure that the value is unique among different requests. The token is case-sensitive and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
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
	// - true: Switch back to the original primary zone.
	//
	// - false: Do not switch back to the original primary zone.
	//
	// example:
	//
	// false
	RollBackForDisaster *bool `json:"RollBackForDisaster,omitempty" xml:"RollBackForDisaster,omitempty"`
	// The ID of the node that you want to promote to the primary node. If you leave this parameter empty, the system automatically elects a node. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query node IDs and other information.
	//
	// > - Automatic election of the primary node: The system identifies all read-only nodes that are eligible for election. The system selects one or more read-only nodes with the highest priority. If the switchover to the first node fails due to reasons such as network issues or abnormal replication status, the system attempts to switch to the next node until the switchover succeeds.
	//
	// >- This parameter is required when the cluster is a PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL cluster. This parameter is optional when the cluster is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-***********
	TargetDBNodeId *string `json:"TargetDBNodeId,omitempty" xml:"TargetDBNodeId,omitempty"`
	// Specifies whether the primary/secondary switchover is performed within the primary zone. Valid values:
	//
	// - Primary: The primary/secondary switchover is performed within the primary zone.
	//
	// - Standby: The switchover is performed to the hot standby storage cluster.
	//
	// example:
	//
	// Primary
	TargetZoneType *string `json:"TargetZoneType,omitempty" xml:"TargetZoneType,omitempty"`
}

func (s FailoverDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterRequest) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FailoverDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *FailoverDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FailoverDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FailoverDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FailoverDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FailoverDBClusterRequest) GetRollBackForDisaster() *bool {
	return s.RollBackForDisaster
}

func (s *FailoverDBClusterRequest) GetTargetDBNodeId() *string {
	return s.TargetDBNodeId
}

func (s *FailoverDBClusterRequest) GetTargetZoneType() *string {
	return s.TargetZoneType
}

func (s *FailoverDBClusterRequest) SetClientToken(v string) *FailoverDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDBClusterRequest) SetDBClusterId(v string) *FailoverDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetOwnerAccount(v string) *FailoverDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetOwnerId(v int64) *FailoverDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerAccount(v string) *FailoverDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerId(v int64) *FailoverDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetRollBackForDisaster(v bool) *FailoverDBClusterRequest {
	s.RollBackForDisaster = &v
	return s
}

func (s *FailoverDBClusterRequest) SetTargetDBNodeId(v string) *FailoverDBClusterRequest {
	s.TargetDBNodeId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetTargetZoneType(v string) *FailoverDBClusterRequest {
	s.TargetZoneType = &v
	return s
}

func (s *FailoverDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
