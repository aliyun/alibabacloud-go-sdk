// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateParameterShrinkRequest
	GetClientToken() *string
	SetConstraints(v string) *CreateParameterShrinkRequest
	GetConstraints() *string
	SetDescription(v string) *CreateParameterShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateParameterShrinkRequest
	GetName() *string
	SetRegionId(v string) *CreateParameterShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateParameterShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *CreateParameterShrinkRequest
	GetTagsShrink() *string
	SetType(v string) *CreateParameterShrinkRequest
	GetType() *string
	SetValue(v string) *CreateParameterShrinkRequest
	GetValue() *string
}

type CreateParameterShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the common parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the common parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the common parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the common parameter.
	//
	// 	- MaxLength: The maximum length of the common parameter.
	//
	// example:
	//
	// {
	//
	//     "AllowedValues": [
	//
	//         "parameter"
	//
	//     ],
	//
	//     "AllowedPattern": "parameter",
	//
	//     "MinLength": 0,
	//
	//     "MaxLength": 20
	//
	// }
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The description of the common parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the parameter. Valid values: String and StringList.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the common parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// parameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParameterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateParameterShrinkRequest) GetConstraints() *string {
	return s.Constraints
}

func (s *CreateParameterShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateParameterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateParameterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateParameterShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateParameterShrinkRequest) GetValue() *string {
	return s.Value
}

func (s *CreateParameterShrinkRequest) SetClientToken(v string) *CreateParameterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetConstraints(v string) *CreateParameterShrinkRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetDescription(v string) *CreateParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetName(v string) *CreateParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetRegionId(v string) *CreateParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetResourceGroupId(v string) *CreateParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetTagsShrink(v string) *CreateParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetType(v string) *CreateParameterShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetValue(v string) *CreateParameterShrinkRequest {
	s.Value = &v
	return s
}

func (s *CreateParameterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
