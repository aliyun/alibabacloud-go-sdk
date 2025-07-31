// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddProjectMemberShrinkRequest
	GetAddCommandShrink() *string
	SetId(v int64) *AddProjectMemberShrinkRequest
	GetId() *int64
	SetOpTenantId(v int64) *AddProjectMemberShrinkRequest
	GetOpTenantId() *int64
}

type AddProjectMemberShrinkRequest struct {
	// This parameter is required.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
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
}

func (s AddProjectMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMemberShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddProjectMemberShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *AddProjectMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddProjectMemberShrinkRequest) SetAddCommandShrink(v string) *AddProjectMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddProjectMemberShrinkRequest) SetId(v int64) *AddProjectMemberShrinkRequest {
	s.Id = &v
	return s
}

func (s *AddProjectMemberShrinkRequest) SetOpTenantId(v int64) *AddProjectMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddProjectMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
