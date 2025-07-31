// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceApiShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommandShrink(v string) *ApplyDataServiceApiShrinkRequest
	GetApplyCommandShrink() *string
	SetOpTenantId(v int64) *ApplyDataServiceApiShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ApplyDataServiceApiShrinkRequest
	GetProjectId() *int32
}

type ApplyDataServiceApiShrinkRequest struct {
	// This parameter is required.
	ApplyCommandShrink *string `json:"ApplyCommand,omitempty" xml:"ApplyCommand,omitempty"`
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

func (s ApplyDataServiceApiShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiShrinkRequest) GetApplyCommandShrink() *string {
	return s.ApplyCommandShrink
}

func (s *ApplyDataServiceApiShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyDataServiceApiShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ApplyDataServiceApiShrinkRequest) SetApplyCommandShrink(v string) *ApplyDataServiceApiShrinkRequest {
	s.ApplyCommandShrink = &v
	return s
}

func (s *ApplyDataServiceApiShrinkRequest) SetOpTenantId(v int64) *ApplyDataServiceApiShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyDataServiceApiShrinkRequest) SetProjectId(v int32) *ApplyDataServiceApiShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyDataServiceApiShrinkRequest) Validate() error {
	return dara.Validate(s)
}
