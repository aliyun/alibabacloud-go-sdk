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
