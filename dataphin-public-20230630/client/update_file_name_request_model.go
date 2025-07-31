// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *UpdateFileNameRequest
	GetFileId() *int64
	SetName(v string) *UpdateFileNameRequest
	GetName() *string
	SetOpTenantId(v int64) *UpdateFileNameRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *UpdateFileNameRequest
	GetProjectId() *int64
}

type UpdateFileNameRequest struct {
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
	// xxNew
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s UpdateFileNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileNameRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateFileNameRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFileNameRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateFileNameRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFileNameRequest) SetFileId(v int64) *UpdateFileNameRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileNameRequest) SetName(v string) *UpdateFileNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateFileNameRequest) SetOpTenantId(v int64) *UpdateFileNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateFileNameRequest) SetProjectId(v int64) *UpdateFileNameRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFileNameRequest) Validate() error {
	return dara.Validate(s)
}
