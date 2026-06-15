// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToAndroidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *PushMessageToAndroidRequest
	GetAppKey() *int64
	SetBody(v string) *PushMessageToAndroidRequest
	GetBody() *string
	SetJobKey(v string) *PushMessageToAndroidRequest
	GetJobKey() *string
	SetStoreOffline(v bool) *PushMessageToAndroidRequest
	GetStoreOffline() *bool
	SetTarget(v string) *PushMessageToAndroidRequest
	GetTarget() *string
	SetTargetValue(v string) *PushMessageToAndroidRequest
	GetTargetValue() *string
	SetTitle(v string) *PushMessageToAndroidRequest
	GetTitle() *string
}

type PushMessageToAndroidRequest struct {
	// The AppKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The content of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// my body
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// A custom ID for the push task. If \\`JobKey\\` is not empty, this field is included in the receipt logs. For more information, see [Receipt logs](https://help.aliyun.com/document_detail/434651.html).
	//
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// Specifies whether to store the message offline. The default value is false.
	//
	// If you store the message and the user is offline, the message is sent again when the user comes online within the time-to-live (TTL) period. The default TTL is 72 hours.
	//
	// example:
	//
	// true
	StoreOffline *bool `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// The push target. Valid values:
	//
	// - **DEVICE**: Pushes messages to devices.
	//
	// - **ACCOUNT**: Pushes messages to accounts.
	//
	// - **ALIAS**: Pushes messages to aliases.
	//
	// - **TAG**: Pushes messages to tags.
	//
	// - **ALL**: Pushes messages to all devices.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// Set this parameter based on the value of \\`Target\\`. Use commas (,) to separate multiple values. If you exceed the limit, send the pushes in batches.
	//
	// - If \\`Target\\` is set to \\`DEVICE\\`, specify device IDs. Example: `deviceid111,deviceid1111`. You can specify up to 1,000 device IDs.
	//
	// - If \\`Target\\` is set to \\`ACCOUNT\\`, specify account IDs. Example: `account111,account222`. You can specify up to 1,000 account IDs.
	//
	// - If \\`Target\\` is set to \\`ALIAS\\`, specify aliases. Example: `alias111,alias222`. You can specify up to 1,000 aliases.
	//
	// - If \\`Target\\` is set to \\`TAG\\`, you can specify one or more tags. For more information about the format, see [Tag format](https://help.aliyun.com/document_detail/434847.html).
	//
	// - If \\`Target\\` is set to \\`ALL\\`, set the value to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
	// The title of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// my title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s PushMessageToAndroidRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToAndroidRequest) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushMessageToAndroidRequest) GetBody() *string {
	return s.Body
}

func (s *PushMessageToAndroidRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushMessageToAndroidRequest) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *PushMessageToAndroidRequest) GetTarget() *string {
	return s.Target
}

func (s *PushMessageToAndroidRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushMessageToAndroidRequest) GetTitle() *string {
	return s.Title
}

func (s *PushMessageToAndroidRequest) SetAppKey(v int64) *PushMessageToAndroidRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetBody(v string) *PushMessageToAndroidRequest {
	s.Body = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetJobKey(v string) *PushMessageToAndroidRequest {
	s.JobKey = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetStoreOffline(v bool) *PushMessageToAndroidRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTarget(v string) *PushMessageToAndroidRequest {
	s.Target = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTargetValue(v string) *PushMessageToAndroidRequest {
	s.TargetValue = &v
	return s
}

func (s *PushMessageToAndroidRequest) SetTitle(v string) *PushMessageToAndroidRequest {
	s.Title = &v
	return s
}

func (s *PushMessageToAndroidRequest) Validate() error {
	return dara.Validate(s)
}
