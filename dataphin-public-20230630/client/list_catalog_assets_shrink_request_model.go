// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogAssetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListCatalogAssetsQueryShrink(v string) *ListCatalogAssetsShrinkRequest
	GetListCatalogAssetsQueryShrink() *string
	SetOpTenantId(v int64) *ListCatalogAssetsShrinkRequest
	GetOpTenantId() *int64
}

type ListCatalogAssetsShrinkRequest struct {
	// This parameter is required.
	ListCatalogAssetsQueryShrink *string `json:"ListCatalogAssetsQuery,omitempty" xml:"ListCatalogAssetsQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListCatalogAssetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsShrinkRequest) GetListCatalogAssetsQueryShrink() *string {
	return s.ListCatalogAssetsQueryShrink
}

func (s *ListCatalogAssetsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListCatalogAssetsShrinkRequest) SetListCatalogAssetsQueryShrink(v string) *ListCatalogAssetsShrinkRequest {
	s.ListCatalogAssetsQueryShrink = &v
	return s
}

func (s *ListCatalogAssetsShrinkRequest) SetOpTenantId(v int64) *ListCatalogAssetsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListCatalogAssetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
