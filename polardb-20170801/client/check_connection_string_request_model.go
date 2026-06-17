// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *CheckConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *CheckConnectionStringRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckConnectionStringRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckConnectionStringRequest
	GetResourceOwnerId() *int64
}

type CheckConnectionStringRequest struct {
	// The prefix of the new connection string. The prefix must meet the following requirements:
	//
	// - It must consist of lowercase letters, digits, and periods (.).
	//
	// - It must start with a letter and end with a letter or a digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************.pg.polardb.pre.rds.aliyuncs.com:5432
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the cluster.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to view the details of all clusters in your account, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *CheckConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CheckConnectionStringRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckConnectionStringRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckConnectionStringRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckConnectionStringRequest) SetConnectionStringPrefix(v string) *CheckConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CheckConnectionStringRequest) SetDBClusterId(v string) *CheckConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckConnectionStringRequest) SetOwnerAccount(v string) *CheckConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckConnectionStringRequest) SetOwnerId(v int64) *CheckConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckConnectionStringRequest) SetResourceOwnerAccount(v string) *CheckConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckConnectionStringRequest) SetResourceOwnerId(v int64) *CheckConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
