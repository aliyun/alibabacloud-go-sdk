// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecretParametersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSecretParametersResponseBody
	GetNextToken() *string
	SetParameters(v []*ListSecretParametersResponseBodyParameters) *ListSecretParametersResponseBody
	GetParameters() []*ListSecretParametersResponseBodyParameters
	SetRequestId(v string) *ListSecretParametersResponseBody
	GetRequestId() *string
}

type ListSecretParametersResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// sPH90GZOVGC6KPDUL0FIIbEtMQHq_19S6_4O_XqA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the parameters.
	Parameters []*ListSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CA9C6248-AF2A-4AE9-A166-88FD901BBB90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSecretParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecretParametersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecretParametersResponseBody) GetParameters() []*ListSecretParametersResponseBodyParameters {
	return s.Parameters
}

func (s *ListSecretParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretParametersResponseBody) SetMaxResults(v int32) *ListSecretParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetNextToken(v string) *ListSecretParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetParameters(v []*ListSecretParametersResponseBodyParameters) *ListSecretParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *ListSecretParametersResponseBody) SetRequestId(v string) *ListSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecretParametersResponseBodyParameters struct {
	// The user who created the parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:28:47Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the parameter.
	//
	// example:
	//
	// p-14ed150fdcd048xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the KMS customer master key (CMK) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *string `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the parameter was updated.
	//
	// example:
	//
	// 2020-09-01T09:35:17Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListSecretParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListSecretParametersResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListSecretParametersResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *ListSecretParametersResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *ListSecretParametersResponseBodyParameters) GetKeyId() *string {
	return s.KeyId
}

func (s *ListSecretParametersResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *ListSecretParametersResponseBodyParameters) GetParameterVersion() *string {
	return s.ParameterVersion
}

func (s *ListSecretParametersResponseBodyParameters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSecretParametersResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *ListSecretParametersResponseBodyParameters) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListSecretParametersResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *ListSecretParametersResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListSecretParametersResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetDescription(v string) *ListSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetId(v string) *ListSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetKeyId(v string) *ListSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetName(v string) *ListSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetParameterVersion(v string) *ListSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *ListSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetShareType(v string) *ListSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetType(v string) *ListSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
