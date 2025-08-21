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
	// my body
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
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
	// all
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
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
