// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineSubmitResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *QueryPipelineSubmitResultRequestContext) *QueryPipelineSubmitResultRequest
	GetContext() *QueryPipelineSubmitResultRequestContext
	SetOpTenantId(v int64) *QueryPipelineSubmitResultRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *QueryPipelineSubmitResultRequest
	GetOpUserId() *string
	SetSubmitId(v int64) *QueryPipelineSubmitResultRequest
	GetSubmitId() *int64
}

type QueryPipelineSubmitResultRequest struct {
	// The request context information.
	//
	// This parameter is required.
	Context *QueryPipelineSubmitResultRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
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
	// The submit detail ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s QueryPipelineSubmitResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultRequest) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultRequest) GetContext() *QueryPipelineSubmitResultRequestContext {
	return s.Context
}

func (s *QueryPipelineSubmitResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *QueryPipelineSubmitResultRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *QueryPipelineSubmitResultRequest) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *QueryPipelineSubmitResultRequest) SetContext(v *QueryPipelineSubmitResultRequestContext) *QueryPipelineSubmitResultRequest {
	s.Context = v
	return s
}

func (s *QueryPipelineSubmitResultRequest) SetOpTenantId(v int64) *QueryPipelineSubmitResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *QueryPipelineSubmitResultRequest) SetOpUserId(v string) *QueryPipelineSubmitResultRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryPipelineSubmitResultRequest) SetSubmitId(v int64) *QueryPipelineSubmitResultRequest {
	s.SubmitId = &v
	return s
}

func (s *QueryPipelineSubmitResultRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPipelineSubmitResultRequestContext struct {
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

func (s QueryPipelineSubmitResultRequestContext) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultRequestContext) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultRequestContext) GetEnv() *string {
	return s.Env
}

func (s *QueryPipelineSubmitResultRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *QueryPipelineSubmitResultRequestContext) SetEnv(v string) *QueryPipelineSubmitResultRequestContext {
	s.Env = &v
	return s
}

func (s *QueryPipelineSubmitResultRequestContext) SetProjectId(v int64) *QueryPipelineSubmitResultRequestContext {
	s.ProjectId = &v
	return s
}

func (s *QueryPipelineSubmitResultRequestContext) Validate() error {
	return dara.Validate(s)
}
