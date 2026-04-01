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
	SetDBInstanceId(v string) *DeleteDBNodesRequest
	GetDBInstanceId() *string
	SetDBNodeId(v []*string) *DeleteDBNodesRequest
	GetDBNodeId() []*string
	SetOwnerAccount(v string) *DeleteDBNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBNodesRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeleteDBNodesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBNodesRequest
	GetResourceOwnerId() *int64
}

type DeleteDBNodesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node IDs.
	//
	// This parameter is required.
	DBNodeId     []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group. You can call the DescribeDBInstanceAttribute operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DeleteDBNodesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBNodesRequest) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *DeleteDBNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *DeleteDBNodesRequest) SetDBInstanceId(v string) *DeleteDBNodesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBNodeId(v []*string) *DeleteDBNodesRequest {
	s.DBNodeId = v
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

func (s *DeleteDBNodesRequest) SetResourceGroupId(v string) *DeleteDBNodesRequest {
	s.ResourceGroupId = &v
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
