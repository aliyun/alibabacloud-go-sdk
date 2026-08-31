// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *ListPipelinesShrinkRequest
	GetContextShrink() *string
	SetListCommandShrink(v string) *ListPipelinesShrinkRequest
	GetListCommandShrink() *string
	SetOpTenantId(v int64) *ListPipelinesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListPipelinesShrinkRequest
	GetOpUserId() *string
}

type ListPipelinesShrinkRequest struct {
	// The request context.
	//
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The query parameters.
	//
	// This parameter is required.
	ListCommandShrink *string `json:"ListCommand,omitempty" xml:"ListCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListPipelinesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *ListPipelinesShrinkRequest) GetListCommandShrink() *string {
	return s.ListCommandShrink
}

func (s *ListPipelinesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPipelinesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListPipelinesShrinkRequest) SetContextShrink(v string) *ListPipelinesShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *ListPipelinesShrinkRequest) SetListCommandShrink(v string) *ListPipelinesShrinkRequest {
	s.ListCommandShrink = &v
	return s
}

func (s *ListPipelinesShrinkRequest) SetOpTenantId(v int64) *ListPipelinesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPipelinesShrinkRequest) SetOpUserId(v string) *ListPipelinesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListPipelinesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
