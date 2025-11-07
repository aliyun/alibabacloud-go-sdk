// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v *CreateSecretParameterResponseBodyParameter) *CreateSecretParameterResponseBody
	GetParameter() *CreateSecretParameterResponseBodyParameter
	SetRequestId(v string) *CreateSecretParameterResponseBody
	GetRequestId() *string
}

type CreateSecretParameterResponseBody struct {
	// The information about the encryption parameter.
	Parameter *CreateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecretParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBody) GetParameter() *CreateSecretParameterResponseBodyParameter {
	return s.Parameter
}

func (s *CreateSecretParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretParameterResponseBody) SetParameter(v *CreateSecretParameterResponseBodyParameter) *CreateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *CreateSecretParameterResponseBody) SetRequestId(v string) *CreateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretParameterResponseBody) Validate() error {
	if s.Parameter != nil {
		if err := s.Parameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecretParameterResponseBodyParameter struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// \\"{ 	"AllowedValues": ["secretparameter"], 	"AllowedPattern": "secretparameter", 	"MinLength": 0, 	"MaxLength": 20 }\\"
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
	// 2020-09-01T09:30:36Z
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
	// p-0b0fff9919c946xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The key ID of KMS that is used to encrypt the parameter.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MyParameter
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
	// 2020-09-01T09:30:36Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateSecretParameterResponseBodyParameter) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBodyParameter) GetConstraints() *string {
	return s.Constraints
}

func (s *CreateSecretParameterResponseBodyParameter) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateSecretParameterResponseBodyParameter) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *CreateSecretParameterResponseBodyParameter) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateSecretParameterResponseBodyParameter) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretParameterResponseBodyParameter) GetId() *string {
	return s.Id
}

func (s *CreateSecretParameterResponseBodyParameter) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateSecretParameterResponseBodyParameter) GetName() *string {
	return s.Name
}

func (s *CreateSecretParameterResponseBodyParameter) GetParameterVersion() *int32 {
	return s.ParameterVersion
}

func (s *CreateSecretParameterResponseBodyParameter) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecretParameterResponseBodyParameter) GetShareType() *string {
	return s.ShareType
}

func (s *CreateSecretParameterResponseBodyParameter) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateSecretParameterResponseBodyParameter) GetType() *string {
	return s.Type
}

func (s *CreateSecretParameterResponseBodyParameter) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *CreateSecretParameterResponseBodyParameter) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *CreateSecretParameterResponseBodyParameter) SetConstraints(v string) *CreateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetDKMSInstanceId(v string) *CreateSecretParameterResponseBodyParameter {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetDescription(v string) *CreateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetId(v string) *CreateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetKeyId(v string) *CreateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetName(v string) *CreateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetShareType(v string) *CreateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetTags(v map[string]interface{}) *CreateSecretParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetType(v string) *CreateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) Validate() error {
	return dara.Validate(s)
}
