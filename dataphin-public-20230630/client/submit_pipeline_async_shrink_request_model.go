// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *SubmitPipelineAsyncShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *SubmitPipelineAsyncShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitPipelineAsyncShrinkRequest
	GetOpUserId() *string
	SetQueryIdShrink(v string) *SubmitPipelineAsyncShrinkRequest
	GetQueryIdShrink() *string
}

type SubmitPipelineAsyncShrinkRequest struct {
	// The request context information.
	//
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
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
	QueryIdShrink *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s SubmitPipelineAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *SubmitPipelineAsyncShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitPipelineAsyncShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitPipelineAsyncShrinkRequest) GetQueryIdShrink() *string {
	return s.QueryIdShrink
}

func (s *SubmitPipelineAsyncShrinkRequest) SetContextShrink(v string) *SubmitPipelineAsyncShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *SubmitPipelineAsyncShrinkRequest) SetOpTenantId(v int64) *SubmitPipelineAsyncShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitPipelineAsyncShrinkRequest) SetOpUserId(v string) *SubmitPipelineAsyncShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitPipelineAsyncShrinkRequest) SetQueryIdShrink(v string) *SubmitPipelineAsyncShrinkRequest {
	s.QueryIdShrink = &v
	return s
}

func (s *SubmitPipelineAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
