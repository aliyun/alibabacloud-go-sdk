// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommandShrink(v string) *ApplyDataServiceAppShrinkRequest
	GetApplyCommandShrink() *string
	SetOpTenantId(v int64) *ApplyDataServiceAppShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ApplyDataServiceAppShrinkRequest
	GetProjectId() *int32
}

type ApplyDataServiceAppShrinkRequest struct {
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

func (s ApplyDataServiceAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceAppShrinkRequest) GetApplyCommandShrink() *string {
	return s.ApplyCommandShrink
}

func (s *ApplyDataServiceAppShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyDataServiceAppShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ApplyDataServiceAppShrinkRequest) SetApplyCommandShrink(v string) *ApplyDataServiceAppShrinkRequest {
	s.ApplyCommandShrink = &v
	return s
}

func (s *ApplyDataServiceAppShrinkRequest) SetOpTenantId(v int64) *ApplyDataServiceAppShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyDataServiceAppShrinkRequest) SetProjectId(v int32) *ApplyDataServiceAppShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyDataServiceAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
