// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteNodeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteNodeRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *DeleteNodeRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DeleteNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNodeRequest
	GetResourceOwnerId() *int64
}

type DeleteNodeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard or mongos node that you want to delete. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/61923.html) operation to query node IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNodeRequest) SetClientToken(v string) *DeleteNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNodeRequest) SetDBInstanceId(v string) *DeleteNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteNodeRequest) SetNodeId(v string) *DeleteNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteNodeRequest) SetOwnerAccount(v string) *DeleteNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNodeRequest) SetOwnerId(v int64) *DeleteNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNodeRequest) SetResourceOwnerAccount(v string) *DeleteNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNodeRequest) SetResourceOwnerId(v int64) *DeleteNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNodeRequest) Validate() error {
	return dara.Validate(s)
}
