// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterMigrationRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterMigrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterMigrationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterMigrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterMigrationRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterMigrationRequest struct {
	// The network type of the endpoint. Valid values:
	//
	// 	- **Public**: the public endpoint
	//
	// 	- **Private**: the internal endpoint (VPC)
	//
	// 	- **Inner**: the internal endpoint (classic network)
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterMigrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterMigrationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterMigrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterMigrationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterMigrationRequest) SetDBClusterId(v string) *DescribeDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) Validate() error {
	return dara.Validate(s)
}
