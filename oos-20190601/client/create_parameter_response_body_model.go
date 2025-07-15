// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *CreateParameterResponseBodyParameter) *CreateParameterResponseBody
	GetParameter() *CreateParameterResponseBodyParameter
	SetRequestId(v string) *CreateParameterResponseBody
	GetRequestId() *string
}

type CreateParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *CreateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBody) GetParameter() *CreateParameterResponseBodyParameter {
	return s.Parameter
}

func (s *CreateParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParameterResponseBody) SetParameter(v *CreateParameterResponseBodyParameter) *CreateParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *CreateParameterResponseBody) SetRequestId(v string) *CreateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateParameterResponseBodyParameter struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// "{\\"AllowedValues\\":[\\"parameter\\"],\\"AllowedPattern\\":\\"parameter\\",\\"MinLength\\":0,\\"MaxLength\\":20}"
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
	// The type of the common parameter.
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

func (s CreateParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *CreateParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *CreateParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *CreateParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *CreateParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *CreateParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *CreateParameterResponseBodyParameter) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *CreateParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *CreateParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *CreateParameterResponseBodyParameter) SetConstraints(v string) *CreateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedBy(v string) *CreateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedDate(v string) *CreateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetDescription(v string) *CreateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetId(v string) *CreateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetName(v string) *CreateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetShareType(v string) *CreateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetTags(v map[string]interface{}) *CreateParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetType(v string) *CreateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
