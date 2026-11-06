// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCollaborationGroupOutputItemPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetCode() *string
	SetContent(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetContent() *string
	SetDegraded(v bool) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetDegraded() *bool
	SetDownloadUrl(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetDownloadUrl() *string
	SetFileExt(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetFileExt() *string
	SetFileName(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetFileName() *string
	SetItemId(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetItemId() *string
	SetMessage(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetMessage() *string
	SetPreviewType(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetPreviewType() *string
	SetPreviewUrl(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetPreviewUrl() *string
	SetRequestId(v string) *GetCollaborationGroupOutputItemPreviewResponseBody
	GetRequestId() *string
}

type GetCollaborationGroupOutputItemPreviewResponseBody struct {
	// 公共业务状态码，由后端Code字段映射；错误时用于识别失败原因。
	//
	// example:
	//
	// ERR.Robject.Auth.TenantAccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 可直接展示的文本正文或预览提示；不提供正文的文件类型可为空。
	//
	// example:
	//
	// # 技术方案
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否降级为下载。true时应展示下载入口，不要按在线文档预览器处理。
	//
	// example:
	//
	// false
	Degraded *bool `json:"degraded,omitempty" xml:"degraded,omitempty"`
	// 短期签名下载地址；普通文件有效期3600秒，邮件正文可为空。过期后重新调用本接口。示例仅展示URL格式，不包含真实签名。
	//
	// example:
	//
	// https://example.com/download/document.md
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// 文件扩展名，不含点；无扩展名时可为空。
	//
	// example:
	//
	// md
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// 文件名，包含扩展名；无文件名时可为空。
	//
	// example:
	//
	// codex-r2-a-baseline.md
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 对话生成文件的产出明细ID。取GetNodeDetail返回的artifacts中ref.kind为USER_OUTPUT_ITEM的ref.id；不是artifactId。
	//
	// example:
	//
	// 51469466
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 公共错误说明，由后端Message字段映射；错误时描述失败原因。
	//
	// example:
	//
	// 无权访问租户 2147483647，用户不属于该租户
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 预览方式，例如CONTENT为文本正文，HTML为网页，IMAGE为图片，AUDIO为音频，VIDEO为视频，OSS_IMM为文档预览。
	//
	// example:
	//
	// CONTENT
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// 短期预览地址；正文可直接展示时可为空。地址过期后重新调用本接口。示例仅展示URL格式，不是可访问凭证。
	//
	// example:
	//
	// https://example.com/preview/document
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// 请求追踪ID，由后端RequestId字段映射，用于排查调用失败。
	//
	// example:
	//
	// c9cf6cb3-3f5c-4994-a8a1-c93379a6f5cf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCollaborationGroupOutputItemPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCollaborationGroupOutputItemPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetDegraded() *bool {
	return s.Degraded
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetItemId() *string {
	return s.ItemId
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetPreviewType() *string {
	return s.PreviewType
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetCode(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetContent(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.Content = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetDegraded(v bool) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.Degraded = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetDownloadUrl(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetFileExt(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.FileExt = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetFileName(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.FileName = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetItemId(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.ItemId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetMessage(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetPreviewType(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.PreviewType = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetPreviewUrl(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) SetRequestId(v string) *GetCollaborationGroupOutputItemPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCollaborationGroupOutputItemPreviewResponseBody) Validate() error {
	return dara.Validate(s)
}
