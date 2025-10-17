// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterVersionZonalRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeDBClusterVersionZonalRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterVersionZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterVersionZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterVersionZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterVersionZonalRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterVersionZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// LATEST_VERSION
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterVersionZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterVersionZonalRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterVersionZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterVersionZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterVersionZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterVersionZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterVersionZonalRequest) SetDBClusterId(v string) *DescribeDBClusterVersionZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) SetDescribeType(v string) *DescribeDBClusterVersionZonalRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) SetOwnerAccount(v string) *DescribeDBClusterVersionZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) SetOwnerId(v int64) *DescribeDBClusterVersionZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterVersionZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) SetResourceOwnerId(v int64) *DescribeDBClusterVersionZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionZonalRequest) Validate() error {
	return dara.Validate(s)
}
