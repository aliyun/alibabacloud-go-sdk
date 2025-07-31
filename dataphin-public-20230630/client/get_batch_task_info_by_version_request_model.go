// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoByVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetBatchTaskInfoByVersionRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *GetBatchTaskInfoByVersionRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetBatchTaskInfoByVersionRequest
	GetProjectId() *int64
	SetVersionId(v int64) *GetBatchTaskInfoByVersionRequest
	GetVersionId() *int64
}

type GetBatchTaskInfoByVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12113111
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
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetBatchTaskInfoByVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskInfoByVersionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBatchTaskInfoByVersionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskInfoByVersionRequest) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetBatchTaskInfoByVersionRequest) SetFileId(v int64) *GetBatchTaskInfoByVersionRequest {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionRequest) SetOpTenantId(v int64) *GetBatchTaskInfoByVersionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionRequest) SetProjectId(v int64) *GetBatchTaskInfoByVersionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionRequest) SetVersionId(v int64) *GetBatchTaskInfoByVersionRequest {
	s.VersionId = &v
	return s
}

func (s *GetBatchTaskInfoByVersionRequest) Validate() error {
	return dara.Validate(s)
}
