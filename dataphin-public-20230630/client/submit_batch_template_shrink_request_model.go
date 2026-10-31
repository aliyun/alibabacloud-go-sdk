// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitBatchTemplateShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitBatchTemplateShrinkRequest
	GetOpUserId() *string
	SetProjectId(v int64) *SubmitBatchTemplateShrinkRequest
	GetProjectId() *int64
	SetSubmitCommandShrink(v string) *SubmitBatchTemplateShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitBatchTemplateShrinkRequest struct {
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
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template submission instruction.
	//
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitBatchTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitBatchTemplateShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitBatchTemplateShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SubmitBatchTemplateShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitBatchTemplateShrinkRequest) SetOpTenantId(v int64) *SubmitBatchTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitBatchTemplateShrinkRequest) SetOpUserId(v string) *SubmitBatchTemplateShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitBatchTemplateShrinkRequest) SetProjectId(v int64) *SubmitBatchTemplateShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitBatchTemplateShrinkRequest) SetSubmitCommandShrink(v string) *SubmitBatchTemplateShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitBatchTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
