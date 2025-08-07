// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGDNId(v string) *DescribeGlobalDatabaseNetworkRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type DescribeGlobalDatabaseNetworkRequest struct {
	// The ID of the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId        *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
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
}

func (s DescribeGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetGDNId(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
