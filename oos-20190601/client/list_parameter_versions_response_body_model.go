// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedBy(v string) *ListParameterVersionsResponseBody
	GetCreatedBy() *string
	SetCreatedDate(v string) *ListParameterVersionsResponseBody
	GetCreatedDate() *string
	SetDescription(v string) *ListParameterVersionsResponseBody
	GetDescription() *string
	SetId(v string) *ListParameterVersionsResponseBody
	GetId() *string
	SetMaxResults(v int32) *ListParameterVersionsResponseBody
	GetMaxResults() *int32
	SetName(v string) *ListParameterVersionsResponseBody
	GetName() *string
	SetNextToken(v string) *ListParameterVersionsResponseBody
	GetNextToken() *string
	SetParameterVersions(v []*ListParameterVersionsResponseBodyParameterVersions) *ListParameterVersionsResponseBody
	GetParameterVersions() []*ListParameterVersionsResponseBodyParameterVersions
	SetRequestId(v string) *ListParameterVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListParameterVersionsResponseBody
	GetTotalCount() *int32
	SetType(v string) *ListParameterVersionsResponseBody
	GetType() *string
}

type ListParameterVersionsResponseBody struct {
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-07T11:37:29Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-a483b520e0axxxxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the version of the common parameter.
	ParameterVersions []*ListParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FD08D89D-B6C8-4AA2-A2B4-521D3F4A39FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParameterVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBody) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListParameterVersionsResponseBody) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListParameterVersionsResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ListParameterVersionsResponseBody) GetId() *string {
	return s.Id
}

func (s *ListParameterVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListParameterVersionsResponseBody) GetName() *string {
	return s.Name
}

func (s *ListParameterVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListParameterVersionsResponseBody) GetParameterVersions() []*ListParameterVersionsResponseBodyParameterVersions {
	return s.ParameterVersions
}

func (s *ListParameterVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParameterVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParameterVersionsResponseBody) GetType() *string {
	return s.Type
}

func (s *ListParameterVersionsResponseBody) SetCreatedBy(v string) *ListParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetCreatedDate(v string) *ListParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetDescription(v string) *ListParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetId(v string) *ListParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetMaxResults(v int32) *ListParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetName(v string) *ListParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetNextToken(v string) *ListParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetParameterVersions(v []*ListParameterVersionsResponseBodyParameterVersions) *ListParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

func (s *ListParameterVersionsResponseBody) SetRequestId(v string) *ListParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetTotalCount(v int32) *ListParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetType(v string) *ListParameterVersionsResponseBody {
	s.Type = &v
	return s
}

func (s *ListParameterVersionsResponseBody) Validate() error {
	if s.ParameterVersions != nil {
		for _, item := range s.ParameterVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParameterVersionsResponseBodyParameterVersions struct {
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was last updated.
	//
	// example:
	//
	// 2020-09-07T11:37:29Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// MyParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParameterVersionsResponseBodyParameterVersions) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyParameterVersions) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *ListParameterVersionsResponseBodyParameterVersions) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListParameterVersionsResponseBodyParameterVersions) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListParameterVersionsResponseBodyParameterVersions) GetValue() *string {
	return s.Value
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) Validate() error {
	return dara.Validate(s)
}
