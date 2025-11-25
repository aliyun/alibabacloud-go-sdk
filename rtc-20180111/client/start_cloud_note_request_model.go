// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartCloudNoteRequest
	GetAppId() *string
	SetAutoChapters(v *StartCloudNoteRequestAutoChapters) *StartCloudNoteRequest
	GetAutoChapters() *StartCloudNoteRequestAutoChapters
	SetChannelId(v string) *StartCloudNoteRequest
	GetChannelId() *string
	SetCustomPrompt(v *StartCloudNoteRequestCustomPrompt) *StartCloudNoteRequest
	GetCustomPrompt() *StartCloudNoteRequestCustomPrompt
	SetLanguageHints(v []*string) *StartCloudNoteRequest
	GetLanguageHints() []*string
	SetMeetingAssistance(v *StartCloudNoteRequestMeetingAssistance) *StartCloudNoteRequest
	GetMeetingAssistance() *StartCloudNoteRequestMeetingAssistance
	SetRealtimeSubtitle(v *StartCloudNoteRequestRealtimeSubtitle) *StartCloudNoteRequest
	GetRealtimeSubtitle() *StartCloudNoteRequestRealtimeSubtitle
	SetServiceInspection(v *StartCloudNoteRequestServiceInspection) *StartCloudNoteRequest
	GetServiceInspection() *StartCloudNoteRequestServiceInspection
	SetSourceLanguage(v string) *StartCloudNoteRequest
	GetSourceLanguage() *string
	SetStorageConfig(v *StartCloudNoteRequestStorageConfig) *StartCloudNoteRequest
	GetStorageConfig() *StartCloudNoteRequestStorageConfig
	SetSummarization(v *StartCloudNoteRequestSummarization) *StartCloudNoteRequest
	GetSummarization() *StartCloudNoteRequestSummarization
	SetTaskId(v string) *StartCloudNoteRequest
	GetTaskId() *string
	SetTextPolish(v *StartCloudNoteRequestTextPolish) *StartCloudNoteRequest
	GetTextPolish() *StartCloudNoteRequestTextPolish
	SetTranscription(v *StartCloudNoteRequestTranscription) *StartCloudNoteRequest
	GetTranscription() *StartCloudNoteRequestTranscription
}

type StartCloudNoteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2ws***z3
	AppId        *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoChapters *StartCloudNoteRequestAutoChapters `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test
	ChannelId         *string                                 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CustomPrompt      *StartCloudNoteRequestCustomPrompt      `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty" type:"Struct"`
	LanguageHints     []*string                               `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	MeetingAssistance *StartCloudNoteRequestMeetingAssistance `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty" type:"Struct"`
	RealtimeSubtitle  *StartCloudNoteRequestRealtimeSubtitle  `json:"RealtimeSubtitle,omitempty" xml:"RealtimeSubtitle,omitempty" type:"Struct"`
	ServiceInspection *StartCloudNoteRequestServiceInspection `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty" type:"Struct"`
	// example:
	//
	// cn
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	StorageConfig *StartCloudNoteRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
	Summarization *StartCloudNoteRequestSummarization `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// rtc
	TaskId        *string                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TextPolish    *StartCloudNoteRequestTextPolish    `json:"TextPolish,omitempty" xml:"TextPolish,omitempty" type:"Struct"`
	Transcription *StartCloudNoteRequestTranscription `json:"Transcription,omitempty" xml:"Transcription,omitempty" type:"Struct"`
}

func (s StartCloudNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequest) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCloudNoteRequest) GetAutoChapters() *StartCloudNoteRequestAutoChapters {
	return s.AutoChapters
}

func (s *StartCloudNoteRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartCloudNoteRequest) GetCustomPrompt() *StartCloudNoteRequestCustomPrompt {
	return s.CustomPrompt
}

func (s *StartCloudNoteRequest) GetLanguageHints() []*string {
	return s.LanguageHints
}

func (s *StartCloudNoteRequest) GetMeetingAssistance() *StartCloudNoteRequestMeetingAssistance {
	return s.MeetingAssistance
}

func (s *StartCloudNoteRequest) GetRealtimeSubtitle() *StartCloudNoteRequestRealtimeSubtitle {
	return s.RealtimeSubtitle
}

func (s *StartCloudNoteRequest) GetServiceInspection() *StartCloudNoteRequestServiceInspection {
	return s.ServiceInspection
}

func (s *StartCloudNoteRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *StartCloudNoteRequest) GetStorageConfig() *StartCloudNoteRequestStorageConfig {
	return s.StorageConfig
}

func (s *StartCloudNoteRequest) GetSummarization() *StartCloudNoteRequestSummarization {
	return s.Summarization
}

func (s *StartCloudNoteRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudNoteRequest) GetTextPolish() *StartCloudNoteRequestTextPolish {
	return s.TextPolish
}

func (s *StartCloudNoteRequest) GetTranscription() *StartCloudNoteRequestTranscription {
	return s.Transcription
}

func (s *StartCloudNoteRequest) SetAppId(v string) *StartCloudNoteRequest {
	s.AppId = &v
	return s
}

func (s *StartCloudNoteRequest) SetAutoChapters(v *StartCloudNoteRequestAutoChapters) *StartCloudNoteRequest {
	s.AutoChapters = v
	return s
}

func (s *StartCloudNoteRequest) SetChannelId(v string) *StartCloudNoteRequest {
	s.ChannelId = &v
	return s
}

func (s *StartCloudNoteRequest) SetCustomPrompt(v *StartCloudNoteRequestCustomPrompt) *StartCloudNoteRequest {
	s.CustomPrompt = v
	return s
}

func (s *StartCloudNoteRequest) SetLanguageHints(v []*string) *StartCloudNoteRequest {
	s.LanguageHints = v
	return s
}

func (s *StartCloudNoteRequest) SetMeetingAssistance(v *StartCloudNoteRequestMeetingAssistance) *StartCloudNoteRequest {
	s.MeetingAssistance = v
	return s
}

func (s *StartCloudNoteRequest) SetRealtimeSubtitle(v *StartCloudNoteRequestRealtimeSubtitle) *StartCloudNoteRequest {
	s.RealtimeSubtitle = v
	return s
}

func (s *StartCloudNoteRequest) SetServiceInspection(v *StartCloudNoteRequestServiceInspection) *StartCloudNoteRequest {
	s.ServiceInspection = v
	return s
}

func (s *StartCloudNoteRequest) SetSourceLanguage(v string) *StartCloudNoteRequest {
	s.SourceLanguage = &v
	return s
}

func (s *StartCloudNoteRequest) SetStorageConfig(v *StartCloudNoteRequestStorageConfig) *StartCloudNoteRequest {
	s.StorageConfig = v
	return s
}

func (s *StartCloudNoteRequest) SetSummarization(v *StartCloudNoteRequestSummarization) *StartCloudNoteRequest {
	s.Summarization = v
	return s
}

func (s *StartCloudNoteRequest) SetTaskId(v string) *StartCloudNoteRequest {
	s.TaskId = &v
	return s
}

func (s *StartCloudNoteRequest) SetTextPolish(v *StartCloudNoteRequestTextPolish) *StartCloudNoteRequest {
	s.TextPolish = v
	return s
}

func (s *StartCloudNoteRequest) SetTranscription(v *StartCloudNoteRequestTranscription) *StartCloudNoteRequest {
	s.Transcription = v
	return s
}

func (s *StartCloudNoteRequest) Validate() error {
	if s.AutoChapters != nil {
		if err := s.AutoChapters.Validate(); err != nil {
			return err
		}
	}
	if s.CustomPrompt != nil {
		if err := s.CustomPrompt.Validate(); err != nil {
			return err
		}
	}
	if s.MeetingAssistance != nil {
		if err := s.MeetingAssistance.Validate(); err != nil {
			return err
		}
	}
	if s.RealtimeSubtitle != nil {
		if err := s.RealtimeSubtitle.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceInspection != nil {
		if err := s.ServiceInspection.Validate(); err != nil {
			return err
		}
	}
	if s.StorageConfig != nil {
		if err := s.StorageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Summarization != nil {
		if err := s.Summarization.Validate(); err != nil {
			return err
		}
	}
	if s.TextPolish != nil {
		if err := s.TextPolish.Validate(); err != nil {
			return err
		}
	}
	if s.Transcription != nil {
		if err := s.Transcription.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCloudNoteRequestAutoChapters struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s StartCloudNoteRequestAutoChapters) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestAutoChapters) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestAutoChapters) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestAutoChapters) SetEnabled(v bool) *StartCloudNoteRequestAutoChapters {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestAutoChapters) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestCustomPrompt struct {
	// This parameter is required.
	CustomPromptContents []*StartCloudNoteRequestCustomPromptCustomPromptContents `json:"CustomPromptContents,omitempty" xml:"CustomPromptContents,omitempty" type:"Repeated"`
	Enabled              *bool                                                    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s StartCloudNoteRequestCustomPrompt) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestCustomPrompt) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestCustomPrompt) GetCustomPromptContents() []*StartCloudNoteRequestCustomPromptCustomPromptContents {
	return s.CustomPromptContents
}

func (s *StartCloudNoteRequestCustomPrompt) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestCustomPrompt) SetCustomPromptContents(v []*StartCloudNoteRequestCustomPromptCustomPromptContents) *StartCloudNoteRequestCustomPrompt {
	s.CustomPromptContents = v
	return s
}

func (s *StartCloudNoteRequestCustomPrompt) SetEnabled(v bool) *StartCloudNoteRequestCustomPrompt {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestCustomPrompt) Validate() error {
	if s.CustomPromptContents != nil {
		for _, item := range s.CustomPromptContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartCloudNoteRequestCustomPromptCustomPromptContents struct {
	// example:
	//
	// tingwu-turbo
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// split-summary-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 请帮我将下面的对话进行总结，根据发言人来总结:\n {Transcription}
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// chat
	TransType *string `json:"TransType,omitempty" xml:"TransType,omitempty"`
}

func (s StartCloudNoteRequestCustomPromptCustomPromptContents) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestCustomPromptCustomPromptContents) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) GetModel() *string {
	return s.Model
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) GetName() *string {
	return s.Name
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) GetPrompt() *string {
	return s.Prompt
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) GetTransType() *string {
	return s.TransType
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) SetModel(v string) *StartCloudNoteRequestCustomPromptCustomPromptContents {
	s.Model = &v
	return s
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) SetName(v string) *StartCloudNoteRequestCustomPromptCustomPromptContents {
	s.Name = &v
	return s
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) SetPrompt(v string) *StartCloudNoteRequestCustomPromptCustomPromptContents {
	s.Prompt = &v
	return s
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) SetTransType(v string) *StartCloudNoteRequestCustomPromptCustomPromptContents {
	s.TransType = &v
	return s
}

func (s *StartCloudNoteRequestCustomPromptCustomPromptContents) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestMeetingAssistance struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	MeetingAssistanceType []*string `json:"MeetingAssistanceType,omitempty" xml:"MeetingAssistanceType,omitempty" type:"Repeated"`
}

func (s StartCloudNoteRequestMeetingAssistance) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestMeetingAssistance) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestMeetingAssistance) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestMeetingAssistance) GetMeetingAssistanceType() []*string {
	return s.MeetingAssistanceType
}

func (s *StartCloudNoteRequestMeetingAssistance) SetEnabled(v bool) *StartCloudNoteRequestMeetingAssistance {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestMeetingAssistance) SetMeetingAssistanceType(v []*string) *StartCloudNoteRequestMeetingAssistance {
	s.MeetingAssistanceType = v
	return s
}

func (s *StartCloudNoteRequestMeetingAssistance) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestRealtimeSubtitle struct {
	AsrCallback *bool                                             `json:"AsrCallback,omitempty" xml:"AsrCallback,omitempty"`
	Enabled     *bool                                             `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Translation *StartCloudNoteRequestRealtimeSubtitleTranslation `json:"Translation,omitempty" xml:"Translation,omitempty" type:"Struct"`
}

