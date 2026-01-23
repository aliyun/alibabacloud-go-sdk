// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetMappingRelationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetMappingQueryShrink(v string) *GetAssetMappingRelationsShrinkRequest
	GetAssetMappingQueryShrink() *string
	SetOpTenantId(v int64) *GetAssetMappingRelationsShrinkRequest
	GetOpTenantId() *int64
}

type GetAssetMappingRelationsShrinkRequest struct {
	AssetMappingQueryShrink *string `json:"AssetMappingQuery,omitempty" xml:"AssetMappingQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetAssetMappingRelationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsShrinkRequest) GetAssetMappingQueryShrink() *string {
	return s.AssetMappingQueryShrink
}

func (s *GetAssetMappingRelationsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAssetMappingRelationsShrinkRequest) SetAssetMappingQueryShrink(v string) *GetAssetMappingRelationsShrinkRequest {
	s.AssetMappingQueryShrink = &v
	return s
}

func (s *GetAssetMappingRelationsShrinkRequest) SetOpTenantId(v int64) *GetAssetMappingRelationsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAssetMappingRelationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
