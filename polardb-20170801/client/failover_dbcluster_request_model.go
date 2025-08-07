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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster.
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
	// Specifies whether to switch back services to the original primary zone when the original primary zone recovers.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	RollBackForDisaster *bool `json:"RollBackForDisaster,omitempty" xml:"RollBackForDisaster,omitempty"`
	// The ID of the read-only node that you want to promote to the primary node. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query node information, such as node IDs.
	//
	// > 	- If you leave this parameter empty, the system selects one or more available read-only nodes that have the highest failover priority as candidate primary nodes. If the failover to the first read-only node fails due to network issues, abnormal replication status, or other reasons, the system attempts to fail over your applications to the next read-only node until the failover is successful.
	//
	// >	- This parameter is required for PolarDB for Oracle and PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// example:
	//
	// pi-***********
	TargetDBNodeId *string `json:"TargetDBNodeId,omitempty" xml:"TargetDBNodeId,omitempty"`
	// Whether it is a primary-standby switch within the primary availability zone, with the following values:
	//
	// Primary: Primary-standby switch within the primary availability zone.
	//
	// Standby: Switch to the storage hot backup cluster.
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
