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
	// The [edition](https://help.aliyun.com/document_detail/183258.html) of the cluster. Valid values:
	//
	// - **Normal**: Cluster Edition
	//
	// - **Basic**: single node
	//
	// - **Archive**: X-Engine
	//
	// - **NormalMultimaster**: Multi-master Cluster Edition
	//
	// - **SENormal**: Standard Edition
	//
	// > 	- The single node edition is not supported on PolarDB for PostgreSQL clusters that run PostgreSQL 11.
	//
	// >
	//
	// > 	- The Standard Edition is supported on PolarDB for MySQL clusters that run MySQL 8.0 or 5.7, and on PolarDB for PostgreSQL clusters that run PostgreSQL 14.
	//
	// >
	//
	// > 	- PolarDB for MySQL clusters that run MySQL 8.0 support X-Engine and the Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the current instance.
	//
	// example:
	//
	// pc-2zed3m89cw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return for the current request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. If the query results are not returned in a single call, this token is returned. Use this token in a subsequent call to retrieve the remaining results.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the order.
	//
	// - **pending**: The task is waiting to start.
	//
	// - **create**: The order is placed and is being processed.
	//
	// - **fail**: The instance failed to be created.
	//
	// - **cancel**: The order is canceled.
	//
	// - **success**: The instance is created.
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
	// The number of entries to return on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
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
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to view the details of regions.
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
