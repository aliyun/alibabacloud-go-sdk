// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParametersByPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetSecretParametersByPathResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetSecretParametersByPathResponseBody
	GetNextToken() *string
	SetParameters(v []*GetSecretParametersByPathResponseBodyParameters) *GetSecretParametersByPathResponseBody
	GetParameters() []*GetSecretParametersByPathResponseBodyParameters
	SetRequestId(v string) *GetSecretParametersByPathResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetSecretParametersByPathResponseBody
	GetTotalCount() *int32
}

type GetSecretParametersByPathResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the encryption parameters.
	Parameters []*GetSecretParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25156E99-7437-4590-AA58-2ACA17DE405C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSecretParametersByPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetSecretParametersByPathResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetSecretParametersByPathResponseBody) GetParameters() []*GetSecretParametersByPathResponseBodyParameters {
	return s.Parameters
}

func (s *GetSecretParametersByPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretParametersByPathResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSecretParametersByPathResponseBody) SetMaxResults(v int32) *GetSecretParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetNextToken(v string) *GetSecretParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetParameters(v []*GetSecretParametersByPathResponseBodyParameters) *GetSecretParametersByPathResponseBody {
	s.Parameters = v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetRequestId(v string) *GetSecretParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetTotalCount(v int32) *GetSecretParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSecretParametersByPathResponseBodyParameters struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// {\\"AllowedPattern\\": \\"^[a-g]*$\\"}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-10-21T06:22:48Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 090xxbex-xexx-xxxx-axfc-ddxxcxxxxcex
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// mySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The data type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 2020-10-21T06:22:48Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParametersByPathResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetConstraints() *string {
	return s.Constraints
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetKeyId() *string {
	return s.KeyId
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetSecretParametersByPathResponseBodyParameters) GetValue() *string {
	return s.Value
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetConstraints(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetDescription(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetKeyId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetName(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetShareType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetValue(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
