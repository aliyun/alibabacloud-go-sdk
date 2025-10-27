// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateTableSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetCreateTableSQLRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *GetCreateTableSQLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetCreateTableSQLRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetCreateTableSQLRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetCreateTableSQLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCreateTableSQLRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *GetCreateTableSQLRequest
	GetSchemaName() *string
	SetTableName(v string) *GetCreateTableSQLRequest
	GetTableName() *string
}

type GetCreateTableSQLRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6wjk5xxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetCreateTableSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCreateTableSQLRequest) GoString() string {
	return s.String()
}

func (s *GetCreateTableSQLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetCreateTableSQLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetCreateTableSQLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCreateTableSQLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCreateTableSQLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCreateTableSQLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCreateTableSQLRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetCreateTableSQLRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetCreateTableSQLRequest) SetDBClusterId(v string) *GetCreateTableSQLRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetOwnerAccount(v string) *GetCreateTableSQLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetOwnerId(v int64) *GetCreateTableSQLRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetRegionId(v string) *GetCreateTableSQLRequest {
	s.RegionId = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetResourceOwnerAccount(v string) *GetCreateTableSQLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetResourceOwnerId(v int64) *GetCreateTableSQLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetSchemaName(v string) *GetCreateTableSQLRequest {
	s.SchemaName = &v
	return s
}

func (s *GetCreateTableSQLRequest) SetTableName(v string) *GetCreateTableSQLRequest {
	s.TableName = &v
	return s
}

func (s *GetCreateTableSQLRequest) Validate() error {
	return dara.Validate(s)
}
