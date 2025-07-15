// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcAsrTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *CreateRtcAsrTaskRequest
	GetAuthKey() *string
	SetAutoTerminateDelay(v int64) *CreateRtcAsrTaskRequest
	GetAutoTerminateDelay() *int64
	SetAutoTerminateEnabled(v bool) *CreateRtcAsrTaskRequest
	GetAutoTerminateEnabled() *bool
	SetCallbackURL(v string) *CreateRtcAsrTaskRequest
	GetCallbackURL() *string
	SetChannelID(v string) *CreateRtcAsrTaskRequest
	GetChannelID() *string
	SetLanguage(v string) *CreateRtcAsrTaskRequest
	GetLanguage() *string
	SetMode(v string) *CreateRtcAsrTaskRequest
	GetMode() *string
	SetOwnerId(v int64) *CreateRtcAsrTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateRtcAsrTaskRequest
	GetRegionId() *string
	SetReportInterval(v int64) *CreateRtcAsrTaskRequest
	GetReportInterval() *int64
	SetRtcUserId(v string) *CreateRtcAsrTaskRequest
	GetRtcUserId() *string
	SetSDKAppID(v string) *CreateRtcAsrTaskRequest
	GetSDKAppID() *string
	SetStreamURL(v string) *CreateRtcAsrTaskRequest
	GetStreamURL() *string
	SetTargetLanguages(v string) *CreateRtcAsrTaskRequest
	GetTargetLanguages() *string
	SetTranslateEnabled(v bool) *CreateRtcAsrTaskRequest
	GetTranslateEnabled() *bool
}

type CreateRtcAsrTaskRequest struct {
	// The AuthKey that is used to generate the MD5 signature in callbacks.
	//
	// example:
	//
	// abcd
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The maximum latency at which the task is automatically stopped. Unit: seconds. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	AutoTerminateDelay *int64 `json:"AutoTerminateDelay,omitempty" xml:"AutoTerminateDelay,omitempty"`
	// Specifies whether to automatically stop the task when the latency exceeds the specified limit. Default value: false.
	//
	// example:
	//
	// true
	AutoTerminateEnabled *bool `json:"AutoTerminateEnabled,omitempty" xml:"AutoTerminateEnabled,omitempty"`
	// The callback URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://xxx.com
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The ID of the channel.
	//
	// >  This parameter is required and takes effect only if you set the Mode parameter to rtc.
	//
	// example:
	//
	// channelId
	ChannelID *string `json:"ChannelID,omitempty" xml:"ChannelID,omitempty"`
	// The source language of the audio. Valid values:
	//
	// 	- ja: Japanese
	//
	// 	- yue: Cantonese
	//
	// 	- fspk: mixed Mandarin and English
	//
	// 	- en: English
	//
	// 	- cn: Mandarin
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The type of the stream. Valid values: live and rtc. The value live specifies a regular live stream, such as a Real-Time Messaging Protocol (RTMP) stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// live
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The interval at which callbacks are returned. Unit: milliseconds. Valid values: -1 and 0 to 500.
	//
	// 	- \\-1: accepts callbacks only for whole sentences, but not for incomplete sentences.
	//
	// 	- 0 or an empty value: returns callbacks in real time.
	//
	// 	- A value that is greater than 0 and less than or equal to 500: returns callbacks at the specified interval.
	//
	// example:
	//
	// 5
	ReportInterval *int64 `json:"ReportInterval,omitempty" xml:"ReportInterval,omitempty"`
	// The ID of the user who ingests the stream.
	//
	// >  This parameter is required and takes effect only if you set the Mode parameter to rtc. You can specify only one user ID.
	//
	// example:
	//
	// user1
	RtcUserId *string `json:"RtcUserId,omitempty" xml:"RtcUserId,omitempty"`
	// The ID of the ApsaraVideo Real-time Communication (ARTC) application.
	//
	// >  This parameter is required and takes effect only if you set the Mode parameter to rtc.
	//
	// example:
	//
	// appId
	SDKAppID *string `json:"SDKAppID,omitempty" xml:"SDKAppID,omitempty"`
	// The URL of the live stream.
	//
	// >  This parameter is required and takes effect only if you set the Mode parameter to live.
	//
	// example:
	//
	// rtmp://xxx
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The language into which the subtitles are translated. Valid values:
	//
	// 	- cn: Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// en
	TargetLanguages *string `json:"TargetLanguages,omitempty" xml:"TargetLanguages,omitempty"`
	// Specifies whether to enable the translation feature.
	//
	// example:
	//
	// true
	TranslateEnabled *bool `json:"TranslateEnabled,omitempty" xml:"TranslateEnabled,omitempty"`
}

func (s CreateRtcAsrTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRtcAsrTaskRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateRtcAsrTaskRequest) GetAutoTerminateDelay() *int64 {
	return s.AutoTerminateDelay
}

func (s *CreateRtcAsrTaskRequest) GetAutoTerminateEnabled() *bool {
	return s.AutoTerminateEnabled
}

func (s *CreateRtcAsrTaskRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *CreateRtcAsrTaskRequest) GetChannelID() *string {
	return s.ChannelID
}

func (s *CreateRtcAsrTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateRtcAsrTaskRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateRtcAsrTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRtcAsrTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRtcAsrTaskRequest) GetReportInterval() *int64 {
	return s.ReportInterval
}

func (s *CreateRtcAsrTaskRequest) GetRtcUserId() *string {
	return s.RtcUserId
}

func (s *CreateRtcAsrTaskRequest) GetSDKAppID() *string {
	return s.SDKAppID
}

func (s *CreateRtcAsrTaskRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *CreateRtcAsrTaskRequest) GetTargetLanguages() *string {
	return s.TargetLanguages
}

func (s *CreateRtcAsrTaskRequest) GetTranslateEnabled() *bool {
	return s.TranslateEnabled
}

func (s *CreateRtcAsrTaskRequest) SetAuthKey(v string) *CreateRtcAsrTaskRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetAutoTerminateDelay(v int64) *CreateRtcAsrTaskRequest {
	s.AutoTerminateDelay = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetAutoTerminateEnabled(v bool) *CreateRtcAsrTaskRequest {
	s.AutoTerminateEnabled = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetCallbackURL(v string) *CreateRtcAsrTaskRequest {
	s.CallbackURL = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetChannelID(v string) *CreateRtcAsrTaskRequest {
	s.ChannelID = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetLanguage(v string) *CreateRtcAsrTaskRequest {
	s.Language = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetMode(v string) *CreateRtcAsrTaskRequest {
	s.Mode = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetOwnerId(v int64) *CreateRtcAsrTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetRegionId(v string) *CreateRtcAsrTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetReportInterval(v int64) *CreateRtcAsrTaskRequest {
	s.ReportInterval = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetRtcUserId(v string) *CreateRtcAsrTaskRequest {
	s.RtcUserId = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetSDKAppID(v string) *CreateRtcAsrTaskRequest {
	s.SDKAppID = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetStreamURL(v string) *CreateRtcAsrTaskRequest {
	s.StreamURL = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetTargetLanguages(v string) *CreateRtcAsrTaskRequest {
	s.TargetLanguages = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) SetTranslateEnabled(v bool) *CreateRtcAsrTaskRequest {
	s.TranslateEnabled = &v
	return s
}

func (s *CreateRtcAsrTaskRequest) Validate() error {
	return dara.Validate(s)
}
