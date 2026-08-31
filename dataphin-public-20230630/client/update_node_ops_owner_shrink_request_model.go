// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeOpsOwnerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandShrink(v string) *UpdateNodeOpsOwnerShrinkRequest
	GetCommandShrink() *string
	SetOpTenantId(v int64) *UpdateNodeOpsOwnerShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateNodeOpsOwnerShrinkRequest
	GetOpUserId() *string
}

type UpdateNodeOpsOwnerShrinkRequest struct {
	// The command for updating O&M owners.
	//
	// This parameter is required.
	CommandShrink *string `json:"Command,omitempty" xml:"Command,omitempty"`
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

func (s UpdateNodeOpsOwnerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeOpsOwnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeOpsOwnerShrinkRequest) GetCommandShrink() *string {
	return s.CommandShrink
}

func (s *UpdateNodeOpsOwnerShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateNodeOpsOwnerShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateNodeOpsOwnerShrinkRequest) SetCommandShrink(v string) *UpdateNodeOpsOwnerShrinkRequest {
	s.CommandShrink = &v
	return s
}

func (s *UpdateNodeOpsOwnerShrinkRequest) SetOpTenantId(v int64) *UpdateNodeOpsOwnerShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateNodeOpsOwnerShrinkRequest) SetOpUserId(v string) *UpdateNodeOpsOwnerShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateNodeOpsOwnerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
