// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceGroupSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ReplaceGroupSourceFileRequest
	GetFileName() *string
	SetFilePath(v string) *ReplaceGroupSourceFileRequest
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceGroupSourceFileRequest
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceGroupSourceFileRequest
	GetFileRecordId() *string
	SetForceSync(v bool) *ReplaceGroupSourceFileRequest
	GetForceSync() *bool
	SetGroupId(v string) *ReplaceGroupSourceFileRequest
	GetGroupId() *string
	SetSourceId(v string) *ReplaceGroupSourceFileRequest
	GetSourceId() *string
	SetTenantId(v string) *ReplaceGroupSourceFileRequest
	GetTenantId() *string
}

type ReplaceGroupSourceFileRequest struct {
	// The new file name. This parameter is optional. If you do not specify this parameter or set it to an empty string, the original file name is retained.
	//
	// example:
	//
	// example
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The OSS persistent storage path of the replacement file.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://example/new.txt
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The OSS persistent storage path of the replacement file.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/new.txt
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID of the replacement file.
	//
	// This parameter is required.
	//
	// example:
	//
	// file_example
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// Specifies whether to synchronously wait for re-parsing to complete. Default value: false, which means the task is asynchronously enqueued.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReplaceGroupSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceGroupSourceFileRequest) GoString() string {
	return s.String()
}

func (s *ReplaceGroupSourceFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ReplaceGroupSourceFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceGroupSourceFileRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceGroupSourceFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceGroupSourceFileRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReplaceGroupSourceFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReplaceGroupSourceFileRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceGroupSourceFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceGroupSourceFileRequest) SetFileName(v string) *ReplaceGroupSourceFileRequest {
	s.FileName = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetFilePath(v string) *ReplaceGroupSourceFileRequest {
	s.FilePath = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetFilePublicUrl(v string) *ReplaceGroupSourceFileRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetFileRecordId(v string) *ReplaceGroupSourceFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetForceSync(v bool) *ReplaceGroupSourceFileRequest {
	s.ForceSync = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetGroupId(v string) *ReplaceGroupSourceFileRequest {
	s.GroupId = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetSourceId(v string) *ReplaceGroupSourceFileRequest {
	s.SourceId = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) SetTenantId(v string) *ReplaceGroupSourceFileRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceGroupSourceFileRequest) Validate() error {
	return dara.Validate(s)
}
