// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateClusterPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AllocateClusterPublicConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateClusterPublicConnectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocateClusterPublicConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateClusterPublicConnectionRequest
	GetResourceOwnerId() *int64
}

type AllocateClusterPublicConnectionRequest struct {
	// The prefix of the public endpoint.
	//
	// 	- The prefix must contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter.
	//
	// 	- The prefix can be up to 30 characters in length.
	//
	// 	- By default, the cluster name is used as the prefix of the public endpoint.
	//
	// example:
	//
	// am-bp278jg9****.ads.aliyuncs.com
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The Data Warehouse Edition Cluster ID of AnalyticDB for MySQL. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to retrieve all Data Warehouse Edition cluster IDs within the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp278jg9****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateClusterPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateClusterPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateClusterPublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateClusterPublicConnectionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AllocateClusterPublicConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateClusterPublicConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateClusterPublicConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateClusterPublicConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateClusterPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateClusterPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetDBClusterId(v string) *AllocateClusterPublicConnectionRequest {
	s.DBClusterId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateClusterPublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateClusterPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}
