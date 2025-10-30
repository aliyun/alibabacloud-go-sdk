// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineAsyncResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncId(v int64) *GetPipelineAsyncResultShrinkRequest
	GetAsyncId() *int64
	SetContextShrink(v string) *GetPipelineAsyncResultShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *GetPipelineAsyncResultShrinkRequest
	GetOpTenantId() *int64
}

type GetPipelineAsyncResultShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	AsyncId *int64 `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPipelineAsyncResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultShrinkRequest) GetAsyncId() *int64 {
	return s.AsyncId
}

func (s *GetPipelineAsyncResultShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *GetPipelineAsyncResultShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPipelineAsyncResultShrinkRequest) SetAsyncId(v int64) *GetPipelineAsyncResultShrinkRequest {
	s.AsyncId = &v
	return s
}

func (s *GetPipelineAsyncResultShrinkRequest) SetContextShrink(v string) *GetPipelineAsyncResultShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *GetPipelineAsyncResultShrinkRequest) SetOpTenantId(v int64) *GetPipelineAsyncResultShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPipelineAsyncResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
