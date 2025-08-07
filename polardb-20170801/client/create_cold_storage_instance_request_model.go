// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateColdStorageInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateColdStorageInstanceRequest
	GetClientToken() *string
	SetColdStorageInstanceDescription(v string) *CreateColdStorageInstanceRequest
	GetColdStorageInstanceDescription() *string
	SetDBClusterId(v string) *CreateColdStorageInstanceRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateColdStorageInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateColdStorageInstanceRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateColdStorageInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateColdStorageInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateColdStorageInstanceRequest
	GetResourceOwnerId() *int64
}

type CreateColdStorageInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the cluster. The description cannot exceed 256 characters in length.
	//
	// example:
	//
	// xxxxxxxxx
	ColdStorageInstanceDescription *string `json:"ColdStorageInstanceDescription,omitempty" xml:"ColdStorageInstanceDescription,omitempty"`
	// The cluster ID. > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters within your account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-wz9ye3xrtj60ua6d1
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateColdStorageInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateColdStorageInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateColdStorageInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateColdStorageInstanceRequest) GetColdStorageInstanceDescription() *string {
	return s.ColdStorageInstanceDescription
}

func (s *CreateColdStorageInstanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateColdStorageInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateColdStorageInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateColdStorageInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateColdStorageInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateColdStorageInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateColdStorageInstanceRequest) SetClientToken(v string) *CreateColdStorageInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetColdStorageInstanceDescription(v string) *CreateColdStorageInstanceRequest {
	s.ColdStorageInstanceDescription = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetDBClusterId(v string) *CreateColdStorageInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetOwnerAccount(v string) *CreateColdStorageInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetOwnerId(v int64) *CreateColdStorageInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetResourceGroupId(v string) *CreateColdStorageInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetResourceOwnerAccount(v string) *CreateColdStorageInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) SetResourceOwnerId(v int64) *CreateColdStorageInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateColdStorageInstanceRequest) Validate() error {
	return dara.Validate(s)
}
