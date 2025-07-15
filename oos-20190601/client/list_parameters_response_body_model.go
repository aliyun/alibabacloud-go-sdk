// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListParametersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListParametersResponseBody
	GetNextToken() *string
	SetParameters(v []*ListParametersResponseBodyParameters) *ListParametersResponseBody
	GetParameters() []*ListParametersResponseBodyParameters
	SetRequestId(v string) *ListParametersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListParametersResponseBody
	GetTotalCount() *int32
}

type ListParametersResponseBody struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the common parameter.
	Parameters []*ListParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListParametersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListParametersResponseBody) GetParameters() []*ListParametersResponseBodyParameters {
	return s.Parameters
}

func (s *ListParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParametersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParametersResponseBody) SetMaxResults(v int32) *ListParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParametersResponseBody) SetNextToken(v string) *ListParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParametersResponseBody) SetParameters(v []*ListParametersResponseBodyParameters) *ListParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *ListParametersResponseBody) SetRequestId(v string) *ListParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParametersResponseBody) SetTotalCount(v int32) *ListParametersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListParametersResponseBodyParameters struct {
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
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The common parameter ID.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *string `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags added to the common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListParametersResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListParametersResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *ListParametersResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *ListParametersResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *ListParametersResponseBodyParameters) GetParameterVersion() *string {
	return s.ParameterVersion
}

func (s *ListParametersResponseBodyParameters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListParametersResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *ListParametersResponseBodyParameters) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListParametersResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *ListParametersResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListParametersResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListParametersResponseBodyParameters) SetCreatedBy(v string) *ListParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetCreatedDate(v string) *ListParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetDescription(v string) *ListParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetId(v string) *ListParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetName(v string) *ListParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetParameterVersion(v string) *ListParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetResourceGroupId(v string) *ListParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetShareType(v string) *ListParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListParametersResponseBodyParameters) SetType(v string) *ListParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedBy(v string) *ListParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedDate(v string) *ListParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *ListParametersResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
