// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailedNotificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationList(v []*DescribeFailedNotificationsResponseBodyNotificationList) *DescribeFailedNotificationsResponseBody
	GetNotificationList() []*DescribeFailedNotificationsResponseBodyNotificationList
	SetPageNumber(v int32) *DescribeFailedNotificationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFailedNotificationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFailedNotificationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFailedNotificationsResponseBody
	GetTotalCount() *int32
}

type DescribeFailedNotificationsResponseBody struct {
	NotificationList []*DescribeFailedNotificationsResponseBodyNotificationList `json:"NotificationList,omitempty" xml:"NotificationList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFailedNotificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailedNotificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFailedNotificationsResponseBody) GetNotificationList() []*DescribeFailedNotificationsResponseBodyNotificationList {
	return s.NotificationList
}

func (s *DescribeFailedNotificationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFailedNotificationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFailedNotificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFailedNotificationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFailedNotificationsResponseBody) SetNotificationList(v []*DescribeFailedNotificationsResponseBodyNotificationList) *DescribeFailedNotificationsResponseBody {
	s.NotificationList = v
	return s
}

func (s *DescribeFailedNotificationsResponseBody) SetPageNumber(v int32) *DescribeFailedNotificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBody) SetPageSize(v int32) *DescribeFailedNotificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBody) SetRequestId(v string) *DescribeFailedNotificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBody) SetTotalCount(v int32) *DescribeFailedNotificationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFailedNotificationsResponseBodyNotificationList struct {
	// example:
	//
	// orderPay
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 1665988512000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// {"productCode":"cmgj00001","orderId":"123456","orderBizId":"111222","action":"orderPay","aliUid":"12211222211","skuId":"cmgj00001-prepay"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// deab3673236843a3b378c7d8d25c5f01
	NotificationRequestId *string `json:"NotificationRequestId,omitempty" xml:"NotificationRequestId,omitempty"`
}

func (s DescribeFailedNotificationsResponseBodyNotificationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailedNotificationsResponseBodyNotificationList) GoString() string {
	return s.String()
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) GetAction() *string {
	return s.Action
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) GetMessage() *string {
	return s.Message
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) GetNotificationRequestId() *string {
	return s.NotificationRequestId
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) SetAction(v string) *DescribeFailedNotificationsResponseBodyNotificationList {
	s.Action = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) SetCreateTime(v int64) *DescribeFailedNotificationsResponseBodyNotificationList {
	s.CreateTime = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) SetMessage(v string) *DescribeFailedNotificationsResponseBodyNotificationList {
	s.Message = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) SetNotificationRequestId(v string) *DescribeFailedNotificationsResponseBodyNotificationList {
	s.NotificationRequestId = &v
	return s
}

func (s *DescribeFailedNotificationsResponseBodyNotificationList) Validate() error {
	return dara.Validate(s)
}
