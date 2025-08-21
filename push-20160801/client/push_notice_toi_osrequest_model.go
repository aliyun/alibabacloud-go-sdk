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
	// This parameter is required.
	//
	// example:
	//
	// DEV
	ApnsEnv *string `json:"ApnsEnv,omitempty" xml:"ApnsEnv,omitempty"`
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
	// Hello World
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// {"k1":"ios","k2":"v2"}
	ExtParameters *string `json:"ExtParameters,omitempty" xml:"ExtParameters,omitempty"`
	// example:
	//
	// 123
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
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
