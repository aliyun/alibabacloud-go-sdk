// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterAttributeRequest
	GetDBClusterId() *string
	SetDescribeType(v string) *DescribeDBClusterAttributeRequest
	GetDescribeType() *string
	SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterAttributeRequest struct {
	// Cluster ID.
	//
	// > You can view detailed information about all clusters under your account, including the cluster ID, through the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Whether to obtain information about AI-related nodes.
	//
	// example:
	//
	// AI
	DescribeType         *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAttributeRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClusterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetDescribeType(v string) *DescribeDBClusterAttributeRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
