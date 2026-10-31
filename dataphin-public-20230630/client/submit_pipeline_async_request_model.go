// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *SubmitPipelineAsyncRequestContext) *SubmitPipelineAsyncRequest
	GetContext() *SubmitPipelineAsyncRequestContext
	SetOpTenantId(v int64) *SubmitPipelineAsyncRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitPipelineAsyncRequest
	GetOpUserId() *string
	SetQueryId(v *SubmitPipelineAsyncRequestQueryId) *SubmitPipelineAsyncRequest
	GetQueryId() *SubmitPipelineAsyncRequestQueryId
}

type SubmitPipelineAsyncRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *SubmitPipelineAsyncRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID used to query the pipeline task.
	//
	// This parameter is required.
	QueryId *SubmitPipelineAsyncRequestQueryId `json:"QueryId,omitempty" xml:"QueryId,omitempty" type:"Struct"`
}

func (s SubmitPipelineAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncRequest) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncRequest) GetContext() *SubmitPipelineAsyncRequestContext {
	return s.Context
}

func (s *SubmitPipelineAsyncRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitPipelineAsyncRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitPipelineAsyncRequest) GetQueryId() *SubmitPipelineAsyncRequestQueryId {
	return s.QueryId
}

func (s *SubmitPipelineAsyncRequest) SetContext(v *SubmitPipelineAsyncRequestContext) *SubmitPipelineAsyncRequest {
	s.Context = v
	return s
}

func (s *SubmitPipelineAsyncRequest) SetOpTenantId(v int64) *SubmitPipelineAsyncRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitPipelineAsyncRequest) SetOpUserId(v string) *SubmitPipelineAsyncRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitPipelineAsyncRequest) SetQueryId(v *SubmitPipelineAsyncRequestQueryId) *SubmitPipelineAsyncRequest {
	s.QueryId = v
	return s
}

func (s *SubmitPipelineAsyncRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.QueryId != nil {
		if err := s.QueryId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitPipelineAsyncRequestContext struct {
	// The current operation environment. Valid values: DEV: development environment. PROD: production environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The ID of the project to which the integration pipeline task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SubmitPipelineAsyncRequestContext) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncRequestContext) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncRequestContext) GetEnv() *string {
	return s.Env
}

func (s *SubmitPipelineAsyncRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitPipelineAsyncRequestContext) SetEnv(v string) *SubmitPipelineAsyncRequestContext {
	s.Env = &v
	return s
}

func (s *SubmitPipelineAsyncRequestContext) SetProjectId(v int64) *SubmitPipelineAsyncRequestContext {
	s.ProjectId = &v
	return s
}

func (s *SubmitPipelineAsyncRequestContext) Validate() error {
	return dara.Validate(s)
}

type SubmitPipelineAsyncRequestQueryId struct {
	// The ID of the integration task file. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the scheduling node for the integration task. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The primary key of the integration pipeline. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s SubmitPipelineAsyncRequestQueryId) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncRequestQueryId) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncRequestQueryId) GetFileId() *int64 {
	return s.FileId
}

func (s *SubmitPipelineAsyncRequestQueryId) GetNodeId() *string {
	return s.NodeId
}

func (s *SubmitPipelineAsyncRequestQueryId) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *SubmitPipelineAsyncRequestQueryId) SetFileId(v int64) *SubmitPipelineAsyncRequestQueryId {
	s.FileId = &v
	return s
}

func (s *SubmitPipelineAsyncRequestQueryId) SetNodeId(v string) *SubmitPipelineAsyncRequestQueryId {
	s.NodeId = &v
	return s
}

func (s *SubmitPipelineAsyncRequestQueryId) SetPipelineId(v int64) *SubmitPipelineAsyncRequestQueryId {
	s.PipelineId = &v
	return s
}

func (s *SubmitPipelineAsyncRequestQueryId) Validate() error {
	return dara.Validate(s)
}
