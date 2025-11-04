// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *AddTemplateResponseBodyTemplate) *AddTemplateResponseBody
	GetTemplate() *AddTemplateResponseBodyTemplate
}

type AddTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template information.
	Template *AddTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTemplateResponseBody) GetTemplate() *AddTemplateResponseBodyTemplate {
	return s.Template
}

func (s *AddTemplateResponseBody) SetRequestId(v string) *AddTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTemplateResponseBody) SetTemplate(v *AddTemplateResponseBodyTemplate) *AddTemplateResponseBody {
	s.Template = v
	return s
}

func (s *AddTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTemplateResponseBodyTemplate struct {
	// The template configurations.
	//
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
	// The source from which the template was created.
	//
	// Valid values:
	//
	// 	- AliyunConsole
	//
	// 	- WebSDK
	//
	// 	- OpenAPI
	//
	// example:
	//
	// OpenAPI
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The source from which the template was modified.
	//
	// Valid values:
	//
	// 	- AliyunConsole
	//
	// 	- WebSDK
	//
	// 	- OpenAPI
	//
	// example:
	//
	// OpenAPI
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// The template name.
	//
	// example:
	//
	// 视频添加水印模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the preview video.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// The template state.
	//
	// Valid values:
	//
	// 	- UploadFailed: Failed to upload the video.
	//
	// 	- ProcessFailed: Failed to process the advanced template.
	//
	// 	- Available: The template is available.
	//
	// 	- Uploading: The video is being uploaded.
	//
	// 	- Created: The template is created but not ready for use.
	//
	// 	- Processing: The advanced template is being processed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template type.
	//
	// Valid values:
	//
	// 	- Timeline: regular template.
	//
	// 	- VETemplate: advanced template.
	//
	// example:
	//
	// Timeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplate) GetConfig() *string {
	return s.Config
}

func (s *AddTemplateResponseBodyTemplate) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *AddTemplateResponseBodyTemplate) GetCreateSource() *string {
	return s.CreateSource
}

func (s *AddTemplateResponseBodyTemplate) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *AddTemplateResponseBodyTemplate) GetName() *string {
	return s.Name
}

func (s *AddTemplateResponseBodyTemplate) GetPreviewMedia() *string {
	return s.PreviewMedia
}

func (s *AddTemplateResponseBodyTemplate) GetStatus() *string {
	return s.Status
}

func (s *AddTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddTemplateResponseBodyTemplate) GetType() *string {
	return s.Type
}

func (s *AddTemplateResponseBodyTemplate) SetConfig(v string) *AddTemplateResponseBodyTemplate {
	s.Config = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetCoverUrl(v string) *AddTemplateResponseBodyTemplate {
	s.CoverUrl = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetCreateSource(v string) *AddTemplateResponseBodyTemplate {
	s.CreateSource = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetModifiedSource(v string) *AddTemplateResponseBodyTemplate {
	s.ModifiedSource = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetName(v string) *AddTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetPreviewMedia(v string) *AddTemplateResponseBodyTemplate {
	s.PreviewMedia = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetStatus(v string) *AddTemplateResponseBodyTemplate {
	s.Status = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetTemplateId(v string) *AddTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetType(v string) *AddTemplateResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
