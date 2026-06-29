// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogAssetDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetCatalogAssetDetailsQuery(v *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) *GetCatalogAssetDetailsRequest
	GetGetCatalogAssetDetailsQuery() *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery
	SetOpTenantId(v int64) *GetCatalogAssetDetailsRequest
	GetOpTenantId() *int64
}

type GetCatalogAssetDetailsRequest struct {
	// The paging query for the listing asset catalog list.
	//
	// This parameter is required.
	GetCatalogAssetDetailsQuery *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery `json:"GetCatalogAssetDetailsQuery,omitempty" xml:"GetCatalogAssetDetailsQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetCatalogAssetDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsRequest) GetGetCatalogAssetDetailsQuery() *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery {
	return s.GetCatalogAssetDetailsQuery
}

func (s *GetCatalogAssetDetailsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetCatalogAssetDetailsRequest) SetGetCatalogAssetDetailsQuery(v *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) *GetCatalogAssetDetailsRequest {
	s.GetCatalogAssetDetailsQuery = v
	return s
}

func (s *GetCatalogAssetDetailsRequest) SetOpTenantId(v int64) *GetCatalogAssetDetailsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetCatalogAssetDetailsRequest) Validate() error {
	if s.GetCatalogAssetDetailsQuery != nil {
		if err := s.GetCatalogAssetDetailsQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery struct {
	// The asset GUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dp_ds_table.300023201.7311626611751680256.load_test.abc
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// Specifies whether to include the field list. Default value: false. Valid values:
	//
	// - false: Not included.
	//
	// - true: Included.
	IncludeColumns *bool `json:"IncludeColumns,omitempty" xml:"IncludeColumns,omitempty"`
	// Specifies whether to include detailed data. Default value: false. Valid values:
	//
	// - false: Not included.
	//
	// - true: Included.
	IncludeDetailedAttributes *bool `json:"IncludeDetailedAttributes,omitempty" xml:"IncludeDetailedAttributes,omitempty"`
}

func (s GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) GetGuid() *string {
	return s.Guid
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) GetIncludeColumns() *bool {
	return s.IncludeColumns
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) GetIncludeDetailedAttributes() *bool {
	return s.IncludeDetailedAttributes
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) SetGuid(v string) *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery {
	s.Guid = &v
	return s
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) SetIncludeColumns(v bool) *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery {
	s.IncludeColumns = &v
	return s
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) SetIncludeDetailedAttributes(v bool) *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery {
	s.IncludeDetailedAttributes = &v
	return s
}

func (s *GetCatalogAssetDetailsRequestGetCatalogAssetDetailsQuery) Validate() error {
	return dara.Validate(s)
}
