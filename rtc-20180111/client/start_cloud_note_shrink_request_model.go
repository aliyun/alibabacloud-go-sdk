// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudNoteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartCloudNoteShrinkRequest
	GetAppId() *string
	SetAutoChaptersShrink(v string) *StartCloudNoteShrinkRequest
	GetAutoChaptersShrink() *string
	SetChannelId(v string) *StartCloudNoteShrinkRequest
	GetChannelId() *string
	SetCustomPromptShrink(v string) *StartCloudNoteShrinkRequest
	GetCustomPromptShrink() *string
	SetLanguageHints(v []*string) *StartCloudNoteShrinkRequest
	GetLanguageHints() []*string
	SetMeetingAssistanceShrink(v string) *StartCloudNoteShrinkRequest
	GetMeetingAssistanceShrink() *string
	SetRealtimeSubtitleShrink(v string) *StartCloudNoteShrinkRequest
	GetRealtimeSubtitleShrink() *string
	SetServiceInspectionShrink(v string) *StartCloudNoteShrinkRequest
	GetServiceInspectionShrink() *string
	SetSourceLanguage(v string) *StartCloudNoteShrinkRequest
	GetSourceLanguage() *string
	SetStorageConfig(v *StartCloudNoteShrinkRequestStorageConfig) *StartCloudNoteShrinkRequest
	GetStorageConfig() *StartCloudNoteShrinkRequestStorageConfig
	SetSummarizationShrink(v string) *StartCloudNoteShrinkRequest
	GetSummarizationShrink() *string
	SetTaskId(v string) *StartCloudNoteShrinkRequest
	GetTaskId() *string
	SetTextPolishShrink(v string) *StartCloudNoteShrinkRequest
	GetTextPolishShrink() *string
	SetTranscriptionShrink(v string) *StartCloudNoteShrinkRequest
	GetTranscriptionShrink() *string
}

type StartCloudNoteShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2ws***z3
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoChaptersShrink *string `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	ChannelId               *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CustomPromptShrink      *string   `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	LanguageHints           []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	MeetingAssistanceShrink *string   `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty"`
	RealtimeSubtitleShrink  *string   `json:"RealtimeSubtitle,omitempty" xml:"RealtimeSubtitle,omitempty"`
	ServiceInspectionShrink *string   `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty"`
	// example:
	//
	// cn
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	StorageConfig       *StartCloudNoteShrinkRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
	SummarizationShrink *string                                   `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TextPolishShrink    *string `json:"TextPolish,omitempty" xml:"TextPolish,omitempty"`
	TranscriptionShrink *string `json:"Transcription,omitempty" xml:"Transcription,omitempty"`
}

func (s StartCloudNoteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartCloudNoteShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCloudNoteShrinkRequest) GetAutoChaptersShrink() *string {
	return s.AutoChaptersShrink
}

func (s *StartCloudNoteShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartCloudNoteShrinkRequest) GetCustomPromptShrink() *string {
	return s.CustomPromptShrink
}

func (s *StartCloudNoteShrinkRequest) GetLanguageHints() []*string {
	return s.LanguageHints
}

func (s *StartCloudNoteShrinkRequest) GetMeetingAssistanceShrink() *string {
	return s.MeetingAssistanceShrink
}

func (s *StartCloudNoteShrinkRequest) GetRealtimeSubtitleShrink() *string {
	return s.RealtimeSubtitleShrink
}

func (s *StartCloudNoteShrinkRequest) GetServiceInspectionShrink() *string {
	return s.ServiceInspectionShrink
}

func (s *StartCloudNoteShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *StartCloudNoteShrinkRequest) GetStorageConfig() *StartCloudNoteShrinkRequestStorageConfig {
	return s.StorageConfig
}

func (s *StartCloudNoteShrinkRequest) GetSummarizationShrink() *string {
	return s.SummarizationShrink
}

func (s *StartCloudNoteShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudNoteShrinkRequest) GetTextPolishShrink() *string {
	return s.TextPolishShrink
}

func (s *StartCloudNoteShrinkRequest) GetTranscriptionShrink() *string {
	return s.TranscriptionShrink
}

func (s *StartCloudNoteShrinkRequest) SetAppId(v string) *StartCloudNoteShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetAutoChaptersShrink(v string) *StartCloudNoteShrinkRequest {
	s.AutoChaptersShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetChannelId(v string) *StartCloudNoteShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetCustomPromptShrink(v string) *StartCloudNoteShrinkRequest {
	s.CustomPromptShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetLanguageHints(v []*string) *StartCloudNoteShrinkRequest {
	s.LanguageHints = v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetMeetingAssistanceShrink(v string) *StartCloudNoteShrinkRequest {
	s.MeetingAssistanceShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetRealtimeSubtitleShrink(v string) *StartCloudNoteShrinkRequest {
	s.RealtimeSubtitleShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetServiceInspectionShrink(v string) *StartCloudNoteShrinkRequest {
	s.ServiceInspectionShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetSourceLanguage(v string) *StartCloudNoteShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetStorageConfig(v *StartCloudNoteShrinkRequestStorageConfig) *StartCloudNoteShrinkRequest {
	s.StorageConfig = v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetSummarizationShrink(v string) *StartCloudNoteShrinkRequest {
	s.SummarizationShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetTaskId(v string) *StartCloudNoteShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetTextPolishShrink(v string) *StartCloudNoteShrinkRequest {
	s.TextPolishShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) SetTranscriptionShrink(v string) *StartCloudNoteShrinkRequest {
	s.TranscriptionShrink = &v
	return s
}

func (s *StartCloudNoteShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteShrinkRequestStorageConfig struct {
	// accessKey。
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAX***
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-bucket-for-recording
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// secretKey。
	//
	// This parameter is required.
	//
	// example:
	//
	// APb6qWYEzKtYxE***
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s StartCloudNoteShrinkRequestStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteShrinkRequestStorageConfig) GoString() string {
	return s.String()
}

func (s *StartCloudNoteShrinkRequestStorageConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *StartCloudNoteShrinkRequestStorageConfig) GetBucket() *string {
	return s.Bucket
}

func (s *StartCloudNoteShrinkRequestStorageConfig) GetRegion() *int32 {
	return s.Region
}

func (s *StartCloudNoteShrinkRequestStorageConfig) GetSecretKey() *string {
	return s.SecretKey
}

func (s *StartCloudNoteShrinkRequestStorageConfig) GetVendor() *int32 {
	return s.Vendor
}

func (s *StartCloudNoteShrinkRequestStorageConfig) SetAccessKey(v string) *StartCloudNoteShrinkRequestStorageConfig {
	s.AccessKey = &v
	return s
}

func (s *StartCloudNoteShrinkRequestStorageConfig) SetBucket(v string) *StartCloudNoteShrinkRequestStorageConfig {
	s.Bucket = &v
	return s
}

func (s *StartCloudNoteShrinkRequestStorageConfig) SetRegion(v int32) *StartCloudNoteShrinkRequestStorageConfig {
	s.Region = &v
	return s
}

func (s *StartCloudNoteShrinkRequestStorageConfig) SetSecretKey(v string) *StartCloudNoteShrinkRequestStorageConfig {
	s.SecretKey = &v
	return s
}

func (s *StartCloudNoteShrinkRequestStorageConfig) SetVendor(v int32) *StartCloudNoteShrinkRequestStorageConfig {
	s.Vendor = &v
	return s
}

func (s *StartCloudNoteShrinkRequestStorageConfig) Validate() error {
	return dara.Validate(s)
}
