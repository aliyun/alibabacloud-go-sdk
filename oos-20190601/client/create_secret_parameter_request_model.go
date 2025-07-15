// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSecretParameterRequest
	GetClientToken() *string
	SetConstraints(v string) *CreateSecretParameterRequest
	GetConstraints() *string
	SetDKMSInstanceId(v string) *CreateSecretParameterRequest
	GetDKMSInstanceId() *string
	SetDescription(v string) *CreateSecretParameterRequest
	GetDescription() *string
	SetKeyId(v string) *CreateSecretParameterRequest
	GetKeyId() *string
	SetName(v string) *CreateSecretParameterRequest
	GetName() *string
	SetRegionId(v string) *CreateSecretParameterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSecretParameterRequest
	GetResourceGroupId() *string
	SetTags(v map[string]interface{}) *CreateSecretParameterRequest
	GetTags() map[string]interface{}
	SetType(v string) *CreateSecretParameterRequest
	GetType() *string
	SetValue(v string) *CreateSecretParameterRequest
	GetValue() *string
}

type CreateSecretParameterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the encryption parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the encryption parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the encryption parameter.
	//
	// 	- MaxLength: The maximum length of the encryption parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":"secretparameter","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key ID of Key Management Service (KMS) that is used to encrypt the parameter.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter. Set the value to Secret.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the encryption parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecretParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecretParameterRequest) GetConstraints() *string {
	return s.Constraints
}

func (s *CreateSecretParameterRequest) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateSecretParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretParameterRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateSecretParameterRequest) GetName() *string {
	return s.Name
}

func (s *CreateSecretParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecretParameterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecretParameterRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateSecretParameterRequest) GetType() *string {
	return s.Type
}

func (s *CreateSecretParameterRequest) GetValue() *string {
	return s.Value
}

func (s *CreateSecretParameterRequest) SetClientToken(v string) *CreateSecretParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecretParameterRequest) SetConstraints(v string) *CreateSecretParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterRequest) SetDKMSInstanceId(v string) *CreateSecretParameterRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetDescription(v string) *CreateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterRequest) SetKeyId(v string) *CreateSecretParameterRequest {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetName(v string) *CreateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterRequest) SetRegionId(v string) *CreateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetResourceGroupId(v string) *CreateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetTags(v map[string]interface{}) *CreateSecretParameterRequest {
	s.Tags = v
	return s
}

func (s *CreateSecretParameterRequest) SetType(v string) *CreateSecretParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterRequest) SetValue(v string) *CreateSecretParameterRequest {
	s.Value = &v
	return s
}

func (s *CreateSecretParameterRequest) Validate() error {
	return dara.Validate(s)
}
