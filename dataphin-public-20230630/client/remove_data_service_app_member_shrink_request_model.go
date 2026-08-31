// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataServiceAppMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveDataServiceAppMemberShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *RemoveDataServiceAppMemberShrinkRequest
	GetOpUserId() *string
	SetRemoveCommandShrink(v string) *RemoveDataServiceAppMemberShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveDataServiceAppMemberShrinkRequest struct {
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// Remove regular members from a data service application
	//
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveDataServiceAppMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataServiceAppMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveDataServiceAppMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveDataServiceAppMemberShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *RemoveDataServiceAppMemberShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveDataServiceAppMemberShrinkRequest) SetOpTenantId(v int64) *RemoveDataServiceAppMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveDataServiceAppMemberShrinkRequest) SetOpUserId(v string) *RemoveDataServiceAppMemberShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveDataServiceAppMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveDataServiceAppMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveDataServiceAppMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
