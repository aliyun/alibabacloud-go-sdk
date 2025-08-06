// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIMediaAuditJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditConfiguration(v string) *SubmitAIMediaAuditJobRequest
	GetMediaAuditConfiguration() *string
	SetMediaId(v string) *SubmitAIMediaAuditJobRequest
	GetMediaId() *string
	SetMediaType(v string) *SubmitAIMediaAuditJobRequest
	GetMediaType() *string
	SetTemplateId(v string) *SubmitAIMediaAuditJobRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitAIMediaAuditJobRequest
	GetUserData() *string
}

type SubmitAIMediaAuditJobRequest struct {
	// The configuration information about the review job.
	//
	// 	- Other configuration items of the review job. Only the ResourceType field is supported. This field is used to specify the type of media files. You can adjust review standards and rules based on the type of media files.
	//
	// 	- If you want to modify the review standard and rules based on ResourceType, submit a ticket. For more information, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// 	- The value of ResourceType can contain only letters, digits, and underscores (_).
	//
	// example:
	//
	// {"ResourceType":"****_movie"}
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	// The ID of the video file. To obtain the file ID, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Review Management*	- > **Content Moderation*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe028d09441afffb138cd7ee****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media file. Only **video*	- is supported.
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The ID of the AI template. You can use one of the following methods to obtain the ID of the AI template:
	//
	// 	- Obtain the value of TemplateId from the response to the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation that you call to create an AI template.
	//
	// 	- Obtain the value of TemplateId from the response to the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation that you call to create an AI template.
	//
	// >  If you do not specify an ID, the ID of the default AI template is used.
	//
	// example:
	//
	// a07a7f7d7d10eb9fd999e56ecc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The custom settings. The value must be a JSON string. You can configure settings such as message callbacks. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// >  To use the callback configurations specified by this parameter, you must configure an HTTP callback URL and specify the types of the callback events in the ApsaraVideo VOD console. Otherwise, the callback configurations do not take effect. For more information about how to configure HTTP callback settings in the ApsaraVideo VOD console, see [Configure callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://test.test.com"},"Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIMediaAuditJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIMediaAuditJobRequest) GoString() string {
	return s.String()
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

func (s *SubmitAIMediaAuditJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitAIMediaAuditJobRequest) GetUserData() *string {
	return s.UserData
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

func (s *SubmitAIMediaAuditJobRequest) SetTemplateId(v string) *SubmitAIMediaAuditJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetUserData(v string) *SubmitAIMediaAuditJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) Validate() error {
	return dara.Validate(s)
}
