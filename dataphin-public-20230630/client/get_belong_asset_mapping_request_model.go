// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBelongAssetMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetMappingQuery(v *GetBelongAssetMappingRequestAssetMappingQuery) *GetBelongAssetMappingRequest
	GetAssetMappingQuery() *GetBelongAssetMappingRequestAssetMappingQuery
	SetOpTenantId(v int64) *GetBelongAssetMappingRequest
	GetOpTenantId() *int64
}

type GetBelongAssetMappingRequest struct {
	AssetMappingQuery *GetBelongAssetMappingRequestAssetMappingQuery `json:"AssetMappingQuery,omitempty" xml:"AssetMappingQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetBelongAssetMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingRequest) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingRequest) GetAssetMappingQuery() *GetBelongAssetMappingRequestAssetMappingQuery {
	return s.AssetMappingQuery
}

func (s *GetBelongAssetMappingRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetBelongAssetMappingRequest) SetAssetMappingQuery(v *GetBelongAssetMappingRequestAssetMappingQuery) *GetBelongAssetMappingRequest {
	s.AssetMappingQuery = v
	return s
}

func (s *GetBelongAssetMappingRequest) SetOpTenantId(v int64) *GetBelongAssetMappingRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetBelongAssetMappingRequest) Validate() error {
	if s.AssetMappingQuery != nil {
		if err := s.AssetMappingQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBelongAssetMappingRequestAssetMappingQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1121
	BelongGuid *string `json:"BelongGuid,omitempty" xml:"BelongGuid,omitempty"`
	// example:
	//
	// VALID
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s GetBelongAssetMappingRequestAssetMappingQuery) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingRequestAssetMappingQuery) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingRequestAssetMappingQuery) GetBelongGuid() *string {
	return s.BelongGuid
}

func (s *GetBelongAssetMappingRequestAssetMappingQuery) GetRelationType() *string {
	return s.RelationType
}

func (s *GetBelongAssetMappingRequestAssetMappingQuery) SetBelongGuid(v string) *GetBelongAssetMappingRequestAssetMappingQuery {
	s.BelongGuid = &v
	return s
}

func (s *GetBelongAssetMappingRequestAssetMappingQuery) SetRelationType(v string) *GetBelongAssetMappingRequestAssetMappingQuery {
	s.RelationType = &v
	return s
}

func (s *GetBelongAssetMappingRequestAssetMappingQuery) Validate() error {
	return dara.Validate(s)
}
