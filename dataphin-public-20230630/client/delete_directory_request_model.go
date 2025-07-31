// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DeleteDirectoryRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *DeleteDirectoryRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *DeleteDirectoryRequest
	GetProjectId() *int64
}

type DeleteDirectoryRequest struct {
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

func (s DeleteDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeleteDirectoryRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDirectoryRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDirectoryRequest) SetFileId(v int64) *DeleteDirectoryRequest {
	s.FileId = &v
	return s
}

func (s *DeleteDirectoryRequest) SetOpTenantId(v int64) *DeleteDirectoryRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDirectoryRequest) SetProjectId(v int64) *DeleteDirectoryRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
