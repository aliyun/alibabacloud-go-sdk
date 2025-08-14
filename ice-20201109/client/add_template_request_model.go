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
	// example:
	//
	// 参见Timeline模板Config文档
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The URL of the template thumbnail.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/cover.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The name of the custom template.
	//
	// example:
	//
	// 视频添加水印模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template preview video.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// The IDs of the materials associated with the template for use by the regular template editor.
	//
	// example:
	//
	// {"video":["1805a0c6ca544fb395a06ca683619655"]}
	RelatedMediaids *string `json:"RelatedMediaids,omitempty" xml:"RelatedMediaids,omitempty"`
	// The source from which the template is created. Valid values:
	//
	// 	- OpenAPI
	//
	// 	- AliyunConsole
	//
	// 	- WebSDK
	//
	// <!---->
	//
	// example:
	//
	// OpenAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The template state. Valid values:
	//
	// 	- Available: The template is available.
	//
	// 	- Created: The template is created but not ready for use.
	//
	// 	- Uploading: The video is being uploaded.
	//
	// 	- Processing: The advanced template is being processed.
	//
	// 	- UploadFailed: Failed to upload the video.
	//
	// 	- ProcessFailed: Failed to process the advanced template.
	//
	// <!---->
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type. Valid values:
	//
	// 	- Timeline: a regular template created based on the timeline of a video editing project, in which multiple materials are arranged in sequence across multiple layers. It can be used to convert text and images into videos, create photo albums, add opening and closing parts, and apply the default watermark.
	//
	// 	- VETemplate: an advanced template created using effects of Adobe After Effects (AE). It can be used to produce complex animations and advanced media effects.
	//
	// <!---->
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
