// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetFileName() *string
	SetFilePath(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetFileRecordId() *string
	SetForceSync(v bool) *ReplaceKnowledgeBaseSourceFileRequest
	GetForceSync() *bool
	SetSourceId(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceKnowledgeBaseSourceFileRequest
	GetTenantId() *string
}

type ReplaceKnowledgeBaseSourceFileRequest struct {
	// 新文件名（可选；不传或空字符串时保持原文件名）
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 新文件的 OSS 持久化地址（由上传签名接口返回）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// 新文件的公开访问 URL（可能携带临时签名）
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 新文件的文件记录 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// 是否同步等待重新解析完成；默认 false，异步入队
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// 待替换的企业知识库 FILE 数据源 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceFileRequest) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetFileName(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.FileName = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetFilePath(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.FilePath = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetFilePublicUrl(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetFileRecordId(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetForceSync(v bool) *ReplaceKnowledgeBaseSourceFileRequest {
	s.ForceSync = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetSourceId(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) SetTenantId(v string) *ReplaceKnowledgeBaseSourceFileRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileRequest) Validate() error {
	return dara.Validate(s)
}
