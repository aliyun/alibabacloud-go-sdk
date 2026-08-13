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
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标个人目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时必须是当前用户在当前数字员工下的已有个人目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 文件后缀名（可选，如 pdf、docx）
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// 原始文件名（可选，含后缀）
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件 OSS 持久化地址（必填，对应 settings.file_path）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 文件公开访问 URL（可选，带签名，对应 settings.file_public_url）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 文件记录 ID（可选，对应 settings.file_record_id）
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// 资源显示名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工名称（已废弃：不再作为个人资源隔离条件，仅保留用于来源追溯）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
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
