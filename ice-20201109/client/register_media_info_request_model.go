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
	// - subtitles
	//
	// - font
	//
	// - watermark
	//
	// - opening
	//
	// - ending
	//
	// - general
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
	// The client token. A 32-character UUID that ensures the idempotence of the request.
	//
	// example:
	//
	// ****0311a423d11a5f7dee713535****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cover image URL.
	//
	// - Maximum length: 128 bytes.
	//
	// - UTF-8 encoded.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The content description.
	//
	// - Maximum length: 1024 bytes.
	//
	// - UTF-8 encoded.
	//
	// example:
	//
	// defaultDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset to be registered in the corresponding system. Once registered, this URL cannot be changed and is attached to the IMS mediaId.
	//
	// - OSS URL. Two formats are supported:
	//
	//
	//
	// http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	//
	// oss://example-bucket/example.mp4
	//
	//  (This format assumes by default that the OSS region is the same as the service registration area.)
	//
	// - VOD media asset:
	//
	// vod://\\*\\*\\*20b48fb04483915d4f2cd8ac****
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4  or  vod://****20b48fb04483915d4f2cd8ac****
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The tags.
	//
	// - Maximum number of tags: 16.
	//
	// - Separate multiple tags with commas.
	//
	// - Maximum length of a single tag: 32 bytes.
	//
	// - UTF-8 encoded.
	//
	// example:
	//
	// tag1,tag2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The media type of the media asset. Valid values:
	//
	// - image
	//
	// - video
	//
	// - audio
	//
	// - text
	//
	// When the value is "text", the businessType must be set to "subtitles" or "font".
	//
	// Specify this field as needed. When the InputURL field is an OSS URL, the media type can also be automatically determined based on the file name extension (only for image, video, and audio file extensions). For the mapping between file extensions and media types, see [File formats](https://help.aliyun.com/document_detail/466207.html).
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Specifies whether to overwrite an existing registered media asset. Default value: false.
	//
	// - true: If the inputUrl is already registered, the existing media asset is deleted and a new media asset is registered.
	//
	// - false: If the inputUrl is already registered, the new media asset is not registered. Duplicate inputUrl values are not supported.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The custom ID. Only lowercase letters, uppercase letters, digits, hyphens (-), and underscores (_) are supported. The length must be 6 to 64 characters. The ID must be unique for each user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The registration configuration.
	//
	// By default, a sprite image is generated for the media asset. To disable this, set the NeedSprite field to false.
	//
	// By default, a snapshot is generated. To disable this, set the NeedSnapshot field to false.
	//
	// To specify the time for the cover image, configure CoverConfig, which contains the following field:
	//
	// - StartTime: The time in seconds at which the cover image is captured from the media asset. Up to four decimal places are supported.
	//
	// After media asset registration, to import the media asset into a custom search library, configure SearchLibName. For information about how to create and use a custom search library, see [Use a custom search library](~~2796619#dd34d8c740yj9~~).
	//
	// example:
	//
	// {
	//
	// 	"NeedSprite": "false",
	//
	// 	"CoverConfig": {
	//
	// 		"StartTime": 1.0
	//
	// 	},
	//
	//        "SearchLibName": "test"
	//
	// }
	RegisterConfig *string `json:"RegisterConfig,omitempty" xml:"RegisterConfig,omitempty"`
	// The intelligent tagging template. Valid values:
	//
	// - S00000101-300080: A system template that includes NLP content understanding.
	//
	// - S00000103-000001: A system template that includes NLP content understanding and all [tagging capabilities](~~2804526#93b27f536airj~~).
	//
	// - S00000103-000002: A system template that includes all [tagging capabilities](~~2804526#93b27f536airj~~) but does not include NLP content understanding.
	//
	// For more information about tagging capabilities, see the documentation.
	//
	// After this field is configured, an intelligent tagging analysis task is automatically initiated upon media asset registration. For billing information, see [Billing of Smart Tag Standard Edition](https://help.aliyun.com/document_detail/600262.html).
	//
	// example:
	//
	// S00000101-300080
	SmartTagTemplateId *string `json:"SmartTagTemplateId,omitempty" xml:"SmartTagTemplateId,omitempty"`
	// The title. If not provided, a default title is automatically generated based on the date.
	//
	// - Maximum length: 128 bytes.
	//
	// - UTF-8 encoded.
	//
	// example:
	//
	// defaultTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user data. Custom callback URL configuration is supported. For configuration instructions, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	//
	// - Maximum length: 1024 bytes.
	//
	// - UTF-8 encoded.
	//
	// - Json format.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"} or{"NotifyAddress":"https://xx.xx.xxx"} or{"NotifyAddress":"ice-callback-demo"}
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
