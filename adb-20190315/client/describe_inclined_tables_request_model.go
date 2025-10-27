// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeInclinedTablesRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeInclinedTablesRequest
	GetLang() *string
	SetOrder(v string) *DescribeInclinedTablesRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeInclinedTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInclinedTablesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInclinedTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInclinedTablesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInclinedTablesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInclinedTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInclinedTablesRequest
	GetResourceOwnerId() *int64
	SetTableType(v string) *DescribeInclinedTablesRequest
	GetTableType() *string
}

type DescribeInclinedTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-k2jofo4pi5zhd****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language in which you want to send requests and receive messages. Default value: zh. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON format.
	//
	// Example:
	//
	// ```
	//
	// [
	//
	//     {
	//
	//         "Field":"Name",
	//
	//         "Type":"Asc"
	//
	//     }
	//
	// ]
	//
	// ```
	//
	// Field specifies the field by which to sort the query results. Set the value to Name. Type specifies the sorting order. Valid values: Desc and Asc.
	//
	// Field and Type are case-insensitive.
	//
	// example:
	//
	// [      {          "Field":"Name",          "Type":"Asc"      }  ]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - 30
	//
	// - 50
	//
	// - 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the table. Valid values:
	//
	//
	// - **FactTable**: the partitioned table.
	//
	// - **DimensionTable**: the dimension table.
	//
	// This parameter is required.
	//
	// example:
	//
	// FactTable
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeInclinedTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeInclinedTablesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInclinedTablesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeInclinedTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInclinedTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInclinedTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInclinedTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInclinedTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInclinedTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInclinedTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInclinedTablesRequest) GetTableType() *string {
	return s.TableType
}

func (s *DescribeInclinedTablesRequest) SetDBClusterId(v string) *DescribeInclinedTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetLang(v string) *DescribeInclinedTablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOrder(v string) *DescribeInclinedTablesRequest {
	s.Order = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageNumber(v int32) *DescribeInclinedTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetPageSize(v int32) *DescribeInclinedTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetRegionId(v string) *DescribeInclinedTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerAccount(v string) *DescribeInclinedTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetResourceOwnerId(v int64) *DescribeInclinedTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInclinedTablesRequest) SetTableType(v string) *DescribeInclinedTablesRequest {
	s.TableType = &v
	return s
}

func (s *DescribeInclinedTablesRequest) Validate() error {
	return dara.Validate(s)
}
