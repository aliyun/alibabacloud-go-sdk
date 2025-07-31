// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskUdfLineagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetBatchTaskUdfLineagesRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *GetBatchTaskUdfLineagesRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetBatchTaskUdfLineagesRequest
	GetProjectId() *int64
}

type GetBatchTaskUdfLineagesRequest struct {
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

func (s GetBatchTaskUdfLineagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskUdfLineagesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBatchTaskUdfLineagesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskUdfLineagesRequest) SetFileId(v int64) *GetBatchTaskUdfLineagesRequest {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesRequest) SetOpTenantId(v int64) *GetBatchTaskUdfLineagesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesRequest) SetProjectId(v int64) *GetBatchTaskUdfLineagesRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesRequest) Validate() error {
	return dara.Validate(s)
}
