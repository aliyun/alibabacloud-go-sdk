// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishKgSchemaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishKgSchemaShrinkRequest
	GetOpTenantId() *int64
	SetPublishCommandShrink(v string) *PublishKgSchemaShrinkRequest
	GetPublishCommandShrink() *string
	SetWorkspaceId(v string) *PublishKgSchemaShrinkRequest
	GetWorkspaceId() *string
}

type PublishKgSchemaShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The publish command and its details.
	//
	// This parameter is required.
	PublishCommandShrink *string `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PublishKgSchemaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishKgSchemaShrinkRequest) GetPublishCommandShrink() *string {
	return s.PublishCommandShrink
}

func (s *PublishKgSchemaShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PublishKgSchemaShrinkRequest) SetOpTenantId(v int64) *PublishKgSchemaShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishKgSchemaShrinkRequest) SetPublishCommandShrink(v string) *PublishKgSchemaShrinkRequest {
	s.PublishCommandShrink = &v
	return s
}

func (s *PublishKgSchemaShrinkRequest) SetWorkspaceId(v string) *PublishKgSchemaShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PublishKgSchemaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
