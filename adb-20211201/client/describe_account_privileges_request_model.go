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
	// The name of the database account whose privileges you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Filters the results by column name. This parameter is used only when `PrivilegeType` is set to `Column`.
	//
	// example:
	//
	// col1
	ColumnPrivilegeObject *string `json:"ColumnPrivilegeObject,omitempty" xml:"ColumnPrivilegeObject,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Filters the results by database name. This parameter is used only when `PrivilegeType` is set to `Database`, `Table`, or `Column`.
	//
	// example:
	//
	// db1
	DatabasePrivilegeObject *string `json:"DatabasePrivilegeObject,omitempty" xml:"DatabasePrivilegeObject,omitempty"`
	// The page number. Pages start at 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The privilege level to query. To obtain the valid values for this parameter, call the `DescribeEnabledPrivileges` operation.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Filters the results by table name. You can use this parameter with `DatabasePrivilegeObject` to refine the search. This parameter is used only when `PrivilegeType` is set to `Table` or `Column`.
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
