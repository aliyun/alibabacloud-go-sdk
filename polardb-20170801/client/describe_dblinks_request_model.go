// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBLinksRequest
	GetDBClusterId() *string
	SetDBLinkName(v string) *DescribeDBLinksRequest
	GetDBLinkName() *string
	SetOwnerAccount(v string) *DescribeDBLinksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBLinksRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBLinksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBLinksRequest
	GetResourceOwnerId() *int64
}

type DescribeDBLinksRequest struct {
	// The ID of the cluster for which you want to query DBLink information.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query the list of PolarDB clusters.
	//
	// > Use only the ID of a supported PolarDB for PostgreSQL (Compatible with Oracle) cluster. Unsupported engines or versions return InvalidEngine.Unsupported.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-a*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the DBLink. If you do not specify this parameter, all DBLinks are returned.
	//
	// example:
	//
	// dblink_test
	DBLinkName           *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBLinksRequest) GetDBLinkName() *string {
	return s.DBLinkName
}

func (s *DescribeDBLinksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBLinksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBLinksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBLinksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBLinksRequest) SetDBClusterId(v string) *DescribeDBLinksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBLinksRequest) SetDBLinkName(v string) *DescribeDBLinksRequest {
	s.DBLinkName = &v
	return s
}

func (s *DescribeDBLinksRequest) SetOwnerAccount(v string) *DescribeDBLinksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBLinksRequest) SetOwnerId(v int64) *DescribeDBLinksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBLinksRequest) SetResourceOwnerAccount(v string) *DescribeDBLinksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBLinksRequest) SetResourceOwnerId(v int64) *DescribeDBLinksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBLinksRequest) Validate() error {
	return dara.Validate(s)
}
