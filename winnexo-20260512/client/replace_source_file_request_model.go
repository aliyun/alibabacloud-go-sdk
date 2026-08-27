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
