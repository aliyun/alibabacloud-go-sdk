// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateKnowledgeBaseFileRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *CreateKnowledgeBaseFileRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateKnowledgeBaseFileRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateKnowledgeBaseFileRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseFileRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseFileRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the alias.
	//
	// example:
	//
	// hangzhou-release-version-3-eventbridge-numeric-queue-fix-20260529
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The file name extension. This parameter is optional. Examples: pdf and docx.
	//
	// example:
	//
	// string_value
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The file name.
	//
	// example:
	//
	// BasicSimilarityScorer.cava
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// bi/batch-query-service.app.yaml
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The publicly accessible URL of the DingTalk online document.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID. This parameter is optional and corresponds to settings.file_record_id.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// Not supported. Ignore this parameter.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// The name of the AI assistant.
	//
	// This parameter is required.
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name). This parameter is optional.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The resource labels. This parameter is optional. Specify a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFileRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFileRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreateKnowledgeBaseFileRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateKnowledgeBaseFileRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateKnowledgeBaseFileRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseFileRequest) SetAppToken(v string) *CreateKnowledgeBaseFileRequest {
	s.AppToken = &v
	return s
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

func (s *CreateKnowledgeBaseFileRequest) SetOpsUserCode(v string) *CreateKnowledgeBaseFileRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateKnowledgeBaseFileRequest) SetOpsUserId(v string) *CreateKnowledgeBaseFileRequest {
	s.OpsUserId = &v
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
