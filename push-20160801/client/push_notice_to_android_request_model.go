// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToAndroidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *PushNoticeToAndroidRequest
	GetAppKey() *int64
	SetBody(v string) *PushNoticeToAndroidRequest
	GetBody() *string
	SetExtParameters(v string) *PushNoticeToAndroidRequest
	GetExtParameters() *string
	SetJobKey(v string) *PushNoticeToAndroidRequest
	GetJobKey() *string
	SetStoreOffline(v bool) *PushNoticeToAndroidRequest
	GetStoreOffline() *bool
	SetTarget(v string) *PushNoticeToAndroidRequest
	GetTarget() *string
	SetTargetValue(v string) *PushNoticeToAndroidRequest
	GetTargetValue() *string
	SetTitle(v string) *PushNoticeToAndroidRequest
	GetTitle() *string
}

type PushNoticeToAndroidRequest struct {
	// Your AppKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The body of the notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// body
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// Custom key-value pairs for Android-specific extensions. Pass this as a JSON object.
	//
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// A custom ID for the push task. If you specify a non-empty JobKey, it appears in the delivery receipt log. For more information, see [Delivery receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Whether to store the notification for offline delivery. Default: false.
	//
	// If enabled, the notification is redelivered when the user comes online within the time-to-live (TTL) period. Default TTL: 72 hours.
	//
	// example:
	//
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The target of the push. Valid values:
	//
	// - **DEVICE**: Push to specific devices.
	//
	// - **ACCOUNT**: Push to specific accounts.
	//
	// - **ALIAS**: Push to users with specific aliases.
	//
	// - **TAG**: Push to users with specific tags.
	//
	// - **ALL**: Push to all devices.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set this based on the Target value. Separate multiple values with commas. If you exceed the limit, send multiple requests.
	//
	// - If Target=DEVICE, use values such as `deviceid111,deviceid1111`. Maximum: 1000 devices.
	//
	// - If Target=ACCOUNT, use values such as `account111,account222`. Maximum: 1000 accounts.
	//
	// - If Target=ALIAS, use values such as `alias111,alias222`. Maximum: 1000 aliases.
	//
	// - If Target=TAG, support single or multiple tags. For format details, see [Tag format](https://help.aliyun.com/document_detail/434847.html).
	//
	// - If Target=ALL, set this to **ALL**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushNoticeToAndroidRequest) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToAndroidRequest) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushNoticeToAndroidRequest) GetBody() *string {
	return s.Body
}

func (s *PushNoticeToAndroidRequest) GetExtParameters() *string {
	return s.ExtParameters
}

func (s *PushNoticeToAndroidRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushNoticeToAndroidRequest) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *PushNoticeToAndroidRequest) GetTarget() *string {
	return s.Target
}

func (s *PushNoticeToAndroidRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushNoticeToAndroidRequest) GetTitle() *string {
	return s.Title
}

func (s *PushNoticeToAndroidRequest) SetAppKey(v int64) *PushNoticeToAndroidRequest {
	s.AppKey = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetBody(v string) *PushNoticeToAndroidRequest {
	s.Body = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetExtParameters(v string) *PushNoticeToAndroidRequest {
	s.ExtParameters = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetJobKey(v string) *PushNoticeToAndroidRequest {
	s.JobKey = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetStoreOffline(v bool) *PushNoticeToAndroidRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTarget(v string) *PushNoticeToAndroidRequest {
	s.Target = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTargetValue(v string) *PushNoticeToAndroidRequest {
	s.TargetValue = &v
	return s
}

func (s *PushNoticeToAndroidRequest) SetTitle(v string) *PushNoticeToAndroidRequest {
	s.Title = &v
	return s
}

func (s *PushNoticeToAndroidRequest) Validate() error {
	return dara.Validate(s)
}
