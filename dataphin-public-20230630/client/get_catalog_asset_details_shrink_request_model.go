// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogAssetDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetCatalogAssetDetailsQueryShrink(v string) *GetCatalogAssetDetailsShrinkRequest
	GetGetCatalogAssetDetailsQueryShrink() *string
	SetOpTenantId(v int64) *GetCatalogAssetDetailsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetCatalogAssetDetailsShrinkRequest
	GetOpUserId() *string
}

type GetCatalogAssetDetailsShrinkRequest struct {
	// The paged query for the listed asset catalog.
	//
	// This parameter is required.
	GetCatalogAssetDetailsQueryShrink *string `json:"GetCatalogAssetDetailsQuery,omitempty" xml:"GetCatalogAssetDetailsQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetCatalogAssetDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsShrinkRequest) GetGetCatalogAssetDetailsQueryShrink() *string {
	return s.GetCatalogAssetDetailsQueryShrink
}

func (s *GetCatalogAssetDetailsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetCatalogAssetDetailsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetCatalogAssetDetailsShrinkRequest) SetGetCatalogAssetDetailsQueryShrink(v string) *GetCatalogAssetDetailsShrinkRequest {
	s.GetCatalogAssetDetailsQueryShrink = &v
	return s
}

func (s *GetCatalogAssetDetailsShrinkRequest) SetOpTenantId(v int64) *GetCatalogAssetDetailsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetCatalogAssetDetailsShrinkRequest) SetOpUserId(v string) *GetCatalogAssetDetailsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetCatalogAssetDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
