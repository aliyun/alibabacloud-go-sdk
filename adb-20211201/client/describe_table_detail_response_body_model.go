// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeTableDetailResponseBody
	GetAccessDeniedDetail() *string
	SetAvgSize(v string) *DescribeTableDetailResponseBody
	GetAvgSize() *string
	SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody
	GetItems() *DescribeTableDetailResponseBodyItems
	SetPageNumber(v string) *DescribeTableDetailResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeTableDetailResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeTableDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeTableDetailResponseBody
	GetTotalCount() *string
}

type DescribeTableDetailResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {
	//
	//     "PolicyType": "AccountLevelIdentityBasedPolicy",
	//
	//     "AuthPrincipalOwnerId": "1*****************7",
	//
	//     "EncodedDiagnosticMessage": "AQIBIAAAAOPdwKY2QLOvgMEc7SkkoJfj1kvZwsaRqNYMh10Tv0wTe0fCzaCdrvgazfNb0EnJKETgXyhR+3BIQjx9WAqZryejBsp1Bl4qI5En/D9dEhcXAtKCxCmE2kZCiEzpy8BoEUt+bs0DmlaGWO5xkEpttypLIB4rUhDvZd+zwPg4EXk4KSSWSWsurxtqDkKEMshKlQFBTKvJcKwyhk62IeYly4hQ+5IpXjkh1GQXuDRCQ==",
	//
	//     "AuthPrincipalType": "SubUser",
	//
	//     "AuthPrincipalDisplayName": "2***************9",
	//
	//     "NoPermissionType": "ImplicitDeny",
	//
	//     "AuthAction": "adb:DescribeExcessivePrimaryKeys"
	//
	// }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The average number of rows in a shard.
	//
	// example:
	//
	// 160
	AvgSize *string `json:"AvgSize,omitempty" xml:"AvgSize,omitempty"`
	// The queried data distribution.
	Items *DescribeTableDetailResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeTableDetailResponseBody) GetAvgSize() *string {
	return s.AvgSize
}

func (s *DescribeTableDetailResponseBody) GetItems() *DescribeTableDetailResponseBodyItems {
	return s.Items
}

func (s *DescribeTableDetailResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeTableDetailResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableDetailResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeTableDetailResponseBody) SetAccessDeniedDetail(v string) *DescribeTableDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetAvgSize(v string) *DescribeTableDetailResponseBody {
	s.AvgSize = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetItems(v *DescribeTableDetailResponseBodyItems) *DescribeTableDetailResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableDetailResponseBody) SetPageNumber(v string) *DescribeTableDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetPageSize(v string) *DescribeTableDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetRequestId(v string) *DescribeTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableDetailResponseBody) SetTotalCount(v string) *DescribeTableDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTableDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTableDetailResponseBodyItems struct {
	// The queried shards.
	Shard []*DescribeTableDetailResponseBodyItemsShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Repeated"`
}

func (s DescribeTableDetailResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItems) GetShard() []*DescribeTableDetailResponseBodyItemsShard {
	return s.Shard
}

func (s *DescribeTableDetailResponseBodyItems) SetShard(v []*DescribeTableDetailResponseBodyItemsShard) *DescribeTableDetailResponseBodyItems {
	s.Shard = v
	return s
}

func (s *DescribeTableDetailResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeTableDetailResponseBodyItemsShard struct {
	// The shard ID. Only the numeric part of the shard name is returned.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 9484858
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTableDetailResponseBodyItemsShard) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponseBodyItemsShard) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponseBodyItemsShard) GetId() *int32 {
	return s.Id
}

func (s *DescribeTableDetailResponseBodyItemsShard) GetSize() *int64 {
	return s.Size
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetId(v int32) *DescribeTableDetailResponseBodyItemsShard {
	s.Id = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) SetSize(v int64) *DescribeTableDetailResponseBodyItemsShard {
	s.Size = &v
	return s
}

func (s *DescribeTableDetailResponseBodyItemsShard) Validate() error {
	return dara.Validate(s)
}
