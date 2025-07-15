// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvalidParameters(v []*string) *GetParametersResponseBody
	GetInvalidParameters() []*string
	SetParameters(v []*GetParametersResponseBodyParameters) *GetParametersResponseBody
	GetParameters() []*GetParametersResponseBodyParameters
	SetRequestId(v string) *GetParametersResponseBody
	GetRequestId() *string
}

type GetParametersResponseBody struct {
	// Invalid parameters.
	InvalidParameters []*string `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	// The information about the common parameters.
	Parameters []*GetParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2597E94B-5346-42D1-BB58-D3333EDD0975
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBody) GetInvalidParameters() []*string {
	return s.InvalidParameters
}

func (s *GetParametersResponseBody) GetParameters() []*GetParametersResponseBodyParameters {
	return s.Parameters
}

func (s *GetParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParametersResponseBody) SetInvalidParameters(v []*string) *GetParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetParametersResponseBody) SetParameters(v []*GetParametersResponseBodyParameters) *GetParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *GetParametersResponseBody) SetRequestId(v string) *GetParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetParametersResponseBodyParameters struct {
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
	// 2020-10-22T03:30:45Z
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
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
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
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
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
	// The time when the parameter was updated.
	//
	// example:
	//
	// 2020-10-22T03:30:45Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// parameter,parameter1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParametersResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBodyParameters) GetConstraints() *string {
	return s.Constraints
}

func (s *GetParametersResponseBodyParameters) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetParametersResponseBodyParameters) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetParametersResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *GetParametersResponseBodyParameters) GetId() *string {
	return s.Id
}

func (s *GetParametersResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *GetParametersResponseBodyParameters) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetParametersResponseBodyParameters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetParametersResponseBodyParameters) GetShareType() *string {
	return s.ShareType
}

func (s *GetParametersResponseBodyParameters) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetParametersResponseBodyParameters) GetType() *string {
	return s.Type
}

func (s *GetParametersResponseBodyParameters) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetParametersResponseBodyParameters) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetParametersResponseBodyParameters) GetValue() *string {
	return s.Value
}

func (s *GetParametersResponseBodyParameters) SetConstraints(v string) *GetParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedBy(v string) *GetParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedDate(v string) *GetParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetDescription(v string) *GetParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetId(v string) *GetParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetName(v string) *GetParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetParameterVersion(v int32) *GetParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetResourceGroupId(v string) *GetParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetShareType(v string) *GetParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetParametersResponseBodyParameters) SetType(v string) *GetParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedBy(v string) *GetParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedDate(v string) *GetParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetValue(v string) *GetParametersResponseBodyParameters {
	s.Value = &v
	return s
}

func (s *GetParametersResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
