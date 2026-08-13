// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PreviewPersonalSourceResponseBody
	GetCode() *string
	SetContent(v string) *PreviewPersonalSourceResponseBody
	GetContent() *string
	SetFileExt(v string) *PreviewPersonalSourceResponseBody
	GetFileExt() *string
	SetFileName(v string) *PreviewPersonalSourceResponseBody
	GetFileName() *string
	SetMessage(v string) *PreviewPersonalSourceResponseBody
	GetMessage() *string
	SetPreviewType(v string) *PreviewPersonalSourceResponseBody
	GetPreviewType() *string
	SetPreviewUrl(v string) *PreviewPersonalSourceResponseBody
	GetPreviewUrl() *string
	SetPublicUrl(v string) *PreviewPersonalSourceResponseBody
	GetPublicUrl() *string
	SetRequestId(v string) *PreviewPersonalSourceResponseBody
	GetRequestId() *string
}

type PreviewPersonalSourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 文本内容（CONTENT类型使用）
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 文件扩展名（OSS_IMM类型使用）
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// 文件名（OSS_IMM类型使用）
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 预览类型：OSS_IMM / IMAGE / AUDIO / VIDEO / HTML / DING_TALK / VOICE_MEETING / CONTENT
	//
	// example:
	//
	// OSS_IMM
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// 预览URL（OSS_IMM、DING_TALK、VOICE_MEETING使用）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// 公开下载URL（可供下载的文件URL）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	PublicUrl *string `json:"publicUrl,omitempty" xml:"publicUrl,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PreviewPersonalSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PreviewPersonalSourceResponseBody) GetContent() *string {
	return s.Content
}

func (s *PreviewPersonalSourceResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *PreviewPersonalSourceResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *PreviewPersonalSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewPersonalSourceResponseBody) GetPreviewType() *string {
	return s.PreviewType
}

func (s *PreviewPersonalSourceResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *PreviewPersonalSourceResponseBody) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *PreviewPersonalSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewPersonalSourceResponseBody) SetCode(v string) *PreviewPersonalSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetContent(v string) *PreviewPersonalSourceResponseBody {
	s.Content = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetFileExt(v string) *PreviewPersonalSourceResponseBody {
	s.FileExt = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetFileName(v string) *PreviewPersonalSourceResponseBody {
	s.FileName = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetMessage(v string) *PreviewPersonalSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetPreviewType(v string) *PreviewPersonalSourceResponseBody {
	s.PreviewType = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetPreviewUrl(v string) *PreviewPersonalSourceResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetPublicUrl(v string) *PreviewPersonalSourceResponseBody {
	s.PublicUrl = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) SetRequestId(v string) *PreviewPersonalSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewPersonalSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
