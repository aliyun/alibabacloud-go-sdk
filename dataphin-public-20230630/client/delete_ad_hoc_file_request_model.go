// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdHocFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DeleteAdHocFileRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *DeleteAdHocFileRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *DeleteAdHocFileRequest
	GetProjectId() *int64
}

type DeleteAdHocFileRequest struct {
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

func (s DeleteAdHocFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeleteAdHocFileRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteAdHocFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteAdHocFileRequest) SetFileId(v int64) *DeleteAdHocFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteAdHocFileRequest) SetOpTenantId(v int64) *DeleteAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteAdHocFileRequest) SetProjectId(v int64) *DeleteAdHocFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteAdHocFileRequest) Validate() error {
	return dara.Validate(s)
}
