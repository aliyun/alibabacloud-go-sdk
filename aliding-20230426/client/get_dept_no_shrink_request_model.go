// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDeptNoShrinkRequest
	GetTenantContextShrink() *string
	SetDeptId(v string) *GetDeptNoShrinkRequest
	GetDeptId() *string
}

type GetDeptNoShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s GetDeptNoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeptNoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDeptNoShrinkRequest) GetDeptId() *string {
	return s.DeptId
}

func (s *GetDeptNoShrinkRequest) SetTenantContextShrink(v string) *GetDeptNoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDeptNoShrinkRequest) SetDeptId(v string) *GetDeptNoShrinkRequest {
	s.DeptId = &v
	return s
}

func (s *GetDeptNoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
