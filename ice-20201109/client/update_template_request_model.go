// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateTemplateRequest
	GetConfig() *string
	SetCoverUrl(v string) *UpdateTemplateRequest
	GetCoverUrl() *string
	SetName(v string) *UpdateTemplateRequest
	GetName() *string
	SetPreviewMedia(v string) *UpdateTemplateRequest
	GetPreviewMedia() *string
	SetRelatedMediaids(v string) *UpdateTemplateRequest
	GetRelatedMediaids() *string
	SetSource(v string) *UpdateTemplateRequest
	GetSource() *string
	SetStatus(v string) *UpdateTemplateRequest
	GetStatus() *string
	SetTemplateId(v string) *UpdateTemplateRequest
	GetTemplateId() *string
}

type UpdateTemplateRequest struct {
	// example:
	//
	// 参见模板Config文档
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The URL of the template thumbnail.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/cover.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The name of the online editing template.
	//
	// example:
	//
	// 视频添加水印模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the preview video.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// The IDs of the materials associated with the template for use by the regular template editor.
	//
	// example:
	//
	// {"video":["******c04f1d4a06996144cc1a******","******cb7db64841b159b4f2ea******"],"audio":["******c04f1d4a06996144cc1a******"],"image":["******c04f1d4a06996144cc1a******"]}
	RelatedMediaids *string `json:"RelatedMediaids,omitempty" xml:"RelatedMediaids,omitempty"`
	// The source from which the template is modified. Default value: OpenAPI. Valid values:
	//
	// 	- AliyunConsole
	//
	// 	- OpenAPI
	//
	// 	- WebSDK
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
	// >  After an advanced template is created, it enters the Processing state. In this case, the template is unavailable. The template can be used only when it is in the Available state. The time required for template processing varies based on the size of the template file. Generally, it ranges from 10 seconds to 5 minutes.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the online editing template. You can obtain the template ID in the [Intelligent Media Services (IMS) console](https://ice.console.aliyun.com/production/template/list/common) or the response parameters of the [AddTemplate](https://help.aliyun.com/document_detail/441161.html) operation.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateTemplateRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *UpdateTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTemplateRequest) GetPreviewMedia() *string {
	return s.PreviewMedia
}

func (s *UpdateTemplateRequest) GetRelatedMediaids() *string {
	return s.RelatedMediaids
}

func (s *UpdateTemplateRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateTemplateRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateRequest) SetConfig(v string) *UpdateTemplateRequest {
	s.Config = &v
	return s
}

func (s *UpdateTemplateRequest) SetCoverUrl(v string) *UpdateTemplateRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateTemplateRequest) SetName(v string) *UpdateTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTemplateRequest) SetPreviewMedia(v string) *UpdateTemplateRequest {
	s.PreviewMedia = &v
	return s
}

func (s *UpdateTemplateRequest) SetRelatedMediaids(v string) *UpdateTemplateRequest {
	s.RelatedMediaids = &v
	return s
}

func (s *UpdateTemplateRequest) SetSource(v string) *UpdateTemplateRequest {
	s.Source = &v
	return s
}

func (s *UpdateTemplateRequest) SetStatus(v string) *UpdateTemplateRequest {
	s.Status = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
