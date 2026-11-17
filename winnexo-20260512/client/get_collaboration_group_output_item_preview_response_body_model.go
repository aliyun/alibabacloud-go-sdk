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
	// The common business status code.
	//
	// example:
	//
	// ERR.Robject.Auth.TenantAccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The text content or preview prompt. This field can be empty.
	//
	// example:
	//
	// # Technical solution
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Indicates whether the preview is downgraded to a download preview.
	//
	// example:
	//
	// false
	Degraded *bool `json:"degraded,omitempty" xml:"degraded,omitempty"`
	// The short-term signed download URL. This field can be empty. If the URL expires, call this operation again.
	//
	// example:
	//
	// https://example.com/download/document.md
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The file name extension. This field can be empty.
	//
	// example:
	//
	// md
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The file name. This field can be empty.
	//
	// example:
	//
	// codex-r2-a-baseline.md
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The output detail ID.
	//
	// example:
	//
	// 51469466
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The common error description.
	//
	// example:
	//
	// Access denied to tenant 2147483647 because the user does not belong to the tenant
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The preview type.
	//
	// example:
	//
	// CONTENT
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// The short-term preview URL. This field can be empty.
	//
	// example:
	//
	// https://example.com/preview/document
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// The request tracking ID.
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
