// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBNodesRequest
	GetClientToken() *string
	SetCloudProvider(v string) *DeleteDBNodesRequest
	GetCloudProvider() *string
	SetDBClusterId(v string) *DeleteDBNodesRequest
	GetDBClusterId() *string
	SetDBNodeId(v []*string) *DeleteDBNodesRequest
	GetDBNodeId() []*string
	SetDBNodeType(v string) *DeleteDBNodesRequest
	GetDBNodeType() *string
	SetOwnerAccount(v string) *DeleteDBNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBNodesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBNodesRequest
	GetResourceOwnerId() *int64
}

type DeleteDBNodesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a42***********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The IDs of the nodes.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/185342.html) operation to query the details of all clusters that belong to your Alibaba Cloud account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-************
	DBNodeId []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
	// The node type. Valid values:
	//
	// 	- RO
	//
	// 	- STANDBY
	//
	// 	- DLNode
	//
	// Enumerated values:
	//
	// 	- DLNode: AI node
	//
	// 	- STANDBY: standby node
	//
	// 	- RO: read-only node
	//
	// example:
	//
	// RO
	DBNodeType           *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBNodesRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DeleteDBNodesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBNodesRequest) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *DeleteDBNodesRequest) GetDBNodeType() *string {
	return s.DBNodeType
}

func (s *DeleteDBNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBNodesRequest) SetClientToken(v string) *DeleteDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBNodesRequest) SetCloudProvider(v string) *DeleteDBNodesRequest {
	s.CloudProvider = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBClusterId(v string) *DeleteDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBNodeId(v []*string) *DeleteDBNodesRequest {
	s.DBNodeId = v
	return s
}

func (s *DeleteDBNodesRequest) SetDBNodeType(v string) *DeleteDBNodesRequest {
	s.DBNodeType = &v
	return s
}

func (s *DeleteDBNodesRequest) SetOwnerAccount(v string) *DeleteDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetOwnerId(v int64) *DeleteDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerAccount(v string) *DeleteDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerId(v int64) *DeleteDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) Validate() error {
	return dara.Validate(s)
}
