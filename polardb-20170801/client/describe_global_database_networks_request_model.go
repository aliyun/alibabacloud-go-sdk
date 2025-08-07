// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeGlobalDatabaseNetworksRequest
	GetDBClusterId() *string
	SetFilterRegion(v string) *DescribeGlobalDatabaseNetworksRequest
	GetFilterRegion() *string
	SetGDNDescription(v string) *DescribeGlobalDatabaseNetworksRequest
	GetGDNDescription() *string
	SetGDNId(v string) *DescribeGlobalDatabaseNetworksRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGlobalDatabaseNetworksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGlobalDatabaseNetworksRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeGlobalDatabaseNetworksRequest
	GetSecurityToken() *string
}

type DescribeGlobalDatabaseNetworksRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specify the region in which you want to query GDNs. You can create secondary clusters for the GDNs.
	//
	// example:
	//
	// cn-beijing
	FilterRegion *string `json:"FilterRegion,omitempty" xml:"FilterRegion,omitempty"`
	// The description of the GDN. The description must meet the following requirements:
	//
	// 	- It cannot start with `http://` or `https://`.
	//
	// 	- It must start with a letter.
	//
	// 	- It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- It must be 2 to 126 characters in length.
	//
	// example:
	//
	// test
	GDNDescription *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	// The ID of the GDN.
	//
	// example:
	//
	// gdn-****************
	GDNId        *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: 1. The value must be an integer that is greater than 0.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s DescribeGlobalDatabaseNetworksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetFilterRegion() *string {
	return s.FilterRegion
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetGDNDescription() *string {
	return s.GDNDescription
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalDatabaseNetworksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetFilterRegion(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.FilterRegion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetGDNDescription(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.GDNDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetGDNId(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetPageNumber(v int32) *DescribeGlobalDatabaseNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetPageSize(v int32) *DescribeGlobalDatabaseNetworksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetSecurityToken(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) Validate() error {
	return dara.Validate(s)
}
