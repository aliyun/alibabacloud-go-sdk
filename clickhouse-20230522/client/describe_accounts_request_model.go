// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAccountsRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *DescribeAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccountsRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAccountsRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeAccountsRequest
	GetRegionId() *string
}

type DescribeAccountsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product code.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccountsRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageNumber(v int32) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v int32) *DescribeAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountsRequest) SetProduct(v string) *DescribeAccountsRequest {
	s.Product = &v
	return s
}

func (s *DescribeAccountsRequest) SetRegionId(v string) *DescribeAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
