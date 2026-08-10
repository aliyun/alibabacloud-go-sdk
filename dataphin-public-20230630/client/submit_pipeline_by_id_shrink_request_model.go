// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *SubmitPipelineByIdShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *SubmitPipelineByIdShrinkRequest
	GetOpTenantId() *int64
	SetQueryIdShrink(v string) *SubmitPipelineByIdShrinkRequest
	GetQueryIdShrink() *string
}

type SubmitPipelineByIdShrinkRequest struct {
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
	// The ID used to query the pipeline task.
	//
	// This parameter is required.
	QueryIdShrink *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s SubmitPipelineByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitPipelineByIdShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *SubmitPipelineByIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitPipelineByIdShrinkRequest) GetQueryIdShrink() *string {
	return s.QueryIdShrink
}

func (s *SubmitPipelineByIdShrinkRequest) SetContextShrink(v string) *SubmitPipelineByIdShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *SubmitPipelineByIdShrinkRequest) SetOpTenantId(v int64) *SubmitPipelineByIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitPipelineByIdShrinkRequest) SetQueryIdShrink(v string) *SubmitPipelineByIdShrinkRequest {
	s.QueryIdShrink = &v
	return s
}

func (s *SubmitPipelineByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
