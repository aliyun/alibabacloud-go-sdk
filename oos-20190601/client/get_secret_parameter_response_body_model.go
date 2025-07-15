// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *GetSecretParameterResponseBodyParameter) *GetSecretParameterResponseBody
	GetParameter() *GetSecretParameterResponseBodyParameter
	SetRequestId(v string) *GetSecretParameterResponseBody
	GetRequestId() *string
}

type GetSecretParameterResponseBody struct {
	// The information about the encryption parameter.
	Parameter *GetSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7F14FB7C-C9BE-44AE-92ED-21ACC02FBFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBody) GetParameter() *GetSecretParameterResponseBodyParameter {
	return s.Parameter
}

func (s *GetSecretParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretParameterResponseBody) SetParameter(v *GetSecretParameterResponseBodyParameter) *GetSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetSecretParameterResponseBody) SetRequestId(v string) *GetSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParameterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSecretParameterResponseBodyParameter struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":".*","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:28:47Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-14ed150fdcd048xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key of Key Management Service (KMS) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
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
	// The share type of the encryption parameter.
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
	// 2020-09-01T09:35:17Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *GetSecretParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetSecretParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetSecretParameterResponseBodyParameter) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *GetSecretParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *GetSecretParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *GetSecretParameterResponseBodyParameter) GetKeyId() *string {
	return s.KeyId
}

func (s *GetSecretParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *GetSecretParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *GetSecretParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSecretParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *GetSecretParameterResponseBodyParameter) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetSecretParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *GetSecretParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetSecretParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetSecretParameterResponseBodyParameter) GetValue() *string {
	return s.Value
}

func (s *GetSecretParameterResponseBodyParameter) SetConstraints(v string) *GetSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetDKMSInstanceId(v string) *GetSecretParameterResponseBodyParameter {
	s.DKMSInstanceId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetDescription(v string) *GetSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetId(v string) *GetSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetKeyId(v string) *GetSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetName(v string) *GetSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *GetSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *GetSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetShareType(v string) *GetSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetSecretParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetType(v string) *GetSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetValue(v string) *GetSecretParameterResponseBodyParameter {
	s.Value = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
