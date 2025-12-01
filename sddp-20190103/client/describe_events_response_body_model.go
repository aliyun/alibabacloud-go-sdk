// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeEventsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeEventsResponseBodyItems) *DescribeEventsResponseBody
	GetItems() []*DescribeEventsResponseBodyItems
	SetPageSize(v int32) *DescribeEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEventsResponseBody
	GetTotalCount() *int32
}

type DescribeEventsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the anomalous events.
	Items []*DescribeEventsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventsResponseBody) GetItems() []*DescribeEventsResponseBodyItems {
	return s.Items
}

func (s *DescribeEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEventsResponseBody) SetCurrentPage(v int32) *DescribeEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsResponseBody) SetItems(v []*DescribeEventsResponseBodyItems) *DescribeEventsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventsResponseBodyItems struct {
	// The time when an alert was triggered for the anomalous event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 154529000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// Indicates whether the detection of anomalous events is enhanced. If the detection of anomalous events is enhanced, the detection accuracy and the rate of triggering alerts for anomalous events are improved. Valid values:
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// false
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The display name of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// yundunsr
	DealDisplayName *string `json:"DealDisplayName,omitempty" xml:"DealDisplayName,omitempty"`
	// The username of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// det1111
	DealLoginName *string `json:"DealLoginName,omitempty" xml:"DealLoginName,omitempty"`
	// The time when the anomalous event was handled. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 12223300
	DealTime *int64 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The ID of the account that is used to handle the anomalous event.
	//
	// example:
	//
	// 229157443385014***
	DealUserId *int64 `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The display name of the account that triggered the anomalous event.
	//
	// example:
	//
	// yundunsr
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the anomalous event occurred. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545829129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The ID of the anomalous event.
	//
	// example:
	//
	// 42233335555
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username of the account that triggered the anomalous event.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The name of the service in which the anomalous event was detected.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The handling status for the anomalous event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: confirmed
	//
	// 	- 2: marked as false positive
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the handling status for the anomalous event.
	//
	// example:
	//
	// Pending
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeName *string `json:"SubTypeName,omitempty" xml:"SubTypeName,omitempty"`
	// The name of the destination service in an anomalous data flow.
	//
	// example:
	//
	// RDS
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	// The code of the anomalous event type.
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The name of the anomalous event type.
	//
	// example:
	//
	// Anomalous data flow
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 1978132506596***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The severity of the anomalous event.
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeEventsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyItems) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *DescribeEventsResponseBodyItems) GetBacked() *bool {
	return s.Backed
}

func (s *DescribeEventsResponseBodyItems) GetDealDisplayName() *string {
	return s.DealDisplayName
}

func (s *DescribeEventsResponseBodyItems) GetDealLoginName() *string {
	return s.DealLoginName
}

func (s *DescribeEventsResponseBodyItems) GetDealTime() *int64 {
	return s.DealTime
}

func (s *DescribeEventsResponseBodyItems) GetDealUserId() *int64 {
	return s.DealUserId
}

func (s *DescribeEventsResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeEventsResponseBodyItems) GetEventTime() *int64 {
	return s.EventTime
}

func (s *DescribeEventsResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventsResponseBodyItems) GetLoginName() *string {
	return s.LoginName
}

func (s *DescribeEventsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeEventsResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventsResponseBodyItems) GetStatusName() *string {
	return s.StatusName
}

func (s *DescribeEventsResponseBodyItems) GetSubTypeCode() *string {
	return s.SubTypeCode
}

func (s *DescribeEventsResponseBodyItems) GetSubTypeName() *string {
	return s.SubTypeName
}

func (s *DescribeEventsResponseBodyItems) GetTargetProductCode() *string {
	return s.TargetProductCode
}

func (s *DescribeEventsResponseBodyItems) GetTypeCode() *string {
	return s.TypeCode
}

func (s *DescribeEventsResponseBodyItems) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeEventsResponseBodyItems) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeEventsResponseBodyItems) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *DescribeEventsResponseBodyItems) SetAlertTime(v int64) *DescribeEventsResponseBodyItems {
	s.AlertTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetBacked(v bool) *DescribeEventsResponseBodyItems {
	s.Backed = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DealDisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealLoginName(v string) *DescribeEventsResponseBodyItems {
	s.DealLoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealTime(v int64) *DescribeEventsResponseBodyItems {
	s.DealTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDealUserId(v int64) *DescribeEventsResponseBodyItems {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetDisplayName(v string) *DescribeEventsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetEventTime(v int64) *DescribeEventsResponseBodyItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetId(v int64) *DescribeEventsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetLoginName(v string) *DescribeEventsResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetProductCode(v string) *DescribeEventsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatus(v int32) *DescribeEventsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetStatusName(v string) *DescribeEventsResponseBodyItems {
	s.StatusName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeCode(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetSubTypeName(v string) *DescribeEventsResponseBodyItems {
	s.SubTypeName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTargetProductCode(v string) *DescribeEventsResponseBodyItems {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTypeCode(v string) *DescribeEventsResponseBodyItems {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetTypeName(v string) *DescribeEventsResponseBodyItems {
	s.TypeName = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetUserId(v int64) *DescribeEventsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) SetWarnLevel(v int32) *DescribeEventsResponseBodyItems {
	s.WarnLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
