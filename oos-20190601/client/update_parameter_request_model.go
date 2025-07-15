// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateParameterRequest
	GetDescription() *string
	SetName(v string) *UpdateParameterRequest
	GetName() *string
	SetRegionId(v string) *UpdateParameterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateParameterRequest
	GetResourceGroupId() *string
	SetTags(v string) *UpdateParameterRequest
	GetTags() *string
	SetValue(v string) *UpdateParameterRequest
	GetValue() *string
}

type UpdateParameterRequest struct {
	// The description of the common parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the common parameter. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags to be added to common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the common parameter. The value must be 1 to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateParameterRequest) GetName() *string {
	return s.Name
}

func (s *UpdateParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateParameterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateParameterRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateParameterRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateParameterRequest) SetDescription(v string) *UpdateParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterRequest) SetName(v string) *UpdateParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateParameterRequest) SetRegionId(v string) *UpdateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateParameterRequest) SetResourceGroupId(v string) *UpdateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateParameterRequest) SetTags(v string) *UpdateParameterRequest {
	s.Tags = &v
	return s
}

func (s *UpdateParameterRequest) SetValue(v string) *UpdateParameterRequest {
	s.Value = &v
	return s
}

func (s *UpdateParameterRequest) Validate() error {
	return dara.Validate(s)
}
