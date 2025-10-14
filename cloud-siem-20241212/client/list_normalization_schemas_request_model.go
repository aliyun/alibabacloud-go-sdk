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
	SetNormalizationSchemaType(v string) *ListNormalizationSchemasRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *ListNormalizationSchemasRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationSchemasRequest
	GetRoleFor() *int64
}

type ListNormalizationSchemasRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// entity。
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
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

func (s *ListNormalizationSchemasRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
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

func (s *ListNormalizationSchemasRequest) SetNormalizationSchemaType(v string) *ListNormalizationSchemasRequest {
	s.NormalizationSchemaType = &v
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
