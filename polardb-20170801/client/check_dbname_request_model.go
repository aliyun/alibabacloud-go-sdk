// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckDBNameRequest
	GetDBClusterId() *string
	SetDBName(v string) *CheckDBNameRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CheckDBNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckDBNameRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckDBNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckDBNameRequest
	GetResourceOwnerId() *int64
}

type CheckDBNameRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a region, such as the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckDBNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDBNameRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckDBNameRequest) GetDBName() *string {
	return s.DBName
}

func (s *CheckDBNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckDBNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckDBNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckDBNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckDBNameRequest) SetDBClusterId(v string) *CheckDBNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckDBNameRequest) SetDBName(v string) *CheckDBNameRequest {
	s.DBName = &v
	return s
}

func (s *CheckDBNameRequest) SetOwnerAccount(v string) *CheckDBNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckDBNameRequest) SetOwnerId(v int64) *CheckDBNameRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDBNameRequest) SetResourceOwnerAccount(v string) *CheckDBNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDBNameRequest) SetResourceOwnerId(v int64) *CheckDBNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDBNameRequest) Validate() error {
	return dara.Validate(s)
}
