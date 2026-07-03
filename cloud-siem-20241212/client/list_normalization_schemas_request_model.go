// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationSchemasRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationSchemasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationSchemasRequest
	GetNextToken() *string
	SetNormalizationCategoryId(v string) *ListNormalizationSchemasRequest
	GetNormalizationCategoryId() *string
	SetNormalizationFieldSource(v string) *ListNormalizationSchemasRequest
	GetNormalizationFieldSource() *string
	SetNormalizationSchemaType(v string) *ListNormalizationSchemasRequest
	GetNormalizationSchemaType() *string
	SetNormalizationSecurityDomainId(v string) *ListNormalizationSchemasRequest
	GetNormalizationSecurityDomainId() *string
	SetRegionId(v string) *ListNormalizationSchemasRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationSchemasRequest
	GetRoleFor() *int64
}

type ListNormalizationSchemasRequest struct {
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the normalization rule category.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// The field source filter. Valid values: normalized / native.
	//
	// example:
	//
	// normalized
	NormalizationFieldSource *string `json:"NormalizationFieldSource,omitempty" xml:"NormalizationFieldSource,omitempty"`
	// The normalization schema type. Valid values:
	//
	// - log: log.
	//
	// - entity: entity.
	//
	// example:
	//
	// entity
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The security domain ID.
	//
	// example:
	//
	// NETWORK_AND_WEB_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: the asset is in the Chinese mainland.
	//
	// - ap-southeast-1: the asset is outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator switches to when viewing as another member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationSchemasRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationSchemasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationSchemasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationSchemasRequest) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationSchemasRequest) GetNormalizationFieldSource() *string {
	return s.NormalizationFieldSource
}

func (s *ListNormalizationSchemasRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *ListNormalizationSchemasRequest) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *ListNormalizationSchemasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationSchemasRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationSchemasRequest) SetLang(v string) *ListNormalizationSchemasRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetMaxResults(v int32) *ListNormalizationSchemasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetNextToken(v string) *ListNormalizationSchemasRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetNormalizationCategoryId(v string) *ListNormalizationSchemasRequest {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetNormalizationFieldSource(v string) *ListNormalizationSchemasRequest {
	s.NormalizationFieldSource = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetNormalizationSchemaType(v string) *ListNormalizationSchemasRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetNormalizationSecurityDomainId(v string) *ListNormalizationSchemasRequest {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetRegionId(v string) *ListNormalizationSchemasRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationSchemasRequest) SetRoleFor(v int64) *ListNormalizationSchemasRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationSchemasRequest) Validate() error {
	return dara.Validate(s)
}
