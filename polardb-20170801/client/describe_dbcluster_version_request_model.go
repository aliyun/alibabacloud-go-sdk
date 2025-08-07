// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterVersionRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeDBClusterVersionRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterVersionRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterVersionRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to query the information about the latest versions or the versions to which the cluster can be updated. Valid values:
	//
	// 	- LATEST_VERSION: the information about the latest versions.
	//
	// 	- AVAILABLE_VERSION: the information about the versions to which the cluster can be updated.
	//
	// example:
	//
	// LATEST_VERSION
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterVersionRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterVersionRequest) SetDBClusterId(v string) *DescribeDBClusterVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetDescribeType(v string) *DescribeDBClusterVersionRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetOwnerAccount(v string) *DescribeDBClusterVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetOwnerId(v int64) *DescribeDBClusterVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetResourceOwnerId(v int64) *DescribeDBClusterVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) Validate() error {
	return dara.Validate(s)
}
