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
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// body
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// {"key1":"value1","api_name":"PushNoticeToAndroidRequest"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// example:
	//
	// 123
	JobKey       *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	StoreOffline *bool   `json:"StoreOffline,omitempty" xml:"StoreOffline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
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
