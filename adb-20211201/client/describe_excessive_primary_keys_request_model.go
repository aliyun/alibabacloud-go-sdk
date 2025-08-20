// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcessivePrimaryKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeExcessivePrimaryKeysRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeExcessivePrimaryKeysRequest
	GetEndTime() *string
	SetLang(v string) *DescribeExcessivePrimaryKeysRequest
	GetLang() *string
	SetOrder(v string) *DescribeExcessivePrimaryKeysRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeExcessivePrimaryKeysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExcessivePrimaryKeysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeExcessivePrimaryKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExcessivePrimaryKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeExcessivePrimaryKeysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeExcessivePrimaryKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeExcessivePrimaryKeysRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeExcessivePrimaryKeysRequest
	GetStartTime() *string
}

type DescribeExcessivePrimaryKeysRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2019-06-01T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// 	- **zh (default)**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order in which table fields are sorted. Specify the value in the JSON format.
	//
	// Example:
	//
	//     [
	//
	//         {
	//
	//             "Field":"Name",
	//
	//             "Type":"Asc"
	//
	//         }
	//
	//     ]
	//
	// In the preceding code, Field specifies the field that is used to sort the table data. Set the value to Name. Type specifies the sorting order. Valid values: Desc and Asc.
	//
	// Field and Type are case-insensitive.
	//
	// example:
	//
	// [{"Field":"TableName", "Type": "Desc" }]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2019-06-01T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeExcessivePrimaryKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcessivePrimaryKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcessivePrimaryKeysRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeExcessivePrimaryKeysRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeExcessivePrimaryKeysRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExcessivePrimaryKeysRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeExcessivePrimaryKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExcessivePrimaryKeysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExcessivePrimaryKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExcessivePrimaryKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExcessivePrimaryKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExcessivePrimaryKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExcessivePrimaryKeysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeExcessivePrimaryKeysRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeExcessivePrimaryKeysRequest) SetDBClusterId(v string) *DescribeExcessivePrimaryKeysRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetEndTime(v string) *DescribeExcessivePrimaryKeysRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetLang(v string) *DescribeExcessivePrimaryKeysRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetOrder(v string) *DescribeExcessivePrimaryKeysRequest {
	s.Order = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetOwnerAccount(v string) *DescribeExcessivePrimaryKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetOwnerId(v int64) *DescribeExcessivePrimaryKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetPageNumber(v int32) *DescribeExcessivePrimaryKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetPageSize(v int32) *DescribeExcessivePrimaryKeysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetRegionId(v string) *DescribeExcessivePrimaryKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetResourceOwnerAccount(v string) *DescribeExcessivePrimaryKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetResourceOwnerId(v int64) *DescribeExcessivePrimaryKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) SetStartTime(v string) *DescribeExcessivePrimaryKeysRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeExcessivePrimaryKeysRequest) Validate() error {
	return dara.Validate(s)
}
