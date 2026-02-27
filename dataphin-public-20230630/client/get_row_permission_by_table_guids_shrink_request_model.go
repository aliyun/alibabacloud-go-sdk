// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRowPermissionByTableGuidsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetRowPermissionByTableGuidsQueryShrink(v string) *GetRowPermissionByTableGuidsShrinkRequest
	GetGetRowPermissionByTableGuidsQueryShrink() *string
	SetOpTenantId(v int64) *GetRowPermissionByTableGuidsShrinkRequest
	GetOpTenantId() *int64
}

type GetRowPermissionByTableGuidsShrinkRequest struct {
	// This parameter is required.
	GetRowPermissionByTableGuidsQueryShrink *string `json:"GetRowPermissionByTableGuidsQuery,omitempty" xml:"GetRowPermissionByTableGuidsQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetRowPermissionByTableGuidsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRowPermissionByTableGuidsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRowPermissionByTableGuidsShrinkRequest) GetGetRowPermissionByTableGuidsQueryShrink() *string {
	return s.GetRowPermissionByTableGuidsQueryShrink
}

func (s *GetRowPermissionByTableGuidsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetRowPermissionByTableGuidsShrinkRequest) SetGetRowPermissionByTableGuidsQueryShrink(v string) *GetRowPermissionByTableGuidsShrinkRequest {
	s.GetRowPermissionByTableGuidsQueryShrink = &v
	return s
}

func (s *GetRowPermissionByTableGuidsShrinkRequest) SetOpTenantId(v int64) *GetRowPermissionByTableGuidsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetRowPermissionByTableGuidsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
