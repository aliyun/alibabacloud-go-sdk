// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToiOSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *PushMessageToiOSRequest
	GetAppKey() *int64
	SetBody(v string) *PushMessageToiOSRequest
	GetBody() *string
	SetJobKey(v string) *PushMessageToiOSRequest
	GetJobKey() *string
	SetStoreOffline(v bool) *PushMessageToiOSRequest
	GetStoreOffline() *bool
	SetTarget(v string) *PushMessageToiOSRequest
	GetTarget() *string
	SetTargetValue(v string) *PushMessageToiOSRequest
	GetTargetValue() *string
	SetTitle(v string) *PushMessageToiOSRequest
	GetTitle() *string
}

type PushMessageToiOSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 24780725
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

func (s PushMessageToiOSRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToiOSRequest) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *PushMessageToiOSRequest) GetBody() *string {
	return s.Body
}

func (s *PushMessageToiOSRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *PushMessageToiOSRequest) GetStoreOffline() *bool {
	return s.StoreOffline
}

func (s *PushMessageToiOSRequest) GetTarget() *string {
	return s.Target
}

func (s *PushMessageToiOSRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *PushMessageToiOSRequest) GetTitle() *string {
	return s.Title
}

func (s *PushMessageToiOSRequest) SetAppKey(v int64) *PushMessageToiOSRequest {
	s.AppKey = &v
	return s
}

func (s *PushMessageToiOSRequest) SetBody(v string) *PushMessageToiOSRequest {
	s.Body = &v
	return s
}

func (s *PushMessageToiOSRequest) SetJobKey(v string) *PushMessageToiOSRequest {
	s.JobKey = &v
	return s
}

func (s *PushMessageToiOSRequest) SetStoreOffline(v bool) *PushMessageToiOSRequest {
	s.StoreOffline = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTarget(v string) *PushMessageToiOSRequest {
	s.Target = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTargetValue(v string) *PushMessageToiOSRequest {
	s.TargetValue = &v
	return s
}

func (s *PushMessageToiOSRequest) SetTitle(v string) *PushMessageToiOSRequest {
	s.Title = &v
	return s
}

func (s *PushMessageToiOSRequest) Validate() error {
	return dara.Validate(s)
}
