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
	SetMixLayoutParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetMixLayoutParamsShrink() *string
	SetMixTranscodeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest
	GetMixTranscodeParamsShrink() *string
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
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// room1024
	ChannelId                *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	MixLayoutParamsShrink    *string `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty"`
	MixTranscodeParamsShrink *string `json:"MixTranscodeParams,omitempty" xml:"MixTranscodeParams,omitempty"`
	// example:
	//
	// http://xxxx/test/mycallback
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// This parameter is required.
	RecordParamsShrink *string `json:"RecordParams,omitempty" xml:"RecordParams,omitempty"`
	// This parameter is required.
	StorageParamsShrink *string `json:"StorageParams,omitempty" xml:"StorageParams,omitempty"`
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

func (s *StartRtcCloudRecordingShrinkRequest) GetMixLayoutParamsShrink() *string {
	return s.MixLayoutParamsShrink
}

func (s *StartRtcCloudRecordingShrinkRequest) GetMixTranscodeParamsShrink() *string {
	return s.MixTranscodeParamsShrink
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

func (s *StartRtcCloudRecordingShrinkRequest) SetMixLayoutParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.MixLayoutParamsShrink = &v
	return s
}

func (s *StartRtcCloudRecordingShrinkRequest) SetMixTranscodeParamsShrink(v string) *StartRtcCloudRecordingShrinkRequest {
	s.MixTranscodeParamsShrink = &v
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
