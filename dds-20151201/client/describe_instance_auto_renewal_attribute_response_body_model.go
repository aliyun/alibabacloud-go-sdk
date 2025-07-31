// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeInstanceAutoRenewalAttributeResponseBodyItems) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetItems() *DescribeInstanceAutoRenewalAttributeResponseBodyItems
	SetItemsNumbers(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetItemsNumbers() *int32
	SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeInstanceAutoRenewalAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAutoRenewalAttributeResponseBody struct {
	// Details about returned entries.
	Items *DescribeInstanceAutoRenewalAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	ItemsNumbers *int32 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries that were returned on the current page.
	//
	// example:
	//
	// 2
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FAB5CB3B-DB9D-473A-9DF1-F57B6B9CB949
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetItems() *DescribeInstanceAutoRenewalAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetItemsNumbers() *int32 {
	return s.ItemsNumbers
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetItems(v *DescribeInstanceAutoRenewalAttributeResponseBodyItems) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetItemsNumbers(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetPageRecordCount(v int32) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAutoRenewalAttributeResponseBodyItems struct {
	Item []*DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItems) GetItem() []*DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItems) SetItem(v []*DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) *DescribeInstanceAutoRenewalAttributeResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem struct {
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// 	- **true**: Auto-renewal is enabled for the instance.
	//
	// 	- **false**: Auto-renewal is disabled for the instance.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- **replicate**: the standalone or replica set instance
	//
	// 	- **sharding**: the sharded cluster instance
	//
	// example:
	//
	// replicate
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// dds-bp2568*****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The auto-renewal period. Unit: months.
	//
	// > 	- This parameter is ruturned only when the returned value of the **AutoRenew*	- parameter is **true**.
	//
	// > 	- You can call the [ModifyInstanceAutoRenewalAttribute](https://help.aliyun.com/document_detail/145979.html) operation to modify the auto-renewal period.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GetDuration() *string {
	return s.Duration
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetAutoRenew(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.AutoRenew = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDBInstanceType(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDbInstanceId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetDuration(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponseBodyItemsItem) Validate() error {
	return dara.Validate(s)
}
