// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *GetBatchTaskVersionsRequest
	GetFileId() *int64
	SetOpTenantId(v int64) *GetBatchTaskVersionsRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetBatchTaskVersionsRequest
	GetProjectId() *int64
}

type GetBatchTaskVersionsRequest struct {
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

func (s GetBatchTaskVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskVersionsRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskVersionsRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskVersionsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBatchTaskVersionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskVersionsRequest) SetFileId(v int64) *GetBatchTaskVersionsRequest {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskVersionsRequest) SetOpTenantId(v int64) *GetBatchTaskVersionsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBatchTaskVersionsRequest) SetProjectId(v int64) *GetBatchTaskVersionsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskVersionsRequest) Validate() error {
	return dara.Validate(s)
}
