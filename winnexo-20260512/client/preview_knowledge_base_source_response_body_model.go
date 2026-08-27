// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewKnowledgeBaseSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetCode() *string
	SetContent(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetContent() *string
	SetFileExt(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetFileExt() *string
	SetFileName(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetFileName() *string
	SetMessage(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetMessage() *string
	SetPreviewType(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetPreviewType() *string
	SetPreviewUrl(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetPreviewUrl() *string
	SetPublicUrl(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetPublicUrl() *string
	SetRequestId(v string) *PreviewKnowledgeBaseSourceResponseBody
	GetRequestId() *string
}

type PreviewKnowledgeBaseSourceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The text content. This is used for the CONTENT type.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The file name extension. This is used for the OSS_IMM type.
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The file name. This is used for the OSS_IMM type.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The preview type. Valid values: OSS_IMM, IMAGE, AUDIO, VIDEO, HTML, DING_TALK, VOICE_MEETING, CONTENT.
	//
	// example:
	//
	// OSS_IMM
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// The preview URL. This is used for the OSS_IMM, DING_TALK, and VOICE_MEETING types.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// The public download URL of the file.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	PublicUrl *string `json:"publicUrl,omitempty" xml:"publicUrl,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PreviewKnowledgeBaseSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewKnowledgeBaseSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetContent() *string {
	return s.Content
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetPreviewType() *string {
	return s.PreviewType
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *PreviewKnowledgeBaseSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetCode(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetContent(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.Content = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetFileExt(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.FileExt = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetFileName(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.FileName = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetMessage(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetPreviewType(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.PreviewType = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetPreviewUrl(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetPublicUrl(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.PublicUrl = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) SetRequestId(v string) *PreviewKnowledgeBaseSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
