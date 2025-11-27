// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceName(v string) *DescribeDBInstanceConnectivityRequest
	GetDbInstanceName() *string
	SetOwnerAccount(v string) *DescribeDBInstanceConnectivityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceConnectivityRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBInstanceConnectivityRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceConnectivityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceConnectivityRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDBInstanceConnectivityRequest
	GetSecurityToken() *string
	SetSourceIpAddress(v string) *DescribeDBInstanceConnectivityRequest
	GetSourceIpAddress() *string
}

type DescribeDBInstanceConnectivityRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4ns09hgoy99i5gez
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
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
	// 172.16.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
}

func (s DescribeDBInstanceConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConnectivityRequest) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstanceConnectivityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceConnectivityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceConnectivityRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceConnectivityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceConnectivityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceConnectivityRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDBInstanceConnectivityRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeDBInstanceConnectivityRequest) SetDbInstanceName(v string) *DescribeDBInstanceConnectivityRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetOwnerAccount(v string) *DescribeDBInstanceConnectivityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetOwnerId(v int64) *DescribeDBInstanceConnectivityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetResourceGroupId(v string) *DescribeDBInstanceConnectivityRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceConnectivityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceConnectivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetSecurityToken(v string) *DescribeDBInstanceConnectivityRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) SetSourceIpAddress(v string) *DescribeDBInstanceConnectivityRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeDBInstanceConnectivityRequest) Validate() error {
	return dara.Validate(s)
}
