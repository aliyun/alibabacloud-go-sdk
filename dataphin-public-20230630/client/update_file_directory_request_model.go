// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectory(v string) *UpdateFileDirectoryRequest
	GetDirectory() *string
	SetFileId(v int64) *UpdateFileDirectoryRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *UpdateFileDirectoryRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *UpdateFileDirectoryRequest
	GetProjectId() *int64
}

type UpdateFileDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /xx测试/目录new
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateFileDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryRequest) GetDirectory() *string {
	return s.Directory
}

func (s *UpdateFileDirectoryRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateFileDirectoryRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateFileDirectoryRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFileDirectoryRequest) SetDirectory(v string) *UpdateFileDirectoryRequest {
	s.Directory = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetFileId(v int64) *UpdateFileDirectoryRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetOpTenantId(v int64) *UpdateFileDirectoryRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetProjectId(v int64) *UpdateFileDirectoryRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFileDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
