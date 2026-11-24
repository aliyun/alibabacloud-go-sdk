// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegeObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountPrivilegeObjectsRequest
	GetAccountName() *string
	SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest
	GetColumnPrivilegeObject() *string
	SetDBClusterId(v string) *DescribeAccountPrivilegeObjectsRequest
	GetDBClusterId() *string
	SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest
	GetDatabasePrivilegeObject() *string
	SetPageNumber(v string) *DescribeAccountPrivilegeObjectsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAccountPrivilegeObjectsRequest
	GetPageSize() *string
	SetPrivilegeType(v string) *DescribeAccountPrivilegeObjectsRequest
	GetPrivilegeType() *string
	SetRegionId(v string) *DescribeAccountPrivilegeObjectsRequest
	GetRegionId() *string
	SetTablePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest
	GetTablePrivilegeObject() *string
}

type DescribeAccountPrivilegeObjectsRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The column name used to filter the results.
	//
	// example:
	//
	// col1
	ColumnPrivilegeObject *string `json:"ColumnPrivilegeObject,omitempty" xml:"ColumnPrivilegeObject,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k3wdmt139****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name used to filter the results.
	//
	// example:
	//
	// database1
	DatabasePrivilegeObject *string `json:"DatabasePrivilegeObject,omitempty" xml:"DatabasePrivilegeObject,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The privilege level. Only the Database, Table, and Column levels are supported. The Global level is not supported.
	//
	// example:
	//
	// Column
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ch-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table name used to filter the results.
	//
	// example:
	//
	// table1
	TablePrivilegeObject *string `json:"TablePrivilegeObject,omitempty" xml:"TablePrivilegeObject,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetColumnPrivilegeObject() *string {
	return s.ColumnPrivilegeObject
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetDatabasePrivilegeObject() *string {
	return s.DatabasePrivilegeObject
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountPrivilegeObjectsRequest) GetTablePrivilegeObject() *string {
	return s.TablePrivilegeObject
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetAccountName(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetColumnPrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.ColumnPrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetDBClusterId(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetDatabasePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.DatabasePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPageNumber(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPageSize(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetPrivilegeType(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetRegionId(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) SetTablePrivilegeObject(v string) *DescribeAccountPrivilegeObjectsRequest {
	s.TablePrivilegeObject = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsRequest) Validate() error {
	return dara.Validate(s)
}
