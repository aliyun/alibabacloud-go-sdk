// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *GetPipelineByIdShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *GetPipelineByIdShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetPipelineByIdShrinkRequest
	GetOpUserId() *string
	SetQueryIdShrink(v string) *GetPipelineByIdShrinkRequest
	GetQueryIdShrink() *string
}

type GetPipelineByIdShrinkRequest struct {
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
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID used to query the pipeline node.
	//
	// This parameter is required.
	QueryIdShrink *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetPipelineByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineByIdShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *GetPipelineByIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetPipelineByIdShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetPipelineByIdShrinkRequest) GetQueryIdShrink() *string {
	return s.QueryIdShrink
}

func (s *GetPipelineByIdShrinkRequest) SetContextShrink(v string) *GetPipelineByIdShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *GetPipelineByIdShrinkRequest) SetOpTenantId(v int64) *GetPipelineByIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPipelineByIdShrinkRequest) SetOpUserId(v string) *GetPipelineByIdShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetPipelineByIdShrinkRequest) SetQueryIdShrink(v string) *GetPipelineByIdShrinkRequest {
	s.QueryIdShrink = &v
	return s
}

func (s *GetPipelineByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
