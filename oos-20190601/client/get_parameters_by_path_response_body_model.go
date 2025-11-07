// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersByPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetParametersByPathResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetParametersByPathResponseBody
	GetNextToken() *string
	SetParameters(v []*GetParametersByPathResponseBodyParameters) *GetParametersByPathResponseBody
	GetParameters() []*GetParametersByPathResponseBodyParameters
	SetRequestId(v string) *GetParametersByPathResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetParametersByPathResponseBody
	GetTotalCount() *int32
}

type GetParametersByPathResponseBody struct {
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
	// The information about the common parameters.
	Parameters []*GetParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25156E99-7437-4590-AA58-2ACA17DE405C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetParametersByPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetParametersByPathResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetParametersByPathResponseBody) GetParameters() []*GetParametersByPathResponseBodyParameters {
	return s.Parameters
}

func (s *GetParametersByPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParametersByPathResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetParametersByPathResponseBody) SetMaxResults(v int32) *GetParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetNextToken(v string) *GetParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetParameters(v []*GetParametersByPathResponseBodyParameters) *GetParametersByPathResponseBody {
	s.Parameters = v
	return s
}

func (s *GetParametersByPathResponseBody) SetRequestId(v string) *GetParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetTotalCount(v int32) *GetParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetParametersByPathResponseBody) Validate() error {
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

type GetParametersByPathResponseBodyParameters struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// {\\"MaxLength\\": 2}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
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
	// 2020-10-21T04:03:12Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// myParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags added to the common parameters.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the common parameter.
	//
	// example:
	//
	// StringList
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 2020-10-21T04:03:12Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// "parameter1,parameter2"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParametersByPathResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBodyParameters) GetConstraints() *string {
	return s.Constraints
}

func (s *GetParametersByPathResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetParametersByPathResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetParametersByPathResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *GetParametersByPathResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *GetParametersByPathResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *GetParametersByPathResponseBodyParameters) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetParametersByPathResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *GetParametersByPathResponseBodyParameters) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetParametersByPathResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *GetParametersByPathResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetParametersByPathResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetParametersByPathResponseBodyParameters) GetValue() *string {
	return s.Value
}

func (s *GetParametersByPathResponseBodyParameters) SetConstraints(v string) *GetParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetDescription(v string) *GetParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetId(v string) *GetParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetName(v string) *GetParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetShareType(v string) *GetParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetTags(v map[string]interface{}) *GetParametersByPathResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetType(v string) *GetParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetValue(v string) *GetParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
