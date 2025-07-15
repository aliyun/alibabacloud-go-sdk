// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSecretParameterRequest
	GetDescription() *string
	SetName(v string) *UpdateSecretParameterRequest
	GetName() *string
	SetRegionId(v string) *UpdateSecretParameterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateSecretParameterRequest
	GetResourceGroupId() *string
	SetTags(v map[string]interface{}) *UpdateSecretParameterRequest
	GetTags() map[string]interface{}
	SetValue(v string) *UpdateSecretParameterRequest
	GetValue() *string
}

type UpdateSecretParameterRequest struct {
	// The description of the parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSecretParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecretParameterRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSecretParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSecretParameterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateSecretParameterRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateSecretParameterRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateSecretParameterRequest) SetDescription(v string) *UpdateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetName(v string) *UpdateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetRegionId(v string) *UpdateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetResourceGroupId(v string) *UpdateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetTags(v map[string]interface{}) *UpdateSecretParameterRequest {
	s.Tags = v
	return s
}

func (s *UpdateSecretParameterRequest) SetValue(v string) *UpdateSecretParameterRequest {
	s.Value = &v
	return s
}

func (s *UpdateSecretParameterRequest) Validate() error {
	return dara.Validate(s)
}
