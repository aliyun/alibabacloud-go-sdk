// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationFieldsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationFieldsRequest
	GetMaxResults() *int32
	SetName(v string) *ListNormalizationFieldsRequest
	GetName() *string
	SetNextToken(v string) *ListNormalizationFieldsRequest
	GetNextToken() *string
	SetNormalizationFieldSource(v string) *ListNormalizationFieldsRequest
	GetNormalizationFieldSource() *string
	SetNormalizationSchemaType(v string) *ListNormalizationFieldsRequest
	GetNormalizationSchemaType() *string
	SetRegionId(v string) *ListNormalizationFieldsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationFieldsRequest
	GetRoleFor() *int64
}

type ListNormalizationFieldsRequest struct {
	// The language of the response. Valid values:
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
	// The field name.
	//
	// example:
	//
	// category
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether a next query token exists. You do not need to specify this parameter for the first query or if no next query exists. If a next query exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// native
	NormalizationFieldSource *string `json:"NormalizationFieldSource,omitempty" xml:"NormalizationFieldSource,omitempty"`
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationFieldsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationFieldsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationFieldsRequest) GetName() *string {
	return s.Name
}

func (s *ListNormalizationFieldsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationFieldsRequest) GetNormalizationFieldSource() *string {
	return s.NormalizationFieldSource
}

func (s *ListNormalizationFieldsRequest) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *ListNormalizationFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationFieldsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationFieldsRequest) SetLang(v string) *ListNormalizationFieldsRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetMaxResults(v int32) *ListNormalizationFieldsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetName(v string) *ListNormalizationFieldsRequest {
	s.Name = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetNextToken(v string) *ListNormalizationFieldsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetNormalizationFieldSource(v string) *ListNormalizationFieldsRequest {
	s.NormalizationFieldSource = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetNormalizationSchemaType(v string) *ListNormalizationFieldsRequest {
	s.NormalizationSchemaType = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetRegionId(v string) *ListNormalizationFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationFieldsRequest) SetRoleFor(v int64) *ListNormalizationFieldsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationFieldsRequest) Validate() error {
	return dara.Validate(s)
}
