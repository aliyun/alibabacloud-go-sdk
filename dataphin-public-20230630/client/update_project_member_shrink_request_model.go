// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateProjectMemberShrinkRequest
	GetId() *int64
	SetOpTenantId(v int64) *UpdateProjectMemberShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateProjectMemberShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateProjectMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 711833
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateProjectMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateProjectMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateProjectMemberShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateProjectMemberShrinkRequest) SetId(v int64) *UpdateProjectMemberShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectMemberShrinkRequest) SetOpTenantId(v int64) *UpdateProjectMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateProjectMemberShrinkRequest) SetUpdateCommandShrink(v string) *UpdateProjectMemberShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateProjectMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
