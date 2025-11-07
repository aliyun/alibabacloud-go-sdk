// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *UpdateParameterResponseBodyParameter) *UpdateParameterResponseBody
	GetParameter() *UpdateParameterResponseBodyParameter
	SetRequestId(v string) *UpdateParameterResponseBody
	GetRequestId() *string
}

type UpdateParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *UpdateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AF1AE6DE-61C4-435E-8687-072CFACCCEC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBody) GetParameter() *UpdateParameterResponseBodyParameter {
	return s.Parameter
}

func (s *UpdateParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateParameterResponseBody) SetParameter(v *UpdateParameterResponseBodyParameter) *UpdateParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *UpdateParameterResponseBody) SetRequestId(v string) *UpdateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateParameterResponseBody) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateParameterResponseBodyParameter struct {
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
	// update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter ID.
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
	// 2
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
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
	// The tag added to the common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	// 2020-09-01T08:04:23Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *UpdateParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *UpdateParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *UpdateParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *UpdateParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *UpdateParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *UpdateParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *UpdateParameterResponseBodyParameter) GetTags() *string {
	return s.Tags
}

func (s *UpdateParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *UpdateParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *UpdateParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateParameterResponseBodyParameter) SetConstraints(v string) *UpdateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetDescription(v string) *UpdateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetId(v string) *UpdateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetName(v string) *UpdateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetShareType(v string) *UpdateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetTags(v string) *UpdateParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetType(v string) *UpdateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
