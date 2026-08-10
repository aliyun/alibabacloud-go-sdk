// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *SubmitPipelineByIdRequestContext) *SubmitPipelineByIdRequest
	GetContext() *SubmitPipelineByIdRequestContext
	SetOpTenantId(v int64) *SubmitPipelineByIdRequest
	GetOpTenantId() *int64
	SetQueryId(v *SubmitPipelineByIdRequestQueryId) *SubmitPipelineByIdRequest
	GetQueryId() *SubmitPipelineByIdRequestQueryId
}

type SubmitPipelineByIdRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *SubmitPipelineByIdRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID used to query the pipeline task.
	//
	// This parameter is required.
	QueryId *SubmitPipelineByIdRequestQueryId `json:"QueryId,omitempty" xml:"QueryId,omitempty" type:"Struct"`
}

func (s SubmitPipelineByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineByIdRequest) GoString() string {
	return s.String()
}

func (s *SubmitPipelineByIdRequest) GetContext() *SubmitPipelineByIdRequestContext {
	return s.Context
}

func (s *SubmitPipelineByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitPipelineByIdRequest) GetQueryId() *SubmitPipelineByIdRequestQueryId {
	return s.QueryId
}

func (s *SubmitPipelineByIdRequest) SetContext(v *SubmitPipelineByIdRequestContext) *SubmitPipelineByIdRequest {
	s.Context = v
	return s
}

func (s *SubmitPipelineByIdRequest) SetOpTenantId(v int64) *SubmitPipelineByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitPipelineByIdRequest) SetQueryId(v *SubmitPipelineByIdRequestQueryId) *SubmitPipelineByIdRequest {
	s.QueryId = v
	return s
}

func (s *SubmitPipelineByIdRequest) Validate() error {
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

type SubmitPipelineByIdRequestContext struct {
	// The current operating environment. Valid values:
	//
	// - DEV: the development environment.
	//
	// - PROD: the production environment.
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

func (s SubmitPipelineByIdRequestContext) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineByIdRequestContext) GoString() string {
	return s.String()
}

func (s *SubmitPipelineByIdRequestContext) GetEnv() *string {
	return s.Env
}

func (s *SubmitPipelineByIdRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitPipelineByIdRequestContext) SetEnv(v string) *SubmitPipelineByIdRequestContext {
	s.Env = &v
	return s
}

func (s *SubmitPipelineByIdRequestContext) SetProjectId(v int64) *SubmitPipelineByIdRequestContext {
	s.ProjectId = &v
	return s
}

func (s *SubmitPipelineByIdRequestContext) Validate() error {
	return dara.Validate(s)
}

type SubmitPipelineByIdRequestQueryId struct {
	// The file ID of the integration task. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The node ID of the integration task scheduling node. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The primary key ID of the integration pipeline. You can specify any one of pipelineId, fileId, or nodeId.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s SubmitPipelineByIdRequestQueryId) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineByIdRequestQueryId) GoString() string {
	return s.String()
}

func (s *SubmitPipelineByIdRequestQueryId) GetFileId() *int64 {
	return s.FileId
}

func (s *SubmitPipelineByIdRequestQueryId) GetNodeId() *string {
	return s.NodeId
}

func (s *SubmitPipelineByIdRequestQueryId) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *SubmitPipelineByIdRequestQueryId) SetFileId(v int64) *SubmitPipelineByIdRequestQueryId {
	s.FileId = &v
	return s
}

func (s *SubmitPipelineByIdRequestQueryId) SetNodeId(v string) *SubmitPipelineByIdRequestQueryId {
	s.NodeId = &v
	return s
}

func (s *SubmitPipelineByIdRequestQueryId) SetPipelineId(v int64) *SubmitPipelineByIdRequestQueryId {
	s.PipelineId = &v
	return s
}

func (s *SubmitPipelineByIdRequestQueryId) Validate() error {
	return dara.Validate(s)
}
