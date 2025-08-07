// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeHotReplicaModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeHotReplicaModeRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *ModifyDBNodeHotReplicaModeRequest
	GetDBNodeId() *string
	SetHotReplicaMode(v string) *ModifyDBNodeHotReplicaModeRequest
	GetHotReplicaMode() *string
	SetOwnerAccount(v string) *ModifyDBNodeHotReplicaModeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeHotReplicaModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBNodeHotReplicaModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeHotReplicaModeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodeHotReplicaModeRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-2vc327c2a14a3u858
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the node in the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-2ze28275h9x5r4wt1
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// Specifies whether to enable the hot standby feature. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	HotReplicaMode       *string `json:"HotReplicaMode,omitempty" xml:"HotReplicaMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBNodeHotReplicaModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeHotReplicaModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetHotReplicaMode() *string {
	return s.HotReplicaMode
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeHotReplicaModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetDBClusterId(v string) *ModifyDBNodeHotReplicaModeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetDBNodeId(v string) *ModifyDBNodeHotReplicaModeRequest {
	s.DBNodeId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetHotReplicaMode(v string) *ModifyDBNodeHotReplicaModeRequest {
	s.HotReplicaMode = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetOwnerAccount(v string) *ModifyDBNodeHotReplicaModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetOwnerId(v int64) *ModifyDBNodeHotReplicaModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeHotReplicaModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) SetResourceOwnerId(v int64) *ModifyDBNodeHotReplicaModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeHotReplicaModeRequest) Validate() error {
	return dara.Validate(s)
}