func (s StartCloudNoteRequestRealtimeSubtitle) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestRealtimeSubtitle) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestRealtimeSubtitle) GetAsrCallback() *bool {
	return s.AsrCallback
}

func (s *StartCloudNoteRequestRealtimeSubtitle) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestRealtimeSubtitle) GetTranslation() *StartCloudNoteRequestRealtimeSubtitleTranslation {
	return s.Translation
}

func (s *StartCloudNoteRequestRealtimeSubtitle) SetAsrCallback(v bool) *StartCloudNoteRequestRealtimeSubtitle {
	s.AsrCallback = &v
	return s
}

func (s *StartCloudNoteRequestRealtimeSubtitle) SetEnabled(v bool) *StartCloudNoteRequestRealtimeSubtitle {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestRealtimeSubtitle) SetTranslation(v *StartCloudNoteRequestRealtimeSubtitleTranslation) *StartCloudNoteRequestRealtimeSubtitle {
	s.Translation = v
	return s
}

func (s *StartCloudNoteRequestRealtimeSubtitle) Validate() error {
	if s.Translation != nil {
		if err := s.Translation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCloudNoteRequestRealtimeSubtitleTranslation struct {
	// example:
	//
	// 1
	TranslateLevel *int32 `json:"TranslateLevel,omitempty" xml:"TranslateLevel,omitempty"`
}

func (s StartCloudNoteRequestRealtimeSubtitleTranslation) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestRealtimeSubtitleTranslation) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestRealtimeSubtitleTranslation) GetTranslateLevel() *int32 {
	return s.TranslateLevel
}

