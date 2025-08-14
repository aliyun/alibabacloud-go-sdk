// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody
	GetTemplate() *GetTemplateResponseBodyTemplate
}

type GetTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template information.
	Template *GetTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetTemplate() *GetTemplateResponseBodyTemplate {
	return s.Template
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateResponseBodyTemplate struct {
	// The clip parameters for submitting a video production job. You can replace mediaId and text with real values to submit a job. References:
	//
	// 	- [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html)
	//
	// 	- [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html)
	//
	// example:
	//
	// {"Media1":"mediaId","Text1":"text"}
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The template configurations.
	//
	// 	- For more information about the configurations of a regular template, see [Config object of a regular template](https://help.aliyun.com/document_detail/456193.html).
	//
	// 	- For more information about the configurations of an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
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
	// The source from which the template was created. Valid values:
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
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The source from which the template was modified. Valid values:
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
	// The state of the preview media asset. Valid values:
	//
	// 	- Init: the initial state, which indicates that the source file is not ready.
	//
	// 	- Preparing: The source file is being prepared. For example, the file is being uploaded or edited.
	//
	// 	- PrepareFail: The source file failed to be prepared. For example, the information about the source file failed to be obtained.
	//
	// 	- Normal: The source file is ready.
	//
	// example:
	//
	// Normal
	PreviewMediaStatus *string `json:"PreviewMediaStatus,omitempty" xml:"PreviewMediaStatus,omitempty"`
	// The IDs of the materials associated with the template for use by the regular template editor.
	//
	// example:
	//
	// {"video":["******c04f1d4a06996144cc1a******"],"audio":["******c04f1d4a06996144cc1a******"],"image":["******c04f1d4a06996144cc1a******"]}
	RelatedMediaids *string `json:"RelatedMediaids,omitempty" xml:"RelatedMediaids,omitempty"`
	// The template state. Valid values:
	//
	// 	- Available
	//
	// 	- Created
	//
	// 	- Uploading
	//
	// 	- Processing
	//
	// 	- UploadFailed
	//
	// 	- ProcessFailed
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
	// The template type. Valid values:
	//
	// 	- Timeline
	//
	// 	- VETemplate
	//
	// example:
	//
	// Timeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTemplate) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *GetTemplateResponseBodyTemplate) GetConfig() *string {
	return s.Config
}

func (s *GetTemplateResponseBodyTemplate) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetTemplateResponseBodyTemplate) GetCreateSource() *string {
	return s.CreateSource
}

func (s *GetTemplateResponseBodyTemplate) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTemplateResponseBodyTemplate) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *GetTemplateResponseBodyTemplate) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetTemplateResponseBodyTemplate) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBodyTemplate) GetPreviewMedia() *string {
	return s.PreviewMedia
}

func (s *GetTemplateResponseBodyTemplate) GetPreviewMediaStatus() *string {
	return s.PreviewMediaStatus
}

func (s *GetTemplateResponseBodyTemplate) GetRelatedMediaids() *string {
	return s.RelatedMediaids
}

func (s *GetTemplateResponseBodyTemplate) GetStatus() *string {
	return s.Status
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateResponseBodyTemplate) GetType() *string {
	return s.Type
}

func (s *GetTemplateResponseBodyTemplate) SetClipsParam(v string) *GetTemplateResponseBodyTemplate {
	s.ClipsParam = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetConfig(v string) *GetTemplateResponseBodyTemplate {
	s.Config = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCoverURL(v string) *GetTemplateResponseBodyTemplate {
	s.CoverURL = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreateSource(v string) *GetTemplateResponseBodyTemplate {
	s.CreateSource = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreationTime(v string) *GetTemplateResponseBodyTemplate {
	s.CreationTime = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetModifiedSource(v string) *GetTemplateResponseBodyTemplate {
	s.ModifiedSource = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetModifiedTime(v string) *GetTemplateResponseBodyTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetName(v string) *GetTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetPreviewMedia(v string) *GetTemplateResponseBodyTemplate {
	s.PreviewMedia = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetPreviewMediaStatus(v string) *GetTemplateResponseBodyTemplate {
	s.PreviewMediaStatus = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetRelatedMediaids(v string) *GetTemplateResponseBodyTemplate {
	s.RelatedMediaids = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetStatus(v string) *GetTemplateResponseBodyTemplate {
	s.Status = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateId(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetType(v string) *GetTemplateResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
