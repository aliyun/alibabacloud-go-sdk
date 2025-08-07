// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterConnectivityRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterConnectivityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterConnectivityRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBClusterConnectivityRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterConnectivityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterConnectivityRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBClusterConnectivityRequest
	GetSecurityToken() *string
	SetSourceIpAddress(v string) *DescribeDBClusterConnectivityRequest
	GetSourceIpAddress() *string
}

type DescribeDBClusterConnectivityRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxxxxxxx
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
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The source IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.***.***.1
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
}

func (s DescribeDBClusterConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConnectivityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConnectivityRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterConnectivityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterConnectivityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterConnectivityRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClusterConnectivityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterConnectivityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterConnectivityRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBClusterConnectivityRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeDBClusterConnectivityRequest) SetDBClusterId(v string) *DescribeDBClusterConnectivityRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetOwnerAccount(v string) *DescribeDBClusterConnectivityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetOwnerId(v int64) *DescribeDBClusterConnectivityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetResourceGroupId(v string) *DescribeDBClusterConnectivityRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterConnectivityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetResourceOwnerId(v int64) *DescribeDBClusterConnectivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetSecurityToken(v string) *DescribeDBClusterConnectivityRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) SetSourceIpAddress(v string) *DescribeDBClusterConnectivityRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeDBClusterConnectivityRequest) Validate() error {
	return dara.Validate(s)
}