func (s *StartCloudNoteRequestRealtimeSubtitleTranslation) SetTranslateLevel(v int32) *StartCloudNoteRequestRealtimeSubtitleTranslation {
	s.TranslateLevel = &v
	return s
}

func (s *StartCloudNoteRequestRealtimeSubtitleTranslation) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestServiceInspection struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	InspectionContents []*StartCloudNoteRequestServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 请监测对话中销售人员表现是否接待热情、态度良好
	InspectionIntroduction *string `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 汽车门店线下销售场景
	SceneIntroduction *string `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
}

func (s StartCloudNoteRequestServiceInspection) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestServiceInspection) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestServiceInspection) GetInspectionContents() []*StartCloudNoteRequestServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *StartCloudNoteRequestServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *StartCloudNoteRequestServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *StartCloudNoteRequestServiceInspection) SetEnabled(v bool) *StartCloudNoteRequestServiceInspection {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestServiceInspection) SetInspectionContents(v []*StartCloudNoteRequestServiceInspectionInspectionContents) *StartCloudNoteRequestServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *StartCloudNoteRequestServiceInspection) SetInspectionIntroduction(v string) *StartCloudNoteRequestServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *StartCloudNoteRequestServiceInspection) SetSceneIntroduction(v string) *StartCloudNoteRequestServiceInspection {
	s.SceneIntroduction = &v
	return s
}

