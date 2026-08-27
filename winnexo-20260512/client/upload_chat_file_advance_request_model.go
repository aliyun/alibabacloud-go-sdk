// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadChatFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UploadChatFileAdvanceRequest
	GetContentType() *string
	SetFileName(v string) *UploadChatFileAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *UploadChatFileAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOperatingObjectName(v string) *UploadChatFileAdvanceRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *UploadChatFileAdvanceRequest
	GetTenantId() *string
}

type UploadChatFileAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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

func (s UploadChatFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadChatFileAdvanceRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *UploadChatFileAdvanceRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UploadChatFileAdvanceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadChatFileAdvanceRequest) SetContentType(v string) *UploadChatFileAdvanceRequest {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetFileName(v string) *UploadChatFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadChatFileAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetOperatingObjectName(v string) *UploadChatFileAdvanceRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) SetTenantId(v string) *UploadChatFileAdvanceRequest {
	s.TenantId = &v
	return s
}

func (s *UploadChatFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
