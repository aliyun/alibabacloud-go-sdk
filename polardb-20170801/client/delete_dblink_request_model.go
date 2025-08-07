// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBLinkRequest
	GetDBClusterId() *string
	SetDBLinkName(v string) *DeleteDBLinkRequest
	GetDBLinkName() *string
	SetOwnerAccount(v string) *DeleteDBLinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBLinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBLinkRequest
	GetResourceOwnerId() *int64
}

type DeleteDBLinkRequest struct {
	// The ID of the source cluster where a database link is to be deleted.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query PolarDB clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database link to be deleted.
	//
	// This parameter is required.
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

func (s DeleteDBLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBLinkRequest) GetDBLinkName() *string {
	return s.DBLinkName
}

func (s *DeleteDBLinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBLinkRequest) SetDBClusterId(v string) *DeleteDBLinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBLinkRequest) SetDBLinkName(v string) *DeleteDBLinkRequest {
	s.DBLinkName = &v
	return s
}

func (s *DeleteDBLinkRequest) SetOwnerAccount(v string) *DeleteDBLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBLinkRequest) SetOwnerId(v int64) *DeleteDBLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBLinkRequest) SetResourceOwnerAccount(v string) *DeleteDBLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBLinkRequest) SetResourceOwnerId(v int64) *DeleteDBLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBLinkRequest) Validate() error {
	return dara.Validate(s)
}
