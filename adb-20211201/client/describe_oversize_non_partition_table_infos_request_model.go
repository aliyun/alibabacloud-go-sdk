// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOversizeNonPartitionTableInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetLang() *string
	SetOrder(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeOversizeNonPartitionTableInfosRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeOversizeNonPartitionTableInfosRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeOversizeNonPartitionTableInfosRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeOversizeNonPartitionTableInfosRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeOversizeNonPartitionTableInfosRequest
	GetStartTime() *string
}

type DescribeOversizeNonPartitionTableInfosRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time is in UTC and follows the <i>yyyy-MM-ddTHH:mm:ssZ</i> format.
	//
	// > The end time must be later than the start time. The interval between the start time and the end time must be less than 7 days.
	//
	// example:
	//
	// 2024-05-11T05:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the title for the downloaded file and specific error messages. Valid values:
	//
	// - **zh**: Simplified Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// - **ja**: Japanese.
	//
	// - **zh-tw**: Traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sorting rule for a specified field. The value is a JSON array. You must specify the fields to be sorted and the sorting order in sequence. Example:
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
	// Field: the name of the field to be sorted. Only the Name field is supported.
	//
	// Type: the sorting order. Valid values: Desc for descending order, and Asc for ascending order.
	//
	// The values are not case-sensitive.
	//
	// example:
	//
	// [{"Field":"SchemaName","Type":"Asc"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer. Default value: **1**.
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
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the query. The time is in UTC and follows the yyyy-MM-ddTHH:mmZ format.
	//
	// > SQL audit logs can be queried only when SQL audit is enabled. You can query SQL audit logs generated within the last 30 days. If SQL audit is disabled and then re-enabled, you can query only the SQL audit logs generated after it is re-enabled.
	//
	// example:
	//
	// 2024-05-11T05:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOversizeNonPartitionTableInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOversizeNonPartitionTableInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetDBClusterId(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetEndTime(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetLang(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetOrder(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.Order = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetOwnerAccount(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetOwnerId(v int64) *DescribeOversizeNonPartitionTableInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetPageNumber(v int32) *DescribeOversizeNonPartitionTableInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetPageSize(v int32) *DescribeOversizeNonPartitionTableInfosRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetRegionId(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetResourceOwnerAccount(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetResourceOwnerId(v int64) *DescribeOversizeNonPartitionTableInfosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) SetStartTime(v string) *DescribeOversizeNonPartitionTableInfosRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosRequest) Validate() error {
	return dara.Validate(s)
}
