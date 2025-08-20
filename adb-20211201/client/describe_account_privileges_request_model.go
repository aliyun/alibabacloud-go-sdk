// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountPrivilegesRequest
	GetAccountName() *string
	SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegesRequest
	GetColumnPrivilegeObject() *string
	SetDBClusterId(v string) *DescribeAccountPrivilegesRequest
	GetDBClusterId() *string
	SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegesRequest
	GetDatabasePrivilegeObject() *string
	SetPageNumber(v string) *DescribeAccountPrivilegesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAccountPrivilegesRequest
	GetPageSize() *string
	SetPrivilegeType(v string) *DescribeAccountPrivilegesRequest
	GetPrivilegeType() *string
	SetRegionId(v string) *DescribeAccountPrivilegesRequest
	GetRegionId() *string
	SetTablePrivilegeObject(v string) *DescribeAccountPrivilegesRequest
	GetTablePrivilegeObject() *string
}

type DescribeAccountPrivilegesRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The columns that you want to query. You can use this parameter to query the permissions of the database account on specific columns. This parameter is available only if the PrivilegeType parameter is set to Column.
	//
	// example:
	//
	// col1
	ColumnPrivilegeObject *string `json:"ColumnPrivilegeObject,omitempty" xml:"ColumnPrivilegeObject,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The databases that you want to query. You can use this parameter to query the permissions of the database account on specific databases. This parameter is available only if the PrivilegeType parameter is set to Database, Table, or Column.
	//
	// example:
	//
	// db1
	DatabasePrivilegeObject *string `json:"DatabasePrivilegeObject,omitempty" xml:"DatabasePrivilegeObject,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The permission level that you want to query. You can call the `DescribeEnabledPrivileges` operation to query the permission level of the database account.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tables that you want to query. You can use this parameter to query the permissions of the database account on specific tables. This parameter can be used together with the DatabasePrivilegeObject parameter. This parameter is available only if the PrivilegeType parameter is set to Table or Column.
	//
	// example:
	//
	// table1
	TablePrivilegeObject *string `json:"TablePrivilegeObject,omitempty" xml:"TablePrivilegeObject,omitempty"`
}

func (s DescribeAccountPrivilegesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountPrivilegesRequest) GetColumnPrivilegeObject() *string {
	return s.ColumnPrivilegeObject
}

func (s *DescribeAccountPrivilegesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountPrivilegesRequest) GetDatabasePrivilegeObject() *string {
	return s.DatabasePrivilegeObject
}

func (s *DescribeAccountPrivilegesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAccountPrivilegesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAccountPrivilegesRequest) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *DescribeAccountPrivilegesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountPrivilegesRequest) GetTablePrivilegeObject() *string {
	return s.TablePrivilegeObject
}

func (s *DescribeAccountPrivilegesRequest) SetAccountName(v string) *DescribeAccountPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.ColumnPrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetDBClusterId(v string) *DescribeAccountPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.DatabasePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPageNumber(v string) *DescribeAccountPrivilegesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPageSize(v string) *DescribeAccountPrivilegesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetPrivilegeType(v string) *DescribeAccountPrivilegesRequest {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetRegionId(v string) *DescribeAccountPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) SetTablePrivilegeObject(v string) *DescribeAccountPrivilegesRequest {
	s.TablePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegesRequest) Validate() error {
	return dara.Validate(s)
}
