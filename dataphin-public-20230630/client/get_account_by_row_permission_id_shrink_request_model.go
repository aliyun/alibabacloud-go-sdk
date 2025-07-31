// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountByRowPermissionIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetAccountByRowPermissionIdQueryShrink(v string) *GetAccountByRowPermissionIdShrinkRequest
	GetGetAccountByRowPermissionIdQueryShrink() *string
	SetOpTenantId(v int64) *GetAccountByRowPermissionIdShrinkRequest
	GetOpTenantId() *int64
}

type GetAccountByRowPermissionIdShrinkRequest struct {
	// This parameter is required.
	GetAccountByRowPermissionIdQueryShrink *string `json:"GetAccountByRowPermissionIdQuery,omitempty" xml:"GetAccountByRowPermissionIdQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetAccountByRowPermissionIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdShrinkRequest) GetGetAccountByRowPermissionIdQueryShrink() *string {
	return s.GetAccountByRowPermissionIdQueryShrink
}

func (s *GetAccountByRowPermissionIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAccountByRowPermissionIdShrinkRequest) SetGetAccountByRowPermissionIdQueryShrink(v string) *GetAccountByRowPermissionIdShrinkRequest {
	s.GetAccountByRowPermissionIdQueryShrink = &v
	return s
}

func (s *GetAccountByRowPermissionIdShrinkRequest) SetOpTenantId(v int64) *GetAccountByRowPermissionIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAccountByRowPermissionIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
