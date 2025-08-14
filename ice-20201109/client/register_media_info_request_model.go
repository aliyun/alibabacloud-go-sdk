// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *RegisterMediaInfoRequest
	GetBusinessType() *string
	SetCateId(v int64) *RegisterMediaInfoRequest
	GetCateId() *int64
	SetClientToken(v string) *RegisterMediaInfoRequest
	GetClientToken() *string
	SetCoverURL(v string) *RegisterMediaInfoRequest
	GetCoverURL() *string
	SetDescription(v string) *RegisterMediaInfoRequest
	GetDescription() *string
	SetInputURL(v string) *RegisterMediaInfoRequest
	GetInputURL() *string
	SetMediaTags(v string) *RegisterMediaInfoRequest
	GetMediaTags() *string
	SetMediaType(v string) *RegisterMediaInfoRequest
	GetMediaType() *string
	SetOverwrite(v bool) *RegisterMediaInfoRequest
	GetOverwrite() *bool
	SetReferenceId(v string) *RegisterMediaInfoRequest
	GetReferenceId() *string
	SetRegisterConfig(v string) *RegisterMediaInfoRequest
	GetRegisterConfig() *string
	SetSmartTagTemplateId(v string) *RegisterMediaInfoRequest
	GetSmartTagTemplateId() *string
	SetTitle(v string) *RegisterMediaInfoRequest
	GetTitle() *string
	SetUserData(v string) *RegisterMediaInfoRequest
	GetUserData() *string
	SetWorkflowId(v string) *RegisterMediaInfoRequest
	GetWorkflowId() *string
}

type RegisterMediaInfoRequest struct {
	// The business type of the media asset. Valid values:
	//
	// 	- subtitles
	//
	// 	- watermark
	//
	// 	- opening
	//
	// 	- ending
	//
	// 	- general
	//
	// example:
	//
	// opening
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 3048
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value must be a UUID that contains 32 characters.
	//
	// example:
	//
	// ****0311a423d11a5f7dee713535****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The thumbnail URL of the media asset.
	//
	// 	- The value can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the media asset.
	//
	// 	- The value can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// defaultDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset in another service. The URL is associated with the ID of the media asset in IMS. The URL cannot be modified once registered. The following types of URLs are supported:
	//
	// 	- OSS URL in one of the following formats:
	//
	// http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	//
	// oss://example-bucket/example.mp4: In this format, it is considered by default that the region of the OSS bucket in which the media asset resides is the same as the region in which IMS is activated.
	//
	// 	- URL of an ApsaraVideo VOD media asset
	//
	// vod://\\*\\*\\*20b48fb04483915d4f2cd8ac\\*\\*\\*\\*
	//
	// This parameter is required.
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The tags of the media asset.
	//
	// 	- Up to 16 tags are supported.
	//
	// 	- Separate multiple tags with commas (,).
	//
	// 	- Each tag can be up to 32 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The type of the media asset. Valid values:
	//
	// 	- image
	//
	// 	- video
	//
	// 	- audio
	//
	// 	- text
	//
	// We recommend that you specify this parameter based on your business requirements. If you set InputURL to an OSS URL, the media asset type can be automatically determined based on the file name extension. For more information
	//
	// <props="china">, see [File formats](https://help.aliyun.com/document_detail/466207.html).
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Specifies whether to overwrite the media asset that has been registered by using the same URL. Default value: false. Valid values:
	//
	// \\- true: If a media asset has been registered by using the same URL, the original media asset is deleted and the new media asset is registered.
	//
	// \\- false: If a media asset has been registered by using the same URL, the new media asset is not registered. A URL cannot be used to register multiple media assets.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The custom ID. The ID can be 6 to 64 characters in length and can contain only letters, digits, hyphens (-), and underscores (_). Make sure that the ID is unique among users.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The registration configurations.
	//
	// By default, a sprite is generated for the media asset. You can set NeedSprite to false to disable automatic sprite generation.
	//
	// By default, a snapshot is generated for the media asset. You can set NeedSnapshot to false to disable automatic snapshot generation.
	//
	// example:
	//
	// {
	//
	//       "NeedSprite": "false"
	//
	// }
	RegisterConfig *string `json:"RegisterConfig,omitempty" xml:"RegisterConfig,omitempty"`
	// The ID of the smart tagging template. Valid values:
	//
	// 	- S00000101-300080: the system template that supports natural language processing (NLP) for content recognition.
	//
	// 	- S00000103-000001: the system template that supports NLP for content recognition and all tagging capabilities.
	//
	// 	- S00000103-000002: the system template that supports all tagging capabilities but does not support NLP for content recognition.
	//
	// After you configure this parameter, a smart tag analysis task is automatically initiated after the media asset is registered. For more information about the billable items<props="china">, see [Smart tagging](https://help.aliyun.com/zh/ims/media-ai-billing?spm=a2c4g.11186623.0.0.3147392dWwlSjL#p-k38-3rb-dug).
	//
	// example:
	//
	// S00000101-300080
	SmartTagTemplateId *string `json:"SmartTagTemplateId,omitempty" xml:"SmartTagTemplateId,omitempty"`
	// The title. If you do not specify this parameter, a default title is automatically generated based on the date.
	//
	// 	- The value can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// defaultTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user data. You can specify a custom callback URL. For more information<props="china"> ,see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	//
	// 	- The value can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// 	- The value must be in the JSON format.
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// ******b4fb044839815d4f2cd8******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s RegisterMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *RegisterMediaInfoRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *RegisterMediaInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RegisterMediaInfoRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *RegisterMediaInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *RegisterMediaInfoRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *RegisterMediaInfoRequest) GetMediaTags() *string {
	return s.MediaTags
}

func (s *RegisterMediaInfoRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *RegisterMediaInfoRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *RegisterMediaInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *RegisterMediaInfoRequest) GetRegisterConfig() *string {
	return s.RegisterConfig
}

func (s *RegisterMediaInfoRequest) GetSmartTagTemplateId() *string {
	return s.SmartTagTemplateId
}

func (s *RegisterMediaInfoRequest) GetTitle() *string {
	return s.Title
}

func (s *RegisterMediaInfoRequest) GetUserData() *string {
	return s.UserData
}

func (s *RegisterMediaInfoRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *RegisterMediaInfoRequest) SetBusinessType(v string) *RegisterMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCateId(v int64) *RegisterMediaInfoRequest {
	s.CateId = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetClientToken(v string) *RegisterMediaInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCoverURL(v string) *RegisterMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetDescription(v string) *RegisterMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetInputURL(v string) *RegisterMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaTags(v string) *RegisterMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaType(v string) *RegisterMediaInfoRequest {
	s.MediaType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetOverwrite(v bool) *RegisterMediaInfoRequest {
	s.Overwrite = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetReferenceId(v string) *RegisterMediaInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetRegisterConfig(v string) *RegisterMediaInfoRequest {
	s.RegisterConfig = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetSmartTagTemplateId(v string) *RegisterMediaInfoRequest {
	s.SmartTagTemplateId = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetTitle(v string) *RegisterMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetUserData(v string) *RegisterMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetWorkflowId(v string) *RegisterMediaInfoRequest {
	s.WorkflowId = &v
	return s
}

func (s *RegisterMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
