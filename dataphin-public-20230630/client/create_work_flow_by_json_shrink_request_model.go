// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkFlowByJsonShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextShrink(v string) *CreateWorkFlowByJsonShrinkRequest
	GetContextShrink() *string
	SetCreateCommandShrink(v string) *CreateWorkFlowByJsonShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateWorkFlowByJsonShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateWorkFlowByJsonShrinkRequest
	GetOpUserId() *string
}

type CreateWorkFlowByJsonShrinkRequest struct {
	// The request context information.
	//
	// This parameter is required.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The JSON script command for creating a workflow.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateWorkFlowByJsonShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *CreateWorkFlowByJsonShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateWorkFlowByJsonShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateWorkFlowByJsonShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateWorkFlowByJsonShrinkRequest) SetContextShrink(v string) *CreateWorkFlowByJsonShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *CreateWorkFlowByJsonShrinkRequest) SetCreateCommandShrink(v string) *CreateWorkFlowByJsonShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateWorkFlowByJsonShrinkRequest) SetOpTenantId(v int64) *CreateWorkFlowByJsonShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateWorkFlowByJsonShrinkRequest) SetOpUserId(v string) *CreateWorkFlowByJsonShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateWorkFlowByJsonShrinkRequest) Validate() error {
	return dara.Validate(s)
}
