// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDataServiceApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RevokeDataServiceApiShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *RevokeDataServiceApiShrinkRequest
	GetProjectId() *int32
	SetRevokeCommandShrink(v string) *RevokeDataServiceApiShrinkRequest
	GetRevokeCommandShrink() *string
}

type RevokeDataServiceApiShrinkRequest struct {
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
	// This parameter is required.
	RevokeCommandShrink *string `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty"`
}

func (s RevokeDataServiceApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeDataServiceApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeDataServiceApiShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RevokeDataServiceApiShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *RevokeDataServiceApiShrinkRequest) GetRevokeCommandShrink() *string {
	return s.RevokeCommandShrink
}

func (s *RevokeDataServiceApiShrinkRequest) SetOpTenantId(v int64) *RevokeDataServiceApiShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeDataServiceApiShrinkRequest) SetProjectId(v int32) *RevokeDataServiceApiShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeDataServiceApiShrinkRequest) SetRevokeCommandShrink(v string) *RevokeDataServiceApiShrinkRequest {
	s.RevokeCommandShrink = &v
	return s
}

func (s *RevokeDataServiceApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
