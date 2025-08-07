// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterTDERequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterTDERequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterTDERequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterTDERequest struct {
	// The ID of the cluster.
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

func (s DescribeDBClusterTDERequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterTDERequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDERequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterTDERequest) SetDBClusterId(v string) *DescribeDBClusterTDERequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetOwnerAccount(v string) *DescribeDBClusterTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetOwnerId(v int64) *DescribeDBClusterTDERequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetResourceOwnerAccount(v string) *DescribeDBClusterTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetResourceOwnerId(v int64) *DescribeDBClusterTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterTDERequest) Validate() error {
	return dara.Validate(s)
}
