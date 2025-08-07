// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterServerlessConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterServerlessConfRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterServerlessConfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterServerlessConfRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterServerlessConfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterServerlessConfRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterServerlessConfRequest struct {
	// Serverless cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterServerlessConfRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterServerlessConfRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterServerlessConfRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterServerlessConfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterServerlessConfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterServerlessConfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterServerlessConfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterServerlessConfRequest) SetDBClusterId(v string) *DescribeDBClusterServerlessConfRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterServerlessConfRequest) SetOwnerAccount(v string) *DescribeDBClusterServerlessConfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterServerlessConfRequest) SetOwnerId(v int64) *DescribeDBClusterServerlessConfRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterServerlessConfRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterServerlessConfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterServerlessConfRequest) SetResourceOwnerId(v int64) *DescribeDBClusterServerlessConfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterServerlessConfRequest) Validate() error {
	return dara.Validate(s)
}
