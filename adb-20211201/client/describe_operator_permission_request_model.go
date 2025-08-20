// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeOperatorPermissionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeOperatorPermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeOperatorPermissionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest
	GetResourceOwnerId() *int64
}

type DescribeOperatorPermissionRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6wjk5xxxxxxxxxx
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeOperatorPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeOperatorPermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeOperatorPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeOperatorPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeOperatorPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeOperatorPermissionRequest) SetDBClusterId(v string) *DescribeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) SetResourceOwnerId(v int64) *DescribeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeOperatorPermissionRequest) Validate() error {
	return dara.Validate(s)
}
