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
	// The new file name. This parameter is optional. If this parameter is not provided or set to an empty string, the original file name is retained.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The OSS persistent storage address of the replacement file.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The public access URL of the new file. The URL may contain a temporary signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID of the replacement file.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// Specifies whether to synchronously wait for re-parsing to complete. Default value: false, which indicates asynchronous queuing.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The ID of the FILE data source in the enterprise knowledge base to be replaced.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using --tenant-id.
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
