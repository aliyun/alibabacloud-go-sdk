// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTemplateShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateBatchTemplateShrinkRequest
	GetOpUserId() *string
	SetProjectId(v int64) *UpdateBatchTemplateShrinkRequest
	GetProjectId() *int64
	SetUpdateCommandShrink(v string) *UpdateBatchTemplateShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBatchTemplateShrinkRequest struct {
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
	// The template update command.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBatchTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTemplateShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBatchTemplateShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBatchTemplateShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBatchTemplateShrinkRequest) SetOpTenantId(v int64) *UpdateBatchTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTemplateShrinkRequest) SetOpUserId(v string) *UpdateBatchTemplateShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateBatchTemplateShrinkRequest) SetProjectId(v int64) *UpdateBatchTemplateShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBatchTemplateShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBatchTemplateShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBatchTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
