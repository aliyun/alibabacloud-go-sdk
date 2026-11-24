// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeProcessListRequest
	GetDBClusterId() *string
	SetInitialQueryId(v string) *DescribeProcessListRequest
	GetInitialQueryId() *string
	SetInitialUser(v string) *DescribeProcessListRequest
	GetInitialUser() *string
	SetKeyword(v string) *DescribeProcessListRequest
	GetKeyword() *string
	SetOrder(v string) *DescribeProcessListRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeProcessListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeProcessListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeProcessListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProcessListRequest
	GetPageSize() *int32
	SetQueryDurationMs(v int32) *DescribeProcessListRequest
	GetQueryDurationMs() *int32
	SetRegionId(v string) *DescribeProcessListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeProcessListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeProcessListRequest
	GetResourceOwnerId() *int64
}

type DescribeProcessListRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query all clusters in the specified region, including their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1190tj036am****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the query statement.
	//
	// example:
	//
	// 6c69d508-f05f-4c74-857c-d982b7e7e79f
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The database account.
	//
	// example:
	//
	// user
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The query keyword.
	//
	// example:
	//
	// join
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The column by which to sort the results. Valid values:
	//
	// - elapsed: cumulative execution time
	//
	// - written_rows: number of written rows
	//
	// - read_rows: number of read rows
	//
	// - memory_usage: memory usage
	//
	// example:
	//
	// elapsed
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: integers greater than 0 and less than or equal to the maximum value of the Integer data type. Default value: 1.
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
	// The minimum query duration. Minimum value: **1000**. Default value: **1000**. Unit: milliseconds. The response returns queries whose duration exceeds this value.
	//
	// example:
	//
	// 500
	QueryDurationMs *int32 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeProcessListRequest) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeProcessListRequest) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeProcessListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeProcessListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeProcessListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeProcessListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeProcessListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProcessListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProcessListRequest) GetQueryDurationMs() *int32 {
	return s.QueryDurationMs
}

func (s *DescribeProcessListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProcessListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeProcessListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeProcessListRequest) SetDBClusterId(v string) *DescribeProcessListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialQueryId(v string) *DescribeProcessListRequest {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListRequest) SetInitialUser(v string) *DescribeProcessListRequest {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetOrder(v string) *DescribeProcessListRequest {
	s.Order = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerAccount(v string) *DescribeProcessListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerId(v int64) *DescribeProcessListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetQueryDurationMs(v int32) *DescribeProcessListRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListRequest) SetRegionId(v string) *DescribeProcessListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerAccount(v string) *DescribeProcessListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerId(v int64) *DescribeProcessListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) Validate() error {
	return dara.Validate(s)
}
