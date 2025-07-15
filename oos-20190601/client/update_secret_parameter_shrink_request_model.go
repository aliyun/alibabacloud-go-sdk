// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretParameterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSecretParameterShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateSecretParameterShrinkRequest
	GetName() *string
	SetRegionId(v string) *UpdateSecretParameterShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateSecretParameterShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *UpdateSecretParameterShrinkRequest
	GetTagsShrink() *string
	SetValue(v string) *UpdateSecretParameterShrinkRequest
	GetValue() *string
}

type UpdateSecretParameterShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSecretParameterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecretParameterShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSecretParameterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSecretParameterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateSecretParameterShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateSecretParameterShrinkRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateSecretParameterShrinkRequest) SetDescription(v string) *UpdateSecretParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetName(v string) *UpdateSecretParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetRegionId(v string) *UpdateSecretParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetResourceGroupId(v string) *UpdateSecretParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetTagsShrink(v string) *UpdateSecretParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetValue(v string) *UpdateSecretParameterShrinkRequest {
	s.Value = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
