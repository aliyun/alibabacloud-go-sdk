// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParameterVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedBy(v string) *ListSecretParameterVersionsResponseBody
	GetCreatedBy() *string
	SetCreatedDate(v string) *ListSecretParameterVersionsResponseBody
	GetCreatedDate() *string
	SetDescription(v string) *ListSecretParameterVersionsResponseBody
	GetDescription() *string
	SetId(v string) *ListSecretParameterVersionsResponseBody
	GetId() *string
	SetMaxResults(v int32) *ListSecretParameterVersionsResponseBody
	GetMaxResults() *int32
	SetName(v string) *ListSecretParameterVersionsResponseBody
	GetName() *string
	SetNextToken(v string) *ListSecretParameterVersionsResponseBody
	GetNextToken() *string
	SetParameterVersions(v []*ListSecretParameterVersionsResponseBodyParameterVersions) *ListSecretParameterVersionsResponseBody
	GetParameterVersions() []*ListSecretParameterVersionsResponseBodyParameterVersions
	SetRequestId(v string) *ListSecretParameterVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSecretParameterVersionsResponseBody
	GetTotalCount() *int32
	SetType(v string) *ListSecretParameterVersionsResponseBody
	GetType() *string
}

type ListSecretParameterVersionsResponseBody struct {
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the version of the encryption parameter.
	ParameterVersions []*ListSecretParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DBA6E6C8-F75D-41DE-AFF5-1FA03F551CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSecretParameterVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListSecretParameterVersionsResponseBody) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListSecretParameterVersionsResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ListSecretParameterVersionsResponseBody) GetId() *string {
	return s.Id
}

func (s *ListSecretParameterVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecretParameterVersionsResponseBody) GetName() *string {
	return s.Name
}

func (s *ListSecretParameterVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecretParameterVersionsResponseBody) GetParameterVersions() []*ListSecretParameterVersionsResponseBodyParameterVersions {
	return s.ParameterVersions
}

func (s *ListSecretParameterVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretParameterVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecretParameterVersionsResponseBody) GetType() *string {
	return s.Type
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedBy(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedDate(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetDescription(v string) *ListSecretParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetId(v string) *ListSecretParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetMaxResults(v int32) *ListSecretParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetName(v string) *ListSecretParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetNextToken(v string) *ListSecretParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetParameterVersions(v []*ListSecretParameterVersionsResponseBodyParameterVersions) *ListSecretParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetRequestId(v string) *ListSecretParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetTotalCount(v int32) *ListSecretParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetType(v string) *ListSecretParameterVersionsResponseBody {
	s.Type = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecretParameterVersionsResponseBodyParameterVersions struct {
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) GetValue() *string {
	return s.Value
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) Validate() error {
	return dara.Validate(s)
}
