// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateBatchTemplateShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateBatchTemplateShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateBatchTemplateShrinkRequest
	GetOpUserId() *string
	SetProjectId(v int64) *CreateBatchTemplateShrinkRequest
	GetProjectId() *int64
}

type CreateBatchTemplateShrinkRequest struct {
	// The command to create a template.
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
}

func (s CreateBatchTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateBatchTemplateShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBatchTemplateShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateBatchTemplateShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBatchTemplateShrinkRequest) SetCreateCommandShrink(v string) *CreateBatchTemplateShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBatchTemplateShrinkRequest) SetOpTenantId(v int64) *CreateBatchTemplateShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBatchTemplateShrinkRequest) SetOpUserId(v string) *CreateBatchTemplateShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBatchTemplateShrinkRequest) SetProjectId(v int64) *CreateBatchTemplateShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateBatchTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
