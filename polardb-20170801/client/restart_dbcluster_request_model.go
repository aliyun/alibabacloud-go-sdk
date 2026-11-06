// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartDBClusterRequest
	GetClientToken() *string
	SetDBClusterId(v string) *RestartDBClusterRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RestartDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBClusterRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *RestartDBClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RestartDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBClusterRequest
	GetResourceOwnerId() *int64
}

type RestartDBClusterRequest struct {
	// The client token that is used to ensure the idempotence of the request. Use different values for different requests.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The PolarDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartDBClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestartDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBClusterRequest) SetClientToken(v string) *RestartDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBClusterRequest) SetDBClusterId(v string) *RestartDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBClusterRequest) SetOwnerAccount(v string) *RestartDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBClusterRequest) SetOwnerId(v int64) *RestartDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceGroupId(v string) *RestartDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceOwnerAccount(v string) *RestartDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceOwnerId(v int64) *RestartDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
