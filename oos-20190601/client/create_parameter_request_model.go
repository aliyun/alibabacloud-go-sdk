// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateParameterRequest
	GetClientToken() *string
	SetConstraints(v string) *CreateParameterRequest
	GetConstraints() *string
	SetDescription(v string) *CreateParameterRequest
	GetDescription() *string
	SetName(v string) *CreateParameterRequest
	GetName() *string
	SetRegionId(v string) *CreateParameterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateParameterRequest
	GetResourceGroupId() *string
	SetTags(v map[string]interface{}) *CreateParameterRequest
	GetTags() map[string]interface{}
	SetType(v string) *CreateParameterRequest
	GetType() *string
	SetValue(v string) *CreateParameterRequest
	GetValue() *string
}

type CreateParameterRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateParameterRequest) GetConstraints() *string {
	return s.Constraints
}

func (s *CreateParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateParameterRequest) GetName() *string {
	return s.Name
}

func (s *CreateParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateParameterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateParameterRequest) GetType() *string {
	return s.Type
}

func (s *CreateParameterRequest) GetValue() *string {
	return s.Value
}

func (s *CreateParameterRequest) SetClientToken(v string) *CreateParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterRequest) SetConstraints(v string) *CreateParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterRequest) SetDescription(v string) *CreateParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterRequest) SetName(v string) *CreateParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterRequest) SetRegionId(v string) *CreateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterRequest) SetResourceGroupId(v string) *CreateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterRequest) SetTags(v map[string]interface{}) *CreateParameterRequest {
	s.Tags = v
	return s
}

func (s *CreateParameterRequest) SetType(v string) *CreateParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterRequest) SetValue(v string) *CreateParameterRequest {
	s.Value = &v
	return s
}

func (s *CreateParameterRequest) Validate() error {
	return dara.Validate(s)
}
