// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetBatchTaskInfoRequest
	GetEnv() *string
	SetFileId(v int64) *GetBatchTaskInfoRequest
	GetFileId() *int64
	SetIncludeAllUpStreams(v bool) *GetBatchTaskInfoRequest
	GetIncludeAllUpStreams() *bool
	SetOpTenantId(v int64) *GetBatchTaskInfoRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetBatchTaskInfoRequest
	GetProjectId() *int64
}

type GetBatchTaskInfoRequest struct {
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12113111
	FileId              *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IncludeAllUpStreams *bool  `json:"IncludeAllUpStreams,omitempty" xml:"IncludeAllUpStreams,omitempty"`
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

func (s GetBatchTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoRequest) GetEnv() *string {
	return s.Env
}

func (s *GetBatchTaskInfoRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *GetBatchTaskInfoRequest) GetIncludeAllUpStreams() *bool {
	return s.IncludeAllUpStreams
}

func (s *GetBatchTaskInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBatchTaskInfoRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTaskInfoRequest) SetEnv(v string) *GetBatchTaskInfoRequest {
	s.Env = &v
	return s
}

func (s *GetBatchTaskInfoRequest) SetFileId(v int64) *GetBatchTaskInfoRequest {
	s.FileId = &v
	return s
}

func (s *GetBatchTaskInfoRequest) SetIncludeAllUpStreams(v bool) *GetBatchTaskInfoRequest {
	s.IncludeAllUpStreams = &v
	return s
}

func (s *GetBatchTaskInfoRequest) SetOpTenantId(v int64) *GetBatchTaskInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBatchTaskInfoRequest) SetProjectId(v int64) *GetBatchTaskInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
