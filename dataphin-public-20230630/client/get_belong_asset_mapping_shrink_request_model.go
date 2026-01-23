// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBelongAssetMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetMappingQueryShrink(v string) *GetBelongAssetMappingShrinkRequest
	GetAssetMappingQueryShrink() *string
	SetOpTenantId(v int64) *GetBelongAssetMappingShrinkRequest
	GetOpTenantId() *int64
}

type GetBelongAssetMappingShrinkRequest struct {
	AssetMappingQueryShrink *string `json:"AssetMappingQuery,omitempty" xml:"AssetMappingQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetBelongAssetMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingShrinkRequest) GetAssetMappingQueryShrink() *string {
	return s.AssetMappingQueryShrink
}

func (s *GetBelongAssetMappingShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBelongAssetMappingShrinkRequest) SetAssetMappingQueryShrink(v string) *GetBelongAssetMappingShrinkRequest {
	s.AssetMappingQueryShrink = &v
	return s
}

func (s *GetBelongAssetMappingShrinkRequest) SetOpTenantId(v int64) *GetBelongAssetMappingShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBelongAssetMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
