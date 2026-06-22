// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessengerSubscriptionTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMessengerSubscriptionTokenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListMessengerSubscriptionTokenResponseBody
	GetCode() *string
	SetData(v []*ListMessengerSubscriptionTokenResponseBodyData) *ListMessengerSubscriptionTokenResponseBody
	GetData() []*ListMessengerSubscriptionTokenResponseBodyData
	SetMessage(v string) *ListMessengerSubscriptionTokenResponseBody
	GetMessage() *string
	SetNextPage(v string) *ListMessengerSubscriptionTokenResponseBody
	GetNextPage() *string
	SetRequestId(v string) *ListMessengerSubscriptionTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMessengerSubscriptionTokenResponseBody
	GetSuccess() *bool
}

type ListMessengerSubscriptionTokenResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - A return value of OK indicates that the request was successful.
	//
	// - For a list of other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListMessengerSubscriptionTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The identifier for the next page. Pass this value to query the next page.
	//
	// example:
	//
	// 3399***
	NextPage *string `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ei**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMessengerSubscriptionTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessengerSubscriptionTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetData() []*ListMessengerSubscriptionTokenResponseBodyData {
	return s.Data
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetNextPage() *string {
	return s.NextPage
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessengerSubscriptionTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetAccessDeniedDetail(v string) *ListMessengerSubscriptionTokenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetCode(v string) *ListMessengerSubscriptionTokenResponseBody {
	s.Code = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetData(v []*ListMessengerSubscriptionTokenResponseBodyData) *ListMessengerSubscriptionTokenResponseBody {
	s.Data = v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetMessage(v string) *ListMessengerSubscriptionTokenResponseBody {
	s.Message = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetNextPage(v string) *ListMessengerSubscriptionTokenResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetRequestId(v string) *ListMessengerSubscriptionTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) SetSuccess(v bool) *ListMessengerSubscriptionTokenResponseBody {
	s.Success = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMessengerSubscriptionTokenResponseBodyData struct {
	// The time when the security token was created.
	//
	// example:
	//
	// 172546854885
	CreationTimestamp *int64 `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The custom audience ID.
	//
	// example:
	//
	// Example
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// The time when the application can send the next marketing message to the recipient.
	//
	// example:
	//
	// 172546854885
	NextEligibleTime *int64 `json:"NextEligibleTime,omitempty" xml:"NextEligibleTime,omitempty"`
	// The re-subscription status.
	//
	// example:
	//
	// Example
	NotificationMessagesReoptin *string `json:"NotificationMessagesReoptin,omitempty" xml:"NotificationMessagesReoptin,omitempty"`
	// The customer\\"s time zone.
	//
	// example:
	//
	// Example
	NotificationMessagesTimezone *string `json:"NotificationMessagesTimezone,omitempty" xml:"NotificationMessagesTimezone,omitempty"`
	// The page-scoped ID of the customer.
	//
	// example:
	//
	// Example
	NotificationMessagesToken *string `json:"NotificationMessagesToken,omitempty" xml:"NotificationMessagesToken,omitempty"`
	// The page ID for Messenger.
	//
	// example:
	//
	// 239***
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The customer\\"s Page-scoped ID (PSID)
	//
	// example:
	//
	// Example
	RecipientId *string `json:"RecipientId,omitempty" xml:"RecipientId,omitempty"`
	// The expiration time of the security token.
	//
	// example:
	//
	// 172546854885
	TokenExpiryTimestamp *int64 `json:"TokenExpiryTimestamp,omitempty" xml:"TokenExpiryTimestamp,omitempty"`
	// The message\\"s title
	//
	// example:
	//
	// Example
	TopicTitle *string `json:"TopicTitle,omitempty" xml:"TopicTitle,omitempty"`
	// The status of the security token.
	//
	// example:
	//
	// Example
	UserTokenStatus *string `json:"UserTokenStatus,omitempty" xml:"UserTokenStatus,omitempty"`
}

func (s ListMessengerSubscriptionTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMessengerSubscriptionTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetCreationTimestamp() *int64 {
	return s.CreationTimestamp
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetNextEligibleTime() *int64 {
	return s.NextEligibleTime
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetNotificationMessagesReoptin() *string {
	return s.NotificationMessagesReoptin
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetNotificationMessagesTimezone() *string {
	return s.NotificationMessagesTimezone
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetNotificationMessagesToken() *string {
	return s.NotificationMessagesToken
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetRecipientId() *string {
	return s.RecipientId
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetTokenExpiryTimestamp() *int64 {
	return s.TokenExpiryTimestamp
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetTopicTitle() *string {
	return s.TopicTitle
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) GetUserTokenStatus() *string {
	return s.UserTokenStatus
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetCreationTimestamp(v int64) *ListMessengerSubscriptionTokenResponseBodyData {
	s.CreationTimestamp = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetCustomAudienceId(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.CustomAudienceId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetNextEligibleTime(v int64) *ListMessengerSubscriptionTokenResponseBodyData {
	s.NextEligibleTime = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetNotificationMessagesReoptin(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.NotificationMessagesReoptin = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetNotificationMessagesTimezone(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.NotificationMessagesTimezone = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetNotificationMessagesToken(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.NotificationMessagesToken = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetPageId(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.PageId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetRecipientId(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.RecipientId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetTokenExpiryTimestamp(v int64) *ListMessengerSubscriptionTokenResponseBodyData {
	s.TokenExpiryTimestamp = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetTopicTitle(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.TopicTitle = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) SetUserTokenStatus(v string) *ListMessengerSubscriptionTokenResponseBodyData {
	s.UserTokenStatus = &v
	return s
}

func (s *ListMessengerSubscriptionTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
