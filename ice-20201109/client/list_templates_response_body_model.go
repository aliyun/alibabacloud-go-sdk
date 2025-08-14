// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody
	GetTemplates() []*ListTemplatesResponseBodyTemplates
	SetTotalCount(v int32) *ListTemplatesResponseBody
	GetTotalCount() *int32
}

type ListTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried templates.
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTemplates() []*ListTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesResponseBodyTemplates struct {
	// The clip parameters.
	//
	// example:
	//
	// {"Media1":"mediaId","Text1":"text"}
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The template configurations.
	//
	// example:
	//
	// 参考Timeline模板配置详解
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The thumbnail URL.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
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
	// The time when the template was created.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
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
	// The time when the template was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The template name.
	//
	// example:
	//
	// 视频添加水印模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The preview media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// The state of the preview media asset.
	//
	// Valid values:
	//
	// 	- PrepareFail
	//
	// 	- Init
	//
	// 	- Normal
	//
	// 	- Preparing
	//
	// example:
	//
	// Normal
	PreviewMediaStatus *string `json:"PreviewMediaStatus,omitempty" xml:"PreviewMediaStatus,omitempty"`
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
	// ****20b48fb04483915d4f2cd8ac****
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

func (s ListTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *ListTemplatesResponseBodyTemplates) GetConfig() *string {
	return s.Config
}

func (s *ListTemplatesResponseBodyTemplates) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ListTemplatesResponseBodyTemplates) GetCreateSource() *string {
	return s.CreateSource
}

func (s *ListTemplatesResponseBodyTemplates) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTemplatesResponseBodyTemplates) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *ListTemplatesResponseBodyTemplates) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *ListTemplatesResponseBodyTemplates) GetPreviewMedia() *string {
	return s.PreviewMedia
}

func (s *ListTemplatesResponseBodyTemplates) GetPreviewMediaStatus() *string {
	return s.PreviewMediaStatus
}

func (s *ListTemplatesResponseBodyTemplates) GetStatus() *string {
	return s.Status
}

func (s *ListTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplatesResponseBodyTemplates) GetType() *string {
	return s.Type
}

func (s *ListTemplatesResponseBodyTemplates) SetClipsParam(v string) *ListTemplatesResponseBodyTemplates {
	s.ClipsParam = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetConfig(v string) *ListTemplatesResponseBodyTemplates {
	s.Config = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCoverURL(v string) *ListTemplatesResponseBodyTemplates {
	s.CoverURL = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateSource(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateSource = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreationTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreationTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetModifiedSource(v string) *ListTemplatesResponseBodyTemplates {
	s.ModifiedSource = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetModifiedTime(v string) *ListTemplatesResponseBodyTemplates {
	s.ModifiedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetName(v string) *ListTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPreviewMedia(v string) *ListTemplatesResponseBodyTemplates {
	s.PreviewMedia = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPreviewMediaStatus(v string) *ListTemplatesResponseBodyTemplates {
	s.PreviewMediaStatus = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetStatus(v string) *ListTemplatesResponseBodyTemplates {
	s.Status = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetType(v string) *ListTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
