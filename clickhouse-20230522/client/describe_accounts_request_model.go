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
	SetPageNumber(v string) *DescribeAccountsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAccountsRequest
	GetPageSize() *string
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
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the cloud service.
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

func (s *DescribeAccountsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAccountsRequest) GetPageSize() *string {
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

func (s *DescribeAccountsRequest) SetPageNumber(v string) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v string) *DescribeAccountsRequest {
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
