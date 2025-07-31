// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListRowPermissionShrinkRequest
	GetOpTenantId() *int64
	SetPageRowPermissionQueryShrink(v string) *ListRowPermissionShrinkRequest
	GetPageRowPermissionQueryShrink() *string
}

type ListRowPermissionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PageRowPermissionQueryShrink *string `json:"PageRowPermissionQuery,omitempty" xml:"PageRowPermissionQuery,omitempty"`
}

func (s ListRowPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRowPermissionShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRowPermissionShrinkRequest) GetPageRowPermissionQueryShrink() *string {
	return s.PageRowPermissionQueryShrink
}

func (s *ListRowPermissionShrinkRequest) SetOpTenantId(v int64) *ListRowPermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRowPermissionShrinkRequest) SetPageRowPermissionQueryShrink(v string) *ListRowPermissionShrinkRequest {
	s.PageRowPermissionQueryShrink = &v
	return s
}

func (s *ListRowPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