func (s *StartCloudNoteRequestServiceInspection) Validate() error {
	if s.InspectionContents != nil {
		for _, item := range s.InspectionContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartCloudNoteRequestServiceInspectionInspectionContents struct {
	// This parameter is required.
	//
	// example:
	//
	// 销售在开场白的时候主动向客户打招呼进行欢迎
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 到店欢迎-欢迎语
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s StartCloudNoteRequestServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *StartCloudNoteRequestServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *StartCloudNoteRequestServiceInspectionInspectionContents) SetContent(v string) *StartCloudNoteRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *StartCloudNoteRequestServiceInspectionInspectionContents) SetTitle(v string) *StartCloudNoteRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *StartCloudNoteRequestServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestStorageConfig struct {
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

func (s StartCloudNoteRequestStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestStorageConfig) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestStorageConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *StartCloudNoteRequestStorageConfig) GetBucket() *string {
	return s.Bucket
}

func (s *StartCloudNoteRequestStorageConfig) GetRegion() *int32 {
	return s.Region
}

func (s *StartCloudNoteRequestStorageConfig) GetSecretKey() *string {
	return s.SecretKey
}

func (s *StartCloudNoteRequestStorageConfig) GetVendor() *int32 {
	return s.Vendor
}

func (s *StartCloudNoteRequestStorageConfig) SetAccessKey(v string) *StartCloudNoteRequestStorageConfig {
	s.AccessKey = &v
	return s
}

func (s *StartCloudNoteRequestStorageConfig) SetBucket(v string) *StartCloudNoteRequestStorageConfig {
	s.Bucket = &v
	return s
}

func (s *StartCloudNoteRequestStorageConfig) SetRegion(v int32) *StartCloudNoteRequestStorageConfig {
	s.Region = &v
	return s
}

func (s *StartCloudNoteRequestStorageConfig) SetSecretKey(v string) *StartCloudNoteRequestStorageConfig {
	s.SecretKey = &v
	return s
}

func (s *StartCloudNoteRequestStorageConfig) SetVendor(v int32) *StartCloudNoteRequestStorageConfig {
	s.Vendor = &v
	return s
}

func (s *StartCloudNoteRequestStorageConfig) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestSummarization struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Type []*string `json:"Type,omitempty" xml:"Type,omitempty" type:"Repeated"`
}

func (s StartCloudNoteRequestSummarization) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestSummarization) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestSummarization) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestSummarization) GetType() []*string {
	return s.Type
}

func (s *StartCloudNoteRequestSummarization) SetEnabled(v bool) *StartCloudNoteRequestSummarization {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestSummarization) SetType(v []*string) *StartCloudNoteRequestSummarization {
	s.Type = v
	return s
}

func (s *StartCloudNoteRequestSummarization) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestTextPolish struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s StartCloudNoteRequestTextPolish) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestTextPolish) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestTextPolish) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartCloudNoteRequestTextPolish) SetEnabled(v bool) *StartCloudNoteRequestTextPolish {
	s.Enabled = &v
	return s
}

func (s *StartCloudNoteRequestTextPolish) Validate() error {
	return dara.Validate(s)
}

type StartCloudNoteRequestTranscription struct {
	DiarizationEnabled *bool `json:"DiarizationEnabled,omitempty" xml:"DiarizationEnabled,omitempty"`
	// example:
	//
	// b401a5da78e94xxxxc3129425c78b6a5
	PhraseId     *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	SpeakerCount *int32  `json:"SpeakerCount,omitempty" xml:"SpeakerCount,omitempty"`
	// example:
	//
	// 1
	TranscriptionLevel *int32 `json:"TranscriptionLevel,omitempty" xml:"TranscriptionLevel,omitempty"`
}

func (s StartCloudNoteRequestTranscription) String() string {
	return dara.Prettify(s)
}

func (s StartCloudNoteRequestTranscription) GoString() string {
	return s.String()
}

func (s *StartCloudNoteRequestTranscription) GetDiarizationEnabled() *bool {
	return s.DiarizationEnabled
}

func (s *StartCloudNoteRequestTranscription) GetPhraseId() *string {
	return s.PhraseId
}

func (s *StartCloudNoteRequestTranscription) GetSpeakerCount() *int32 {
	return s.SpeakerCount
}

func (s *StartCloudNoteRequestTranscription) GetTranscriptionLevel() *int32 {
	return s.TranscriptionLevel
}

func (s *StartCloudNoteRequestTranscription) SetDiarizationEnabled(v bool) *StartCloudNoteRequestTranscription {
	s.DiarizationEnabled = &v
	return s
}

func (s *StartCloudNoteRequestTranscription) SetPhraseId(v string) *StartCloudNoteRequestTranscription {
	s.PhraseId = &v
	return s
}

func (s *StartCloudNoteRequestTranscription) SetSpeakerCount(v int32) *StartCloudNoteRequestTranscription {
	s.SpeakerCount = &v
	return s
}

func (s *StartCloudNoteRequestTranscription) SetTranscriptionLevel(v int32) *StartCloudNoteRequestTranscription {
	s.TranscriptionLevel = &v
	return s
}

func (s *StartCloudNoteRequestTranscription) Validate() error {
	return dara.Validate(s)
}
