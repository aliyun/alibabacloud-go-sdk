// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTablesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTablesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTablesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTablesRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeTablesRequest
	GetSchemaName() *string
}

type DescribeTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
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
}

func (s DescribeTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTablesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerAccount(v string) *DescribeTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerId(v int64) *DescribeTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerAccount(v string) *DescribeTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerId(v int64) *DescribeTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesRequest) Validate() error {
	return dara.Validate(s)
}
