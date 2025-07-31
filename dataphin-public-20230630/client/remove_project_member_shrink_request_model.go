// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RemoveProjectMemberShrinkRequest
	GetId() *int64
	SetOpTenantId(v int64) *RemoveProjectMemberShrinkRequest
	GetOpTenantId() *int64
	SetRemoveCommandShrink(v string) *RemoveProjectMemberShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveProjectMemberShrinkRequest struct {
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
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveProjectMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveProjectMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveProjectMemberShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveProjectMemberShrinkRequest) SetId(v int64) *RemoveProjectMemberShrinkRequest {
	s.Id = &v
	return s
}

func (s *RemoveProjectMemberShrinkRequest) SetOpTenantId(v int64) *RemoveProjectMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveProjectMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveProjectMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveProjectMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
