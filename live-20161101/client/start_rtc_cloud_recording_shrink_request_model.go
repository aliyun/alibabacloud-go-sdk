// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudRecordingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRtcCloudRecordingShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *StartRtcCloudRecordingShrinkRequest
	GetChannelId() *string
	SetMaxIdleTime(v int64) *StartRtcCloudRecordingShrinkRequest
	GetMaxIdleTime() *int64
	SetMixLayoutParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetMixLayoutParamsShrink() *string
	SetMixTranscodeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetMixTranscodeParamsShrink() *string
	SetNotifyAuthKey(v string) *StartRtcCloudRecordingShrinkRequest
	GetNotifyAuthKey() *string
	SetNotifyFileUploadedFormat(v []*string) *StartRtcCloudRecordingShrinkRequest
	GetNotifyFileUploadedFormat() []*string
	SetNotifyUrl(v string) *StartRtcCloudRecordingShrinkRequest
	GetNotifyUrl() *string
	SetRecordParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetRecordParamsShrink() *string
	SetStorageParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetStorageParamsShrink() *string
	SetSubscribeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetSubscribeParamsShrink() *string
}

type StartRtcCloudRecordingShrinkRequest struct {
	// The ID of the app to which the channel to be recorded belongs. The app must belong to the primary account associated with the current API caller\\"s account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to be recorded. Make sure that the channel has active users when you call this operation. Otherwise, the recording task fails to be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// room1024
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The idle timeout period. When the task remains idle for longer than MaxIdleTime, the task is automatically stopped. Unit: seconds. The value must be within [10,14400], which is a maximum of 4 hours. Default value: 300.
	//
	// example:
	//
	// 600
	MaxIdleTime *int64 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The layout parameters. This parameter is not required in single-stream recording mode and is required in stream mixing recording mode when the output is not audio-only.
	MixLayoutParamsShrink *string `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty"`
	// The transcoding parameters. This parameter is not required in single-stream recording mode and is required in stream mixing recording mode.
	MixTranscodeParamsShrink *string `json:"MixTranscodeParams,omitempty" xml:"MixTranscodeParams,omitempty"`
	// The authentication key for callback messages. Leave this parameter empty to skip authentication. If specified, the key must be 16 to 64 characters in length and consist of only uppercase and lowercase letters and digits.
	//
	// example:
	//
	// mytestkeymytestkey
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// The specified formats for which a callback message is sent when the recording file upload event (RecordFileUploaded) is triggered.
	NotifyFileUploadedFormat []*string `json:"NotifyFileUploadedFormat,omitempty" xml:"NotifyFileUploadedFormat,omitempty" type:"Repeated"`
	// The URL for receiving callback messages. Task status messages are pushed to this URL in JSON format by using the POST method. The maximum length is 2048 characters.
	//
	// example:
	//
	// http://xxxx/test/mycallback
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The recording parameters.
	//
	// This parameter is required.
	RecordParamsShrink *string `json:"RecordParams,omitempty" xml:"RecordParams,omitempty"`
	// The storage parameters.
	//
	// This parameter is required.
	StorageParamsShrink *string `json:"StorageParams,omitempty" xml:"StorageParams,omitempty"`
	// The subscription parameters.
	//
	// This parameter is required.
	SubscribeParamsShrink *string `json:"SubscribeParams,omitempty" xml:"SubscribeParams,omitempty"`
}

func (s StartRtcCloudRecordingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRtcCloudRecordingShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcCloudRecordingShrinkRequest) GetMaxIdleTime() *int64 {
	return s.MaxIdleTime
}

func (s *StartRtcCloudRecordingShrinkRequest) GetMixLayoutParamsShrink() *string {
	return s.MixLayoutParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) GetMixTranscodeParamsShrink() *string {
	return s.MixTranscodeParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *StartRtcCloudRecordingShrinkRequest) GetNotifyFileUploadedFormat() []*string {
	return s.NotifyFileUploadedFormat
}

func (s *StartRtcCloudRecordingShrinkRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *StartRtcCloudRecordingShrinkRequest) GetRecordParamsShrink() *string {
	return s.RecordParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) GetStorageParamsShrink() *string {
	return s.StorageParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) GetSubscribeParamsShrink() *string {
	return s.SubscribeParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) SetAppId(v string) *StartRtcCloudRecordingShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetChannelId(v string) *StartRtcCloudRecordingShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetMaxIdleTime(v int64) *StartRtcCloudRecordingShrinkRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetMixLayoutParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.MixLayoutParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetMixTranscodeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.MixTranscodeParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetNotifyAuthKey(v string) *StartRtcCloudRecordingShrinkRequest {
	s.NotifyAuthKey = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetNotifyFileUploadedFormat(v []*string) *StartRtcCloudRecordingShrinkRequest {
	s.NotifyFileUploadedFormat = v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetNotifyUrl(v string) *StartRtcCloudRecordingShrinkRequest {
	s.NotifyUrl = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetRecordParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.RecordParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetStorageParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.StorageParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetSubscribeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.SubscribeParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
