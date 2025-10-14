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
	SetRegionId(v string) *ListNormalizationFieldsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationFieldsRequest
	GetRoleFor() *int64
}

type ListNormalizationFieldsRequest struct {
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
	// category。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
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
