// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncId(v int64) *GetPipelineAsyncResultRequest
	GetAsyncId() *int64
	SetContext(v *GetPipelineAsyncResultRequestContext) *GetPipelineAsyncResultRequest
	GetContext() *GetPipelineAsyncResultRequestContext
	SetOpTenantId(v int64) *GetPipelineAsyncResultRequest
	GetOpTenantId() *int64
}

type GetPipelineAsyncResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	AsyncId *int64 `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	// This parameter is required.
	Context *GetPipelineAsyncResultRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPipelineAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultRequest) GetAsyncId() *int64 {
	return s.AsyncId
}

func (s *GetPipelineAsyncResultRequest) GetContext() *GetPipelineAsyncResultRequestContext {
	return s.Context
}

func (s *GetPipelineAsyncResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPipelineAsyncResultRequest) SetAsyncId(v int64) *GetPipelineAsyncResultRequest {
	s.AsyncId = &v
	return s
}

func (s *GetPipelineAsyncResultRequest) SetContext(v *GetPipelineAsyncResultRequestContext) *GetPipelineAsyncResultRequest {
	s.Context = v
	return s
}

func (s *GetPipelineAsyncResultRequest) SetOpTenantId(v int64) *GetPipelineAsyncResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPipelineAsyncResultRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineAsyncResultRequestContext struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPipelineAsyncResultRequestContext) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultRequestContext) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultRequestContext) GetEnv() *string {
	return s.Env
}

func (s *GetPipelineAsyncResultRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPipelineAsyncResultRequestContext) SetEnv(v string) *GetPipelineAsyncResultRequestContext {
	s.Env = &v
	return s
}

func (s *GetPipelineAsyncResultRequestContext) SetProjectId(v int64) *GetPipelineAsyncResultRequestContext {
	s.ProjectId = &v
	return s
}

func (s *GetPipelineAsyncResultRequestContext) Validate() error {
	return dara.Validate(s)
}
