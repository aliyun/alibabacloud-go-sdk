// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetTypeAttributeCodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *GetAssetTypeAttributeCodesRequest
	GetAssetType() *string
	SetOpTenantId(v int64) *GetAssetTypeAttributeCodesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetAssetTypeAttributeCodesRequest
	GetOpUserId() *string
}

type GetAssetTypeAttributeCodesRequest struct {
	// The asset type filter. Valid values:
	//
	// - TABLE: table.
	//
	// - COLUMN: column.
	//
	// - INDEX: metric.
	//
	// - BIZ_INDEX: business metric.
	//
	// - API: API.
	//
	// - PAGE: dashboard.
	//
	// example:
	//
	// TABLE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetAssetTypeAttributeCodesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetTypeAttributeCodesRequest) GoString() string {
	return s.String()
}

func (s *GetAssetTypeAttributeCodesRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *GetAssetTypeAttributeCodesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAssetTypeAttributeCodesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetAssetTypeAttributeCodesRequest) SetAssetType(v string) *GetAssetTypeAttributeCodesRequest {
	s.AssetType = &v
	return s
}

func (s *GetAssetTypeAttributeCodesRequest) SetOpTenantId(v int64) *GetAssetTypeAttributeCodesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAssetTypeAttributeCodesRequest) SetOpUserId(v string) *GetAssetTypeAttributeCodesRequest {
	s.OpUserId = &v
	return s
}

func (s *GetAssetTypeAttributeCodesRequest) Validate() error {
	return dara.Validate(s)
}
