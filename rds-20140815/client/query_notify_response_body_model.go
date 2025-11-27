// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNotifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryNotifyResponseBodyData) *QueryNotifyResponseBody
	GetData() *QueryNotifyResponseBodyData
	SetRequestId(v string) *QueryNotifyResponseBody
	GetRequestId() *string
}

type QueryNotifyResponseBody struct {
	// The response parameters.
	Data *QueryNotifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 94CB8D93-017A-5AE7-A118-6E0F89D93C0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNotifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBody) GetData() *QueryNotifyResponseBodyData {
	return s.Data
}

func (s *QueryNotifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNotifyResponseBody) SetData(v *QueryNotifyResponseBodyData) *QueryNotifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryNotifyResponseBody) SetRequestId(v string) *QueryNotifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNotifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryNotifyResponseBodyData struct {
	// The details of notifications.
	NotifyItemList []*QueryNotifyResponseBodyDataNotifyItemList `json:"NotifyItemList,omitempty" xml:"NotifyItemList,omitempty" type:"Repeated"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 25
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s QueryNotifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBodyData) GetNotifyItemList() []*QueryNotifyResponseBodyDataNotifyItemList {
	return s.NotifyItemList
}

func (s *QueryNotifyResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNotifyResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNotifyResponseBodyData) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *QueryNotifyResponseBodyData) SetNotifyItemList(v []*QueryNotifyResponseBodyDataNotifyItemList) *QueryNotifyResponseBodyData {
	s.NotifyItemList = v
	return s
}

func (s *QueryNotifyResponseBodyData) SetPageNumber(v int32) *QueryNotifyResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryNotifyResponseBodyData) SetPageSize(v int32) *QueryNotifyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryNotifyResponseBodyData) SetTotalRecordCount(v int32) *QueryNotifyResponseBodyData {
	s.TotalRecordCount = &v
	return s
}

func (s *QueryNotifyResponseBodyData) Validate() error {
	if s.NotifyItemList != nil {
		for _, item := range s.NotifyItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryNotifyResponseBodyDataNotifyItemList struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 22973492**********
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Indicates whether the notification has been confirmed. You can call the [ConfirmNotify](https://help.aliyun.com/document_detail/610444.html) operation to mark the notification as confirmed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ConfirmFlag *bool `json:"ConfirmFlag,omitempty" xml:"ConfirmFlag,omitempty"`
	// The UID of the contact who called the [ConfirmNotify](https://help.aliyun.com/document_detail/610444.html) operation to mark the notification as confirmed. The contact belongs to the current Alibaba Cloud account.
	//
	// The value **0*	- indicates that the notification is automatically confirmed by the system.
	//
	// example:
	//
	// 0
	Confirmor *int64 `json:"Confirmor,omitempty" xml:"Confirmor,omitempty"`
	// The time when the notification was created.
	//
	// example:
	//
	// 2022-04-21T02:04:04Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the notification was modified.
	//
	// example:
	//
	// 2022-04-21T02:10:47Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the notification.
	//
	// example:
	//
	// 103499
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of times that repeatedly sent notifications are blocked.
	//
	// example:
	//
	// 0
	IdempotentCount *string `json:"IdempotentCount,omitempty" xml:"IdempotentCount,omitempty"`
	// This parameter ensures the idempotence of the notification and prevents the notification from being repeatedly sent.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	IdempotentId *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	// The level of the notification. Valid values:
	//
	// 	- **help**
	//
	// 	- **success**
	//
	// 	- **warning**
	//
	// 	- **error**
	//
	// 	- **loading**
	//
	// 	- **notice**
	//
	// example:
	//
	// error
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The element in the notification template. This parameter is a JSON string. Fields in the JSON string vary based on the value of the **TemplateName*	- parameter.
	//
	// 	- If the **TemplateName*	- parameter is **RenewalRecommend**, the JSON string contains the following fields:
	//
	//     	- **instanceName**: the ID of the instance that is about to expire
	//
	//     	- **reservedTime**: the remaining validity period of the instance in days
	//
	// 	- If the **TemplateName*	- parameter is **InstanceCreateFailed**, the JSON string contains the following fields:
	//
	//     	- **orderId**: the ID of the order to purchase the instance
	//
	//     	- **reason**: the cause of the instance creation failure
	//
	// example:
	//
	// {\\"orderId\\":21466**********}
	NotifyElement *string `json:"NotifyElement,omitempty" xml:"NotifyElement,omitempty"`
	// The template of the notification. Valid values:
	//
	// 	- **RenewalRecommend**: The template that is used to notify of renewal suggestions.
	//
	// 	- **InstanceCreateFailed**: The template that is used to notify that an instance fails to be created and is refunded.
	//
	// example:
	//
	// InstanceCreateFailed
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the notification. Valid values:
	//
	// 	- **Sell**: sales notification
	//
	// 	- **Operation**: O\\&M notification
	//
	// 	- **Promotion**: promotion notification
	//
	// example:
	//
	// Sell
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryNotifyResponseBodyDataNotifyItemList) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBodyDataNotifyItemList) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetConfirmFlag() *bool {
	return s.ConfirmFlag
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetConfirmor() *int64 {
	return s.Confirmor
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetId() *int64 {
	return s.Id
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetIdempotentCount() *string {
	return s.IdempotentCount
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetLevel() *string {
	return s.Level
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetNotifyElement() *string {
	return s.NotifyElement
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetType() *string {
	return s.Type
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetAliUid(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.AliUid = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetConfirmFlag(v bool) *QueryNotifyResponseBodyDataNotifyItemList {
	s.ConfirmFlag = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetConfirmor(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Confirmor = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetGmtCreated(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.GmtCreated = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetGmtModified(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.GmtModified = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetId(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Id = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetIdempotentCount(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.IdempotentCount = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetIdempotentId(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.IdempotentId = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetLevel(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Level = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetNotifyElement(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.NotifyElement = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetTemplateName(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.TemplateName = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetType(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Type = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) Validate() error {
	return dara.Validate(s)
}
