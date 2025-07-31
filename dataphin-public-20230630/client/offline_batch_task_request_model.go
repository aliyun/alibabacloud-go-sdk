// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *OfflineBatchTaskRequest
	GetComment() *string
	SetFileId(v int64) *OfflineBatchTaskRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *OfflineBatchTaskRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *OfflineBatchTaskRequest
	GetProjectId() *int64
}

type OfflineBatchTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test xx
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
}

func (s OfflineBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *OfflineBatchTaskRequest) GetComment() *string {
	return s.Comment
}

func (s *OfflineBatchTaskRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *OfflineBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflineBatchTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *OfflineBatchTaskRequest) SetComment(v string) *OfflineBatchTaskRequest {
	s.Comment = &v
	return s
}

func (s *OfflineBatchTaskRequest) SetFileId(v int64) *OfflineBatchTaskRequest {
	s.FileId = &v
	return s
}

func (s *OfflineBatchTaskRequest) SetOpTenantId(v int64) *OfflineBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflineBatchTaskRequest) SetProjectId(v int64) *OfflineBatchTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *OfflineBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}
