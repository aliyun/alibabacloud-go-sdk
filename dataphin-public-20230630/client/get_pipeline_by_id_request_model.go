// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *GetPipelineByIdRequestContext) *GetPipelineByIdRequest
	GetContext() *GetPipelineByIdRequestContext
	SetOpTenantId(v int64) *GetPipelineByIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetPipelineByIdRequest
	GetOpUserId() *string
	SetQueryId(v *GetPipelineByIdRequestQueryId) *GetPipelineByIdRequest
	GetQueryId() *GetPipelineByIdRequestQueryId
}

type GetPipelineByIdRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *GetPipelineByIdRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID used to query the pipeline node.
	//
	// This parameter is required.
	QueryId *GetPipelineByIdRequestQueryId `json:"QueryId,omitempty" xml:"QueryId,omitempty" type:"Struct"`
}

func (s GetPipelineByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdRequest) GetContext() *GetPipelineByIdRequestContext {
	return s.Context
}

func (s *GetPipelineByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPipelineByIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetPipelineByIdRequest) GetQueryId() *GetPipelineByIdRequestQueryId {
	return s.QueryId
}

func (s *GetPipelineByIdRequest) SetContext(v *GetPipelineByIdRequestContext) *GetPipelineByIdRequest {
	s.Context = v
	return s
}

func (s *GetPipelineByIdRequest) SetOpTenantId(v int64) *GetPipelineByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPipelineByIdRequest) SetOpUserId(v string) *GetPipelineByIdRequest {
	s.OpUserId = &v
	return s
}

func (s *GetPipelineByIdRequest) SetQueryId(v *GetPipelineByIdRequestQueryId) *GetPipelineByIdRequest {
	s.QueryId = v
	return s
}

func (s *GetPipelineByIdRequest) Validate() error {
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

type GetPipelineByIdRequestContext struct {
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
	// The ID of the project to which the integration pipeline node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPipelineByIdRequestContext) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdRequestContext) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdRequestContext) GetEnv() *string {
	return s.Env
}

func (s *GetPipelineByIdRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPipelineByIdRequestContext) SetEnv(v string) *GetPipelineByIdRequestContext {
	s.Env = &v
	return s
}

func (s *GetPipelineByIdRequestContext) SetProjectId(v int64) *GetPipelineByIdRequestContext {
	s.ProjectId = &v
	return s
}

func (s *GetPipelineByIdRequestContext) Validate() error {
	return dara.Validate(s)
}

type GetPipelineByIdRequestQueryId struct {
	// The file ID of the integration node. You can use this parameter to query the pipeline node.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The scheduling node ID of the integration node. You can use this parameter to query the pipeline node.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The primary key ID of the integration pipeline. You can use this parameter to query the pipeline node.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s GetPipelineByIdRequestQueryId) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdRequestQueryId) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdRequestQueryId) GetFileId() *int64 {
	return s.FileId
}

func (s *GetPipelineByIdRequestQueryId) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPipelineByIdRequestQueryId) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetPipelineByIdRequestQueryId) SetFileId(v int64) *GetPipelineByIdRequestQueryId {
	s.FileId = &v
	return s
}

func (s *GetPipelineByIdRequestQueryId) SetNodeId(v string) *GetPipelineByIdRequestQueryId {
	s.NodeId = &v
	return s
}

func (s *GetPipelineByIdRequestQueryId) SetPipelineId(v int64) *GetPipelineByIdRequestQueryId {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineByIdRequestQueryId) Validate() error {
	return dara.Validate(s)
}
