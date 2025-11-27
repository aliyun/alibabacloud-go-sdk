// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesByExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstancesByExpireTimeResponseBodyItems) *DescribeDBInstancesByExpireTimeResponseBody
	GetItems() *DescribeDBInstancesByExpireTimeResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstancesByExpireTimeResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstancesByExpireTimeResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstancesByExpireTimeResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstancesByExpireTimeResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstancesByExpireTimeResponseBody struct {
	// The details of the instances.
	Items *DescribeDBInstancesByExpireTimeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page. Valid values: any **non-zero*	- positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on the current page.
	//
	// example:
	//
	// 2
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 200
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesByExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) GetItems() *DescribeDBInstancesByExpireTimeResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) SetItems(v *DescribeDBInstancesByExpireTimeResponseBodyItems) *DescribeDBInstancesByExpireTimeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) SetPageNumber(v int32) *DescribeDBInstancesByExpireTimeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesByExpireTimeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) SetRequestId(v string) *DescribeDBInstancesByExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesByExpireTimeResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesByExpireTimeResponseBodyItems struct {
	DBInstanceExpireTime []*DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime `json:"DBInstanceExpireTime,omitempty" xml:"DBInstanceExpireTime,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesByExpireTimeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByExpireTimeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItems) GetDBInstanceExpireTime() []*DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	return s.DBInstanceExpireTime
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItems) SetDBInstanceExpireTime(v []*DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) *DescribeDBInstancesByExpireTimeResponseBodyItems {
	s.DBInstanceExpireTime = v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItems) Validate() error {
	if s.DBInstanceExpireTime != nil {
		for _, item := range s.DBInstanceExpireTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime struct {
	// The description of the instance.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance state table](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The expiration time of the instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > : Pay-as-you-go instances never expire.
	//
	// example:
	//
	// 2019-03-27T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked after it expires.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before it is rolled back.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked after its storage capacity is exhausted.
	//
	// 	- **LockReadInstanceByDiskQuota**: The instance is a read-only instance and is automatically locked after its storage capacity is exhausted.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetDBInstanceDescription(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetDBInstanceId(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetDBInstanceStatus(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetExpireTime(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetLockMode(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) SetPayType(v string) *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesByExpireTimeResponseBodyItemsDBInstanceExpireTime) Validate() error {
	return dara.Validate(s)
}
