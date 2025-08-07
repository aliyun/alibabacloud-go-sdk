// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterAccessWhitelistRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterAccessWhitelistRequest struct {
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAccessWhitelistRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterAccessWhitelistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterAccessWhitelistRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterAccessWhitelistRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
