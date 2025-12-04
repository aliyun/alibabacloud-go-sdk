// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAllDataSourcesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAllDataSourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAllDataSourcesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAllDataSourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAllDataSourcesRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeAllDataSourcesRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeAllDataSourcesRequest
	GetTableName() *string
}

type DescribeAllDataSourcesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAllDataSourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAllDataSourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAllDataSourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAllDataSourcesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourcesRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourcesRequest) SetDBClusterId(v string) *DescribeAllDataSourcesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetOwnerAccount(v string) *DescribeAllDataSourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetOwnerId(v int64) *DescribeAllDataSourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetResourceOwnerAccount(v string) *DescribeAllDataSourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetResourceOwnerId(v int64) *DescribeAllDataSourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetSchemaName(v string) *DescribeAllDataSourcesRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) SetTableName(v string) *DescribeAllDataSourcesRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
