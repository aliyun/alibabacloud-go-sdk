// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceSourceFileRequest
	GetAppToken() *string
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
	SetOpsUserCode(v string) *ReplaceSourceFileRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceSourceFileRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReplaceSourceFileRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceSourceFileRequest
	GetTenantId() *string
}

type ReplaceSourceFileRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new file name. This parameter is optional. If you do not specify this parameter or set it to an empty string, the original file name is retained.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The OSS persistent address of the new file, returned by the upload signing operation.
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
	// The file record ID of the new file.
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
	// The ID of the personal FILE data source to be replaced. This ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
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

func (s *ReplaceSourceFileRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *ReplaceSourceFileRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceSourceFileRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceSourceFileRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceSourceFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceSourceFileRequest) SetAppToken(v string) *ReplaceSourceFileRequest {
	s.AppToken = &v
	return s
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

func (s *ReplaceSourceFileRequest) SetOpsUserCode(v string) *ReplaceSourceFileRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceSourceFileRequest) SetOpsUserId(v string) *ReplaceSourceFileRequest {
	s.OpsUserId = &v
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
