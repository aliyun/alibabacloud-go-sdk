// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeRoleTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateNodeRoleTagRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CreateNodeRoleTagRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNodeRoleTagRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateNodeRoleTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNodeRoleTagRequest
	GetResourceOwnerId() *int64
	SetShardList(v string) *CreateNodeRoleTagRequest
	GetShardList() *string
}

type CreateNodeRoleTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dds-2ze09d7be1fxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// d-l5xf98b78b9fxxxx,d-l5xf98b7sf9fxxxx
	ShardList *string `json:"ShardList,omitempty" xml:"ShardList,omitempty"`
}

func (s CreateNodeRoleTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeRoleTagRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRoleTagRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateNodeRoleTagRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNodeRoleTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNodeRoleTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNodeRoleTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNodeRoleTagRequest) GetShardList() *string {
	return s.ShardList
}

func (s *CreateNodeRoleTagRequest) SetDBInstanceId(v string) *CreateNodeRoleTagRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNodeRoleTagRequest) SetOwnerAccount(v string) *CreateNodeRoleTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNodeRoleTagRequest) SetOwnerId(v int64) *CreateNodeRoleTagRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNodeRoleTagRequest) SetResourceOwnerAccount(v string) *CreateNodeRoleTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNodeRoleTagRequest) SetResourceOwnerId(v int64) *CreateNodeRoleTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNodeRoleTagRequest) SetShardList(v string) *CreateNodeRoleTagRequest {
	s.ShardList = &v
	return s
}

func (s *CreateNodeRoleTagRequest) Validate() error {
	return dara.Validate(s)
}
