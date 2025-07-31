// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetAdHocFileRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *GetAdHocFileRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetAdHocFileRequest
	GetProjectId() *int64
}

type GetAdHocFileRequest struct {
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

func (s GetAdHocFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *GetAdHocFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetAdHocFileRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAdHocFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAdHocFileRequest) SetFileId(v int64) *GetAdHocFileRequest {
	s.FileId = &v
	return s
}

func (s *GetAdHocFileRequest) SetOpTenantId(v int64) *GetAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAdHocFileRequest) SetProjectId(v int64) *GetAdHocFileRequest {
	s.ProjectId = &v
	return s
}

func (s *GetAdHocFileRequest) Validate() error {
	return dara.Validate(s)
}
