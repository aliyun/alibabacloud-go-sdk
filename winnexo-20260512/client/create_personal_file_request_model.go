// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalFileRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFileRequest
	GetDirectoryId() *string
	SetFileExt(v string) *CreatePersonalFileRequest
	GetFileExt() *string
	SetFileName(v string) *CreatePersonalFileRequest
	GetFileName() *string
	SetFilePath(v string) *CreatePersonalFileRequest
	GetFilePath() *string
	SetFilePublicUrl(v string) *CreatePersonalFileRequest
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *CreatePersonalFileRequest
	GetFileRecordId() *string
	SetName(v string) *CreatePersonalFileRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalFileRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalFileRequest
	GetTenantId() *string
}

type CreatePersonalFileRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// created by eventbridge
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The file extension (optional, such as pdf or docx).
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The file name.
	//
	// example:
	//
	// 0250705120003-2026-04-28-19-22-20.wav
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// bi/batch-query-service.app.yaml
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The publicly accessible URL of the Alibaba DingTalk online document.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID (optional, corresponding to settings.file_record_id).
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The pipeline name.
	//
	// This parameter is required.
	//
	// example:
	//
	// sys_first_new_v3_b
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// 1798284341201499
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFileRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFileRequest) GetFileExt() *string {
	return s.FileExt
}

func (s *CreatePersonalFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreatePersonalFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreatePersonalFileRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreatePersonalFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *CreatePersonalFileRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFileRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFileRequest) SetDescription(v string) *CreatePersonalFileRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFileRequest) SetDirectoryId(v string) *CreatePersonalFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFileRequest) SetFileExt(v string) *CreatePersonalFileRequest {
	s.FileExt = &v
	return s
}

func (s *CreatePersonalFileRequest) SetFileName(v string) *CreatePersonalFileRequest {
	s.FileName = &v
	return s
}

func (s *CreatePersonalFileRequest) SetFilePath(v string) *CreatePersonalFileRequest {
	s.FilePath = &v
	return s
}

func (s *CreatePersonalFileRequest) SetFilePublicUrl(v string) *CreatePersonalFileRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreatePersonalFileRequest) SetFileRecordId(v string) *CreatePersonalFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *CreatePersonalFileRequest) SetName(v string) *CreatePersonalFileRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalFileRequest) SetOperatingObjectName(v string) *CreatePersonalFileRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFileRequest) SetTenantId(v string) *CreatePersonalFileRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFileRequest) Validate() error {
	return dara.Validate(s)
}
