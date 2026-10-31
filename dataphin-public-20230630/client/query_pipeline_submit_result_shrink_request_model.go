// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineSubmitResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *QueryPipelineSubmitResultShrinkRequest
	GetContextShrink() *string
	SetOpTenantId(v int64) *QueryPipelineSubmitResultShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *QueryPipelineSubmitResultShrinkRequest
	GetOpUserId() *string
	SetSubmitId(v int64) *QueryPipelineSubmitResultShrinkRequest
	GetSubmitId() *int64
}

type QueryPipelineSubmitResultShrinkRequest struct {
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
	// The submit detail ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s QueryPipelineSubmitResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineSubmitResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryPipelineSubmitResultShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *QueryPipelineSubmitResultShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *QueryPipelineSubmitResultShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *QueryPipelineSubmitResultShrinkRequest) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *QueryPipelineSubmitResultShrinkRequest) SetContextShrink(v string) *QueryPipelineSubmitResultShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *QueryPipelineSubmitResultShrinkRequest) SetOpTenantId(v int64) *QueryPipelineSubmitResultShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *QueryPipelineSubmitResultShrinkRequest) SetOpUserId(v string) *QueryPipelineSubmitResultShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryPipelineSubmitResultShrinkRequest) SetSubmitId(v int64) *QueryPipelineSubmitResultShrinkRequest {
	s.SubmitId = &v
	return s
}

func (s *QueryPipelineSubmitResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
