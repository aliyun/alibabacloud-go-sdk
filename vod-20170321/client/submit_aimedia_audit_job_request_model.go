// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIMediaAuditJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCensorProvider(v string) *SubmitAIMediaAuditJobRequest
	GetCensorProvider() *string
	SetMediaAuditConfiguration(v string) *SubmitAIMediaAuditJobRequest
	GetMediaAuditConfiguration() *string
	SetMediaId(v string) *SubmitAIMediaAuditJobRequest
	GetMediaId() *string
	SetMediaType(v string) *SubmitAIMediaAuditJobRequest
	GetMediaType() *string
	SetServiceParameters(v string) *SubmitAIMediaAuditJobRequest
	GetServiceParameters() *string
	SetTemplateId(v string) *SubmitAIMediaAuditJobRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitAIMediaAuditJobRequest
	GetUserData() *string
	SetVideoService(v string) *SubmitAIMediaAuditJobRequest
	GetVideoService() *string
	SetVoiceService(v string) *SubmitAIMediaAuditJobRequest
	GetVoiceService() *string
}

type SubmitAIMediaAuditJobRequest struct {
	CensorProvider *string `json:"CensorProvider,omitempty" xml:"CensorProvider,omitempty"`
	// The configuration of the review job.
	//
	// - For other configuration items of the review job, only the ResourceType field is currently supported. This field controls the media file type, and you can adjust the review standards and rules for the specified type.
	//
	// - To adjust the review standards and rules for a ResourceType, submit a ticket for technical support. For information about how to submit a ticket, refer to [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// - Usage notes for ResourceType: Only letters, digits, and underscores (_) are allowed.
	//
	// example:
	//
	// {"ResourceType":"****_movie"}
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	// The audio or video ID. Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Review Management*	- > **Video Review*	- to view the audio or video ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe028d09441afffb138cd7ee****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The media type. Currently, only **video*	- is supported.
	//
	// example:
	//
	// video
	MediaType         *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	// The AI template ID. You can obtain the ID by using one of the following methods:
	//
	// - When you call the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template, the AI template ID is the value of the TemplateId response parameter.
	//
	// - After the AI template is added, call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the AI template ID, which is the value of the TemplateId response parameter.
	//
	// > If you do not specify an AI template ID, the default AI template ID for automated review is used.
	//
	// example:
	//
	// a07a7f7d7d10eb9fd999e56ecc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks. For more information, refer to [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// > To use message callbacks in this parameter, you must configure an HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. For information about how to configure HTTP callbacks in the console, refer to [Callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://test.test.com"},"Extend":{"localId":"xxx","test":"www"}}
	UserData     *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoService *string `json:"VideoService,omitempty" xml:"VideoService,omitempty"`
	VoiceService *string `json:"VoiceService,omitempty" xml:"VoiceService,omitempty"`
}

func (s SubmitAIMediaAuditJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIMediaAuditJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobRequest) GetCensorProvider() *string {
	return s.CensorProvider
}

func (s *SubmitAIMediaAuditJobRequest) GetMediaAuditConfiguration() *string {
	return s.MediaAuditConfiguration
}

func (s *SubmitAIMediaAuditJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIMediaAuditJobRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SubmitAIMediaAuditJobRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *SubmitAIMediaAuditJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitAIMediaAuditJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIMediaAuditJobRequest) GetVideoService() *string {
	return s.VideoService
}

func (s *SubmitAIMediaAuditJobRequest) GetVoiceService() *string {
	return s.VoiceService
}

func (s *SubmitAIMediaAuditJobRequest) SetCensorProvider(v string) *SubmitAIMediaAuditJobRequest {
	s.CensorProvider = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaAuditConfiguration(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaAuditConfiguration = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaId(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaType(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaType = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetServiceParameters(v string) *SubmitAIMediaAuditJobRequest {
	s.ServiceParameters = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetTemplateId(v string) *SubmitAIMediaAuditJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetUserData(v string) *SubmitAIMediaAuditJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetVideoService(v string) *SubmitAIMediaAuditJobRequest {
	s.VideoService = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetVoiceService(v string) *SubmitAIMediaAuditJobRequest {
	s.VoiceService = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) Validate() error {
	return dara.Validate(s)
}
