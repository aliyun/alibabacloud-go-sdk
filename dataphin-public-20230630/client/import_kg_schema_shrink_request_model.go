// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKgSchemaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportCommandShrink(v string) *ImportKgSchemaShrinkRequest
	GetImportCommandShrink() *string
	SetOpTenantId(v int64) *ImportKgSchemaShrinkRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *ImportKgSchemaShrinkRequest
	GetWorkspaceId() *string
}

type ImportKgSchemaShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	ImportCommandShrink *string `json:"ImportCommand,omitempty" xml:"ImportCommand,omitempty"`
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
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ImportKgSchemaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaShrinkRequest) GetImportCommandShrink() *string {
	return s.ImportCommandShrink
}

func (s *ImportKgSchemaShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ImportKgSchemaShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ImportKgSchemaShrinkRequest) SetImportCommandShrink(v string) *ImportKgSchemaShrinkRequest {
	s.ImportCommandShrink = &v
	return s
}

func (s *ImportKgSchemaShrinkRequest) SetOpTenantId(v int64) *ImportKgSchemaShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ImportKgSchemaShrinkRequest) SetWorkspaceId(v string) *ImportKgSchemaShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ImportKgSchemaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
