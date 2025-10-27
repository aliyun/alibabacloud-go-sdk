// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTableDetailRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeTableDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTableDetailRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTableDetailRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTableDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTableDetailRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeTableDetailRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeTableDetailRequest
	GetTableName() *string
}

type DescribeTableDetailRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s DescribeTableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTableDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTableDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTableDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTableDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTableDetailRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTableDetailRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableDetailRequest) SetDBClusterId(v string) *DescribeTableDetailRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerAccount(v string) *DescribeTableDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetOwnerId(v int64) *DescribeTableDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetRegionId(v string) *DescribeTableDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerAccount(v string) *DescribeTableDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTableDetailRequest) SetResourceOwnerId(v int64) *DescribeTableDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetSchemaName(v string) *DescribeTableDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetTableName(v string) *DescribeTableDetailRequest {
	s.TableName = &v
	return s
}

func (s *DescribeTableDetailRequest) Validate() error {
	return dara.Validate(s)
}
