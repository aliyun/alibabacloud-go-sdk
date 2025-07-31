// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionByUserIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListRowPermissionByUserIdQueryShrink(v string) *ListRowPermissionByUserIdShrinkRequest
	GetListRowPermissionByUserIdQueryShrink() *string
	SetOpTenantId(v int64) *ListRowPermissionByUserIdShrinkRequest
	GetOpTenantId() *int64
}

type ListRowPermissionByUserIdShrinkRequest struct {
	// This parameter is required.
	ListRowPermissionByUserIdQueryShrink *string `json:"ListRowPermissionByUserIdQuery,omitempty" xml:"ListRowPermissionByUserIdQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListRowPermissionByUserIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdShrinkRequest) GetListRowPermissionByUserIdQueryShrink() *string {
	return s.ListRowPermissionByUserIdQueryShrink
}

func (s *ListRowPermissionByUserIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRowPermissionByUserIdShrinkRequest) SetListRowPermissionByUserIdQueryShrink(v string) *ListRowPermissionByUserIdShrinkRequest {
	s.ListRowPermissionByUserIdQueryShrink = &v
	return s
}

func (s *ListRowPermissionByUserIdShrinkRequest) SetOpTenantId(v int64) *ListRowPermissionByUserIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRowPermissionByUserIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
