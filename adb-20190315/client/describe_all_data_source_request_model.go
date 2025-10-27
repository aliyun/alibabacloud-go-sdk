// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAllDataSourceRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAllDataSourceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAllDataSourceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAllDataSourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAllDataSourceRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeAllDataSourceRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeAllDataSourceRequest
	GetTableName() *string
}

type DescribeAllDataSourceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612241.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAllDataSourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAllDataSourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAllDataSourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAllDataSourceRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourceRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerAccount(v string) *DescribeAllDataSourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetResourceOwnerId(v int64) *DescribeAllDataSourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetSchemaName(v string) *DescribeAllDataSourceRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetTableName(v string) *DescribeAllDataSourceRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
