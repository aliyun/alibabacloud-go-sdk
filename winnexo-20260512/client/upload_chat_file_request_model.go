// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UploadChatFileRequest
	GetContentType() *string
	SetFileName(v string) *UploadChatFileRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadChatFileRequest
	GetFileUrl() *string
	SetOperatingObjectName(v string) *UploadChatFileRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *UploadChatFileRequest
	GetTenantId() *string
}

type UploadChatFileRequest struct {
	// The content type of the file. Valid values:
	//
	// - **image**: image
	//
	// - **document**: general document
	//
	// - **alidoc**: Alibaba document
	//
	// - **text**: text
	//
	// - **video**: video
	//
	// - **audio**: audio
	//
	// - **archive**: archive
	//
	// - **app**: application
	//
	// - **link**: shortcut
	//
	// - **other**: other
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The full path name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// report.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The attachment address.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://winnexo-file-transfer.oss-cn-hangzhou.aliyuncs.com/openapi/2026-08-06/9f8c2a1b
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UploadChatFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileRequest) GoString() string {
	return s.String()
}

func (s *UploadChatFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadChatFileRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UploadChatFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileRequest) SetContentType(v string) *UploadChatFileRequest {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileRequest) SetFileName(v string) *UploadChatFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadChatFileRequest) SetFileUrl(v string) *UploadChatFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadChatFileRequest) SetOperatingObjectName(v string) *UploadChatFileRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UploadChatFileRequest) SetTenantId(v string) *UploadChatFileRequest {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileRequest) Validate() error {
	return dara.Validate(s)
}
