// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseFileRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFileRequest
	GetDirectoryId() *string
	SetFileExt(v string) *CreateKnowledgeBaseFileRequest
	GetFileExt() *string
	SetFileName(v string) *CreateKnowledgeBaseFileRequest
	GetFileName() *string
	SetFilePath(v string) *CreateKnowledgeBaseFileRequest
	GetFilePath() *string
	SetFilePublicUrl(v string) *CreateKnowledgeBaseFileRequest
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *CreateKnowledgeBaseFileRequest
	GetFileRecordId() *string
	SetKnowledgeId(v string) *CreateKnowledgeBaseFileRequest
	GetKnowledgeId() *string
	SetName(v string) *CreateKnowledgeBaseFileRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBaseFileRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateKnowledgeBaseFileRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseFileRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseFileRequest struct {
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标企业知识库目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时必须是当前租户下已有的企业知识库目录
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
	// 知识库 ID（可选，透传给 document_agent）
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// 资源显示名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工名称（运营对象 name，可选）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 资源标签（可选，JSON 字符串列表，如 ["tagA","tagB"]）
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFileRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFileRequest) GetFileExt() *string {
	return s.FileExt
}

func (s *CreateKnowledgeBaseFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateKnowledgeBaseFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateKnowledgeBaseFileRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateKnowledgeBaseFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *CreateKnowledgeBaseFileRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *CreateKnowledgeBaseFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFileRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseFileRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseFileRequest) SetDescription(v string) *CreateKnowledgeBaseFileRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetDirectoryId(v string) *CreateKnowledgeBaseFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetFileExt(v string) *CreateKnowledgeBaseFileRequest {
	s.FileExt = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetFileName(v string) *CreateKnowledgeBaseFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetFilePath(v string) *CreateKnowledgeBaseFileRequest {
	s.FilePath = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetFilePublicUrl(v string) *CreateKnowledgeBaseFileRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetFileRecordId(v string) *CreateKnowledgeBaseFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetKnowledgeId(v string) *CreateKnowledgeBaseFileRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetName(v string) *CreateKnowledgeBaseFileRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseFileRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetSourceTags(v string) *CreateKnowledgeBaseFileRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetTenantId(v string) *CreateKnowledgeBaseFileRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) Validate() error {
	return dara.Validate(s)
}
