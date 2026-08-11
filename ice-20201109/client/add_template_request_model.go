// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *AddTemplateRequest
	GetConfig() *string
	SetCoverUrl(v string) *AddTemplateRequest
	GetCoverUrl() *string
	SetName(v string) *AddTemplateRequest
	GetName() *string
	SetPreviewMedia(v string) *AddTemplateRequest
	GetPreviewMedia() *string
	SetRelatedMediaids(v string) *AddTemplateRequest
	GetRelatedMediaids() *string
	SetSource(v string) *AddTemplateRequest
	GetSource() *string
	SetStatus(v string) *AddTemplateRequest
	GetStatus() *string
	SetType(v string) *AddTemplateRequest
	GetType() *string
}

type AddTemplateRequest struct {
	// - The standard template Config is an encapsulation based on the cloud editing Timeline. Custom template Config provides more flexibility. If you have special requirements, familiarize yourself with the Config structure and customize the template. For details, see [Standard template Config details](https://help.aliyun.com/document_detail/456193.html).
	//
	// - For more template Config examples, see [Common standard template Config examples](https://help.aliyun.com/document_detail/451634.html).
	//
	// - For one-click video production template Config configurations, see [One-click video production template configuration parameter description](https://help.aliyun.com/document_detail/2878274.html).
	//
	// example:
	//
	// See the Timeline template Config documentation
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The template cover URL.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/cover.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The custom template name.
	//
	// example:
	//
	// Template name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The media asset ID of the template preview video.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// The media assets associated with the template, used by the standard template editor.
	//
	// example:
	//
	// {"video":["1805a0c6ca544fb395a06ca683619655"]}
	RelatedMediaids *string `json:"RelatedMediaids,omitempty" xml:"RelatedMediaids,omitempty"`
	// The template creation source. Valid values:
	//
	// - OpenAPI: created by using OpenAPI.
	//
	// - AliyunConsole: created by using the Alibaba Cloud Management Console.
	//
	// - WebSDK: created by using WebSDK.
	//
	// example:
	//
	// OpenAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The template status. Valid values:
	//
	// - Available: normal.
	//
	// - Created: created but not yet available.
	//
	// - Uploading: uploading.
	//
	// - Processing: advanced template is being analyzed.
	//
	// - UploadFailed: upload failed.
	//
	// - ProcessFailed: advanced template analysis failed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The templatetype. Valid values:
	//
	// - Timeline (standard template): a template created based on the Timeline of a video clip node (multiple materials on multiple tracks are concatenated in sequence). This templatetype can be used to implement effects such as image-to-video conversion, photo albums, intros and outros, and default watermarks.
	//
	// - VETemplate (advanced template): a template created based on Adobe After Effects (AE) effects. This templatetype can be used to implement advanced media effects with complex animations.
	//
	// - BatchEditing (one-click video production template): supports configurations for sticker watermarks, background music, background images, narration subtitle styles, title subtitle styles, and output resolution parameters. These configurations are automatically applied when you commit a one-click video production node.
	//
	// example:
	//
	// Timeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *AddTemplateRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *AddTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddTemplateRequest) GetPreviewMedia() *string {
	return s.PreviewMedia
}

func (s *AddTemplateRequest) GetRelatedMediaids() *string {
	return s.RelatedMediaids
}

func (s *AddTemplateRequest) GetSource() *string {
	return s.Source
}

func (s *AddTemplateRequest) GetStatus() *string {
	return s.Status
}

func (s *AddTemplateRequest) GetType() *string {
	return s.Type
}

func (s *AddTemplateRequest) SetConfig(v string) *AddTemplateRequest {
	s.Config = &v
	return s
}

func (s *AddTemplateRequest) SetCoverUrl(v string) *AddTemplateRequest {
	s.CoverUrl = &v
	return s
}

func (s *AddTemplateRequest) SetName(v string) *AddTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddTemplateRequest) SetPreviewMedia(v string) *AddTemplateRequest {
	s.PreviewMedia = &v
	return s
}

func (s *AddTemplateRequest) SetRelatedMediaids(v string) *AddTemplateRequest {
	s.RelatedMediaids = &v
	return s
}

func (s *AddTemplateRequest) SetSource(v string) *AddTemplateRequest {
	s.Source = &v
	return s
}

func (s *AddTemplateRequest) SetStatus(v string) *AddTemplateRequest {
	s.Status = &v
	return s
}

func (s *AddTemplateRequest) SetType(v string) *AddTemplateRequest {
	s.Type = &v
	return s
}

func (s *AddTemplateRequest) Validate() error {
	return dara.Validate(s)
}
