// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceProjectMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommandShrink(v string) *AddDataServiceProjectMemberShrinkRequest
	GetAddCommandShrink() *string
	SetOpTenantId(v int64) *AddDataServiceProjectMemberShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *AddDataServiceProjectMemberShrinkRequest
	GetProjectId() *int32
}

type AddDataServiceProjectMemberShrinkRequest struct {
	// This parameter is required.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddDataServiceProjectMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberShrinkRequest) GetAddCommandShrink() *string {
	return s.AddCommandShrink
}

func (s *AddDataServiceProjectMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddDataServiceProjectMemberShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *AddDataServiceProjectMemberShrinkRequest) SetAddCommandShrink(v string) *AddDataServiceProjectMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddDataServiceProjectMemberShrinkRequest) SetOpTenantId(v int64) *AddDataServiceProjectMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddDataServiceProjectMemberShrinkRequest) SetProjectId(v int32) *AddDataServiceProjectMemberShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *AddDataServiceProjectMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
