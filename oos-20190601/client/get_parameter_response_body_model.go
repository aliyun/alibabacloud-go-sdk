// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody
	GetParameter() *GetParameterResponseBodyParameter
	SetRequestId(v string) *GetParameterResponseBody
	GetRequestId() *string
}

type GetParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *GetParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BA326372-2A10-4C3B-BE3E-6439DB7557CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBody) GetParameter() *GetParameterResponseBodyParameter {
	return s.Parameter
}

func (s *GetParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParameterResponseBody) SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetParameterResponseBody) SetRequestId(v string) *GetParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParameterResponseBody) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParameterResponseBodyParameter struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["parameter"],"AllowedPattern":"parameter","MinLength":0,"MaxLength":20}\\"
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
	// 2020-09-01T08:01:43Z
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
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m*****
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
	// {"k1":"v1","k2":"v2"}
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
	// The value of the common parameter.
	//
	// example:
	//
	// parameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s GetParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *GetParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *GetParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *GetParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *GetParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *GetParameterResponseBodyParameter) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *GetParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetParameterResponseBodyParameter) GetValue() *string {
	return s.Value
}

func (s *GetParameterResponseBodyParameter) SetConstraints(v string) *GetParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedBy(v string) *GetParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedDate(v string) *GetParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetDescription(v string) *GetParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetId(v string) *GetParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetName(v string) *GetParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetParameterVersion(v int32) *GetParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetResourceGroupId(v string) *GetParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetShareType(v string) *GetParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetParameterResponseBodyParameter) SetType(v string) *GetParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedBy(v string) *GetParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedDate(v string) *GetParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetValue(v string) *GetParameterResponseBodyParameter {
	s.Value = &v
	return s
}

func (s *GetParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
