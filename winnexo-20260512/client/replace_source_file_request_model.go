// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ReplaceSourceFileRequest
	GetFileName() *string
	SetFilePath(v string) *ReplaceSourceFileRequest
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceSourceFileRequest
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceSourceFileRequest
	GetFileRecordId() *string
	SetForceSync(v bool) *ReplaceSourceFileRequest
	GetForceSync() *bool
	SetSourceId(v string) *ReplaceSourceFileRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceSourceFileRequest
	GetTenantId() *string
}

type ReplaceSourceFileRequest struct {
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
	// 待替换的个人 FILE 数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSourceFileRequest) GoString() string {
	return s.String()
}

func (s *ReplaceSourceFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ReplaceSourceFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceSourceFileRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceSourceFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceSourceFileRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReplaceSourceFileRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceSourceFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceSourceFileRequest) SetFileName(v string) *ReplaceSourceFileRequest {
	s.FileName = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetFilePath(v string) *ReplaceSourceFileRequest {
	s.FilePath = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetFilePublicUrl(v string) *ReplaceSourceFileRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetFileRecordId(v string) *ReplaceSourceFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetForceSync(v bool) *ReplaceSourceFileRequest {
	s.ForceSync = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetSourceId(v string) *ReplaceSourceFileRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetTenantId(v string) *ReplaceSourceFileRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceSourceFileRequest) Validate() error {
	return dara.Validate(s)
}
