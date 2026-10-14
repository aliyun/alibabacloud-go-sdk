// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListOrdersRequest
	GetCategory() *string
	SetInstanceId(v string) *ListOrdersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListOrdersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOrdersRequest
	GetNextToken() *string
	SetOrderStatus(v string) *ListOrdersRequest
	GetOrderStatus() *string
	SetPageNumber(v int32) *ListOrdersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOrdersRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListOrdersRequest
	GetProductCode() *string
	SetRegionId(v string) *ListOrdersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListOrdersRequest
	GetResourceGroupId() *string
}

type ListOrdersRequest struct {
	// The [product edition](https://help.aliyun.com/document_detail/183258.html). Valid values:
	//
	// 	- **Normal**: Cluster Edition.
	//
	// 	- **Basic**: Single Node Edition.
	//
	// 	- **Archive**: X-Engine (Warm Data).
	//
	// 	- **NormalMultimaster**: Multi-master Cluster Edition.
	//
	// 	- **SENormal**: Standard Edition.
	//
	// > 	- PolarDB for PostgreSQL 11 does not support Single Node Edition.
	//
	// >	- PolarDB for MySQL 8.0, PolarDB for MySQL 5.7, and PolarDB for PostgreSQL 14 support Standard Edition.
	//
	// >	- PolarDB for MySQL 8.0 supports X-Engine (Warm Data) and Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pc-2zed3m89cw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used to retrieve more results. If a single request does not return all results, this parameter is returned. You can pass the returned token in a subsequent request to continue retrieving results.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The order status. Valid values:
	//
	// - **pending**: The task is waiting to start.
	//
	// - **create**: The order has been placed and entered the production process.
	//
	// - **fail**: Production failed.
	//
	// - **cancel**: The order is canceled.
	//
	// - **success**: Production succeeded.
	//
	// example:
	//
	// success
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product code.
	//
	// example:
	//
	// polardb
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region information.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListOrdersRequest) GetCategory() *string {
	return s.Category
}

func (s *ListOrdersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrdersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOrdersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrdersRequest) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListOrdersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrdersRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListOrdersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOrdersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListOrdersRequest) SetCategory(v string) *ListOrdersRequest {
	s.Category = &v
	return s
}

func (s *ListOrdersRequest) SetInstanceId(v string) *ListOrdersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrdersRequest) SetMaxResults(v int32) *ListOrdersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOrdersRequest) SetNextToken(v string) *ListOrdersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOrdersRequest) SetOrderStatus(v string) *ListOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListOrdersRequest) SetPageNumber(v int32) *ListOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrdersRequest) SetPageSize(v int32) *ListOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrdersRequest) SetProductCode(v string) *ListOrdersRequest {
	s.ProductCode = &v
	return s
}

func (s *ListOrdersRequest) SetRegionId(v string) *ListOrdersRequest {
	s.RegionId = &v
	return s
}

func (s *ListOrdersRequest) SetResourceGroupId(v string) *ListOrdersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListOrdersRequest) Validate() error {
	return dara.Validate(s)
}
