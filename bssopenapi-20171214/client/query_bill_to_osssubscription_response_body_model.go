// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillToOSSSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBillToOSSSubscriptionResponseBody
	GetCode() *string
	SetData(v *QueryBillToOSSSubscriptionResponseBodyData) *QueryBillToOSSSubscriptionResponseBody
	GetData() *QueryBillToOSSSubscriptionResponseBodyData
	SetMessage(v string) *QueryBillToOSSSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBillToOSSSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBillToOSSSubscriptionResponseBody
	GetSuccess() *bool
}

type QueryBillToOSSSubscriptionResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryBillToOSSSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9EEAE43F-1E4D-4734-AE93-5049878AC103
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBillToOSSSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBillToOSSSubscriptionResponseBody) GetData() *QueryBillToOSSSubscriptionResponseBodyData {
	return s.Data
}

func (s *QueryBillToOSSSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBillToOSSSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBillToOSSSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetCode(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetData(v *QueryBillToOSSSubscriptionResponseBodyData) *QueryBillToOSSSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetMessage(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetRequestId(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetSuccess(v bool) *QueryBillToOSSSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBillToOSSSubscriptionResponseBodyData struct {
	// The ID of the current query account.
	//
	// example:
	//
	// 185xxxxx03489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The name of the current query account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string                                          `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items       *QueryBillToOSSSubscriptionResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryBillToOSSSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) GetItems() *QueryBillToOSSSubscriptionResponseBodyDataItems {
	return s.Items
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetAccountID(v string) *QueryBillToOSSSubscriptionResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetAccountName(v string) *QueryBillToOSSSubscriptionResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetItems(v *QueryBillToOSSSubscriptionResponseBodyDataItems) *QueryBillToOSSSubscriptionResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBillToOSSSubscriptionResponseBodyDataItems struct {
	Item []*QueryBillToOSSSubscriptionResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItems) GetItem() []*QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	return s.Item
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItems) SetItem(v []*QueryBillToOSSSubscriptionResponseBodyDataItemsItem) *QueryBillToOSSSubscriptionResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryBillToOSSSubscriptionResponseBodyDataItemsItem struct {
	BucketOwnerId     *int64  `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	BucketPath        *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	RowLimitPerFile   *int32  `json:"RowLimitPerFile,omitempty" xml:"RowLimitPerFile,omitempty"`
	SubscribeBucket   *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	SubscribeLanguage *string `json:"SubscribeLanguage,omitempty" xml:"SubscribeLanguage,omitempty"`
	SubscribeTime     *string `json:"SubscribeTime,omitempty" xml:"SubscribeTime,omitempty"`
	SubscribeType     *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetBucketOwnerId() *int64 {
	return s.BucketOwnerId
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetBucketPath() *string {
	return s.BucketPath
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetRowLimitPerFile() *int32 {
	return s.RowLimitPerFile
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetSubscribeBucket() *string {
	return s.SubscribeBucket
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetSubscribeLanguage() *string {
	return s.SubscribeLanguage
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetSubscribeTime() *string {
	return s.SubscribeTime
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetBucketOwnerId(v int64) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.BucketOwnerId = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetBucketPath(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.BucketPath = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetRowLimitPerFile(v int32) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.RowLimitPerFile = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeBucket(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeBucket = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeLanguage(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeLanguage = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeTime(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeTime = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeType(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeType = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
