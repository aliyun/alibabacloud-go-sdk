// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetMappingRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetMappingQuery(v *GetAssetMappingRelationsRequestAssetMappingQuery) *GetAssetMappingRelationsRequest
	GetAssetMappingQuery() *GetAssetMappingRelationsRequestAssetMappingQuery
	SetOpTenantId(v int64) *GetAssetMappingRelationsRequest
	GetOpTenantId() *int64
}

type GetAssetMappingRelationsRequest struct {
	AssetMappingQuery *GetAssetMappingRelationsRequestAssetMappingQuery `json:"AssetMappingQuery,omitempty" xml:"AssetMappingQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetAssetMappingRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsRequest) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsRequest) GetAssetMappingQuery() *GetAssetMappingRelationsRequestAssetMappingQuery {
	return s.AssetMappingQuery
}

func (s *GetAssetMappingRelationsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAssetMappingRelationsRequest) SetAssetMappingQuery(v *GetAssetMappingRelationsRequestAssetMappingQuery) *GetAssetMappingRelationsRequest {
	s.AssetMappingQuery = v
	return s
}

func (s *GetAssetMappingRelationsRequest) SetOpTenantId(v int64) *GetAssetMappingRelationsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAssetMappingRelationsRequest) Validate() error {
	if s.AssetMappingQuery != nil {
		if err := s.AssetMappingQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssetMappingRelationsRequestAssetMappingQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// INDEX
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1121
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VALID
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s GetAssetMappingRelationsRequestAssetMappingQuery) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsRequestAssetMappingQuery) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) GetAssetType() *string {
	return s.AssetType
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) GetGuid() *string {
	return s.Guid
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) GetRelationType() *string {
	return s.RelationType
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) SetAssetType(v string) *GetAssetMappingRelationsRequestAssetMappingQuery {
	s.AssetType = &v
	return s
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) SetGuid(v string) *GetAssetMappingRelationsRequestAssetMappingQuery {
	s.Guid = &v
	return s
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) SetRelationType(v string) *GetAssetMappingRelationsRequestAssetMappingQuery {
	s.RelationType = &v
	return s
}

func (s *GetAssetMappingRelationsRequestAssetMappingQuery) Validate() error {
	return dara.Validate(s)
}
