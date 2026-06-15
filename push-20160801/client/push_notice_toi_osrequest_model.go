// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToiOSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApnsEnv(v string) *PushNoticeToiOSRequest
	GetApnsEnv() *string
	SetAppKey(v int64) *PushNoticeToiOSRequest
	GetAppKey() *int64
	SetBody(v string) *PushNoticeToiOSRequest
	GetBody() *string
	SetExtParameters(v string) *PushNoticeToiOSRequest
	GetExtParameters() *string
	SetJobKey(v string) *PushNoticeToiOSRequest
	GetJobKey() *string
	SetTarget(v string) *PushNoticeToiOSRequest
	GetTarget() *string
	SetTargetValue(v string) *PushNoticeToiOSRequest
	GetTargetValue() *string
	SetTitle(v string) *PushNoticeToiOSRequest
	GetTitle() *string
}

type PushNoticeToiOSRequest struct {
	// iOS notifications use Apple’s APNs service. Specify the environment.
	//
	// - DEV: Development environment.
	//
	// - PRODUCT: Production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	ApnsEnv *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
	// Your AppKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The body text of the notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello World
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// A custom key-value map for developer extensions.
	//
	// > For iOS devices, pass this parameter as a JSON object. Otherwise, parsing fails.
	//
	// example:
	//
	// {"k1":"ios","k2":"v2"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// A custom ID for the push task. If you specify a JobKey, the delivery log includes this field. For more information, see [Delivery logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
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
	// Values depend on the Target value. Separate multiple values with commas. If you exceed the limit, send multiple requests.
	//
	// - If Target=DEVICE, use values such as `deviceid111,deviceid1111`. Maximum: 1000.
	//
	// - If Target=ACCOUNT, use values such as `account111,account222`. Maximum: 1000.
	//
	// - If Target=ALIAS, use values such as `alias111,alias222`. Maximum: 1000.
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
	// - iOS 10 and later: Displays as the notification title.
	//
	// - iOS 8.2 through iOS 9.x: Replaces the app name in the notification.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushNoticeToiOSRequest) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToiOSRequest) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSRequest) GetApnsEnv() *string {
	return s.ApnsEnv
}

func (s *PushNoticeToiOSRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushNoticeToiOSRequest) GetBody() *string {
	return s.Body
}

func (s *PushNoticeToiOSRequest) GetExtParameters() *string {
	return s.ExtParameters
}

func (s *PushNoticeToiOSRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushNoticeToiOSRequest) GetTarget() *string {
	return s.Target
}

func (s *PushNoticeToiOSRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushNoticeToiOSRequest) GetTitle() *string {
	return s.Title
}

func (s *PushNoticeToiOSRequest) SetApnsEnv(v string) *PushNoticeToiOSRequest {
	s.ApnsEnv = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetAppKey(v int64) *PushNoticeToiOSRequest {
	s.AppKey = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetBody(v string) *PushNoticeToiOSRequest {
	s.Body = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetExtParameters(v string) *PushNoticeToiOSRequest {
	s.ExtParameters = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetJobKey(v string) *PushNoticeToiOSRequest {
	s.JobKey = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTarget(v string) *PushNoticeToiOSRequest {
	s.Target = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTargetValue(v string) *PushNoticeToiOSRequest {
	s.TargetValue = &v
	return s
}

func (s *PushNoticeToiOSRequest) SetTitle(v string) *PushNoticeToiOSRequest {
	s.Title = &v
	return s
}

func (s *PushNoticeToiOSRequest) Validate() error {
	return dara.Validate(s)
}
