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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ws***z3
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The auto chapters feature. When enabled, chapter titles and summaries are generated.
	AutoChaptersShrink *string `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The custom prompt control parameter object.
	CustomPromptShrink *string `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify it.
	LanguageHints []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	// The intelligent meeting minutes feature. When enabled, keywords, key content, and to-do items are generated.
	MeetingAssistanceShrink *string `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty"`
	// The real-time subtitles.
	RealtimeSubtitleShrink *string `json:"RealtimeSubtitle,omitempty" xml:"RealtimeSubtitle,omitempty"`
	// The service inspection feature.
	ServiceInspectionShrink *string `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty"`
	// The source language model used by the channel. Valid values:
	//
	// - cn: Chinese
	//
	// - yue: Cantonese
	//
	// - en: English
	//
	// - multilingual: Multilingual
	//
	// - ja: Japanese
	//
	// - ko: Korean
	//
	// - de: German
	//
	// - fr: French
	//
	// - ru: Russian
	//
	// - es: Spanish
	//
	// - vi: Vietnamese
	//
	// - it: Italian
	//
	// - sv: Swedish
	//
	// - cs: Czech
	//
	// - pl: Polish
	//
	// - th: Thai
	//
	// - fi: Finnish
	//
	// - hi: Hindi
	//
	// - id: Indonesian
	//
	// - pt: Portuguese
	//
	// - ar: Arabic
	//
	// - fil: Filipino
	//
	// - ms: Malay
	//
	// - tr: Turkish
	//
	// - hu: Hungarian
	//
	// - lo: Lao
	//
	// - pt-br: Portuguese (Brazil)
	//
	// - es-ar: Spanish (Argentina)
	//
	// - es-mx: Spanish (Mexico)
	//
	// example:
	//
	// cn
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The storage configuration.
	//
	// This parameter is required.
	StorageConfig *StartCloudNoteShrinkRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
	// The summarization feature. When enabled, full-text summaries and speaker summaries are generated.
	SummarizationShrink *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The text polishing feature.
	TextPolishShrink *string `json:"TextPolish,omitempty" xml:"TextPolish,omitempty"`
	// The transcription feature. Enabled by default.
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
	if s.StorageConfig != nil {
		if err := s.StorageConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCloudNoteShrinkRequestStorageConfig struct {
	// The AccessKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAX***
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The bucket name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-bucket-for-recording
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The storage region.
	//
	// | Region | AliYun OSS |
	//
	// |--------|------------|
	//
	// | 0 | CN_Hangzhou |
	//
	// | 1 | CN_Shanghai |
	//
	// | 2 | CN_Qingdao |
	//
	// | 3 | CN_Beijing |
	//
	// | 4 | CN_Zhangjiakou |
	//
	// | 5 | CN_Huhehaote |
	//
	// | 6 | CN_Shenzhen |
	//
	// | 7 | CN_Hongkong |
	//
	// | 8 | US_West_1 |
	//
	// | 9 | US_East_1 |
	//
	// | 10 | AP_Southeast_1 |
	//
	// | 11 | AP_Southeast_2 |
	//
	// | 12 | AP_Southeast_3 |
	//
	// | 13 | AP_Southeast_5 |
	//
	// | 14 | AP_Northeast_1 |
	//
	// | 15 | AP_South_1 (Offline. Do not use.) |
	//
	// | 16 | EU_Central_1 |
	//
	// | 17 | EU_West_1 |
	//
	// | 18 | EU_East_1(ME_East_1) |
	//
	// | 19 | AP_Southeast_6 |
	//
	// | 20 | CN_Heyuan |
	//
	// | 21 | CN_Guangzhou |
	//
	// | 22 | CN_Chengdu |
	//
	// | 23 | CN_Nanjing |
	//
	// | 24 | CN_Fuzhou |
	//
	// | 25 | CN_Wulanchabu |
	//
	// | 26 | AP_Northeast_2 |
	//
	// | 27 | AP_Southeast_7 |
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// The AccessKey secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// APb6qWYEzKtYxE***
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The storage platform. Valid values:
	//
	// | vendor | Object storage service | Remarks |
	//
	// |--------|---------|---------|
	//
	// | 1 | AliYun OSS | Alibaba Cloud |
	//
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
