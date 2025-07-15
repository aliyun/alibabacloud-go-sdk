// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *UpdateSecretParameterResponseBodyParameter) *UpdateSecretParameterResponseBody
	GetParameter() *UpdateSecretParameterResponseBodyParameter
	SetRequestId(v string) *UpdateSecretParameterResponseBody
	GetRequestId() *string
}

type UpdateSecretParameterResponseBody struct {
	// The information about the parameter.
	Parameter *UpdateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecretParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBody) GetParameter() *UpdateSecretParameterResponseBodyParameter {
	return s.Parameter
}

func (s *UpdateSecretParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecretParameterResponseBody) SetParameter(v *UpdateSecretParameterResponseBodyParameter) *UpdateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *UpdateSecretParameterResponseBody) SetRequestId(v string) *UpdateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretParameterResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSecretParameterResponseBodyParameter struct {
	// The constraints of the parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":".*","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
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
	// 2020-09-01T09:30:36Z
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
	// p-0b0fff9919c946xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of customer master key (CMK) of Key Management Service (KMS) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the parameter.
	//
	// example:
	//
	// 2
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
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
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	// 2020-09-01T09:33:11Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateSecretParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *UpdateSecretParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *UpdateSecretParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateSecretParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecretParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *UpdateSecretParameterResponseBodyParameter) GetKeyId() *string {
	return s.KeyId
}

func (s *UpdateSecretParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *UpdateSecretParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *UpdateSecretParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateSecretParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *UpdateSecretParameterResponseBodyParameter) GetTags() *string {
	return s.Tags
}

func (s *UpdateSecretParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *UpdateSecretParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *UpdateSecretParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateSecretParameterResponseBodyParameter) SetConstraints(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetDescription(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetKeyId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetName(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetShareType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetTags(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
