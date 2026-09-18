// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGroupSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PreviewGroupSourceResponseBody
	GetCode() *string
	SetContent(v string) *PreviewGroupSourceResponseBody
	GetContent() *string
	SetDegraded(v bool) *PreviewGroupSourceResponseBody
	GetDegraded() *bool
	SetFileExt(v string) *PreviewGroupSourceResponseBody
	GetFileExt() *string
	SetFileName(v string) *PreviewGroupSourceResponseBody
	GetFileName() *string
	SetMessage(v string) *PreviewGroupSourceResponseBody
	GetMessage() *string
	SetPreviewType(v string) *PreviewGroupSourceResponseBody
	GetPreviewType() *string
	SetPreviewUrl(v string) *PreviewGroupSourceResponseBody
	GetPreviewUrl() *string
	SetPublicUrl(v string) *PreviewGroupSourceResponseBody
	GetPublicUrl() *string
	SetRequestId(v string) *PreviewGroupSourceResponseBody
	GetRequestId() *string
}

type PreviewGroupSourceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The text content. This field is used for the CONTENT preview type.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Indicates whether the preview is degraded to a download, meaning the resource cannot be opened in the online previewer. Valid values:
	//
	// - true: The preview is degraded to a download.
	//
	// - false: The resource can be previewed online.
	Degraded *bool `json:"degraded,omitempty" xml:"degraded,omitempty"`
	// The file name extension. This field is used for the OSS_IMM preview type.
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The file name. This field is used for the OSS_IMM preview type.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The error details.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The preview type. Valid values: OSS_IMM, IMAGE, AUDIO, VIDEO, HTML, DING_TALK, VOICE_MEETING, and CONTENT.
	//
	// example:
	//
	// OSS_IMM
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// The preview URL. This field is used for the OSS_IMM, DING_TALK, and VOICE_MEETING preview types.
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
	// The request ID.
	//
	// example:
	//
	// E68654BD-F7BA-5837-8686-5645D739A47C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PreviewGroupSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewGroupSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewGroupSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PreviewGroupSourceResponseBody) GetContent() *string {
	return s.Content
}

func (s *PreviewGroupSourceResponseBody) GetDegraded() *bool {
	return s.Degraded
}

func (s *PreviewGroupSourceResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *PreviewGroupSourceResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *PreviewGroupSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewGroupSourceResponseBody) GetPreviewType() *string {
	return s.PreviewType
}

func (s *PreviewGroupSourceResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *PreviewGroupSourceResponseBody) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *PreviewGroupSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewGroupSourceResponseBody) SetCode(v string) *PreviewGroupSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetContent(v string) *PreviewGroupSourceResponseBody {
	s.Content = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetDegraded(v bool) *PreviewGroupSourceResponseBody {
	s.Degraded = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetFileExt(v string) *PreviewGroupSourceResponseBody {
	s.FileExt = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetFileName(v string) *PreviewGroupSourceResponseBody {
	s.FileName = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetMessage(v string) *PreviewGroupSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetPreviewType(v string) *PreviewGroupSourceResponseBody {
	s.PreviewType = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetPreviewUrl(v string) *PreviewGroupSourceResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetPublicUrl(v string) *PreviewGroupSourceResponseBody {
	s.PublicUrl = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) SetRequestId(v string) *PreviewGroupSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewGroupSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
