// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeColumnsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeColumnsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeColumnsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeColumnsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeColumnsRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeColumnsRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeColumnsRequest
	GetTableName() *string
}

type DescribeColumnsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-2zeux3ua25242****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name. You can call the [DescribeSchemas](https://help.aliyun.com/document_detail/350931.html) operation to query database names.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name. You can call the [DescribeTables](https://help.aliyun.com/document_detail/350932.html) operation to query table names.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeColumnsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeColumnsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeColumnsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeColumnsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeColumnsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerAccount(v string) *DescribeColumnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetOwnerId(v int64) *DescribeColumnsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerAccount(v string) *DescribeColumnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColumnsRequest) SetResourceOwnerId(v int64) *DescribeColumnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColumnsRequest) SetSchemaName(v string) *DescribeColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsRequest) Validate() error {
	return dara.Validate(s)
}
