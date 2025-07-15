// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesShrinkRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceType(v string) *UntagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeysShrink(v string) *UntagResourcesShrinkRequest
	GetTagKeysShrink() *string
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags. This parameter takes effect only if TagKeys is left empty. Valid values: true and false. Default value: false. TagKeys is required if this parameter is set to false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources for which you want to modify the resource group. The number of resource IDs is 1 to 50.
	//
	// 	- If you set ResourceType to template, specify ResourceIds in the ["TemplateName1","TemplateName2"] format.
	//
	// 	- If you set ResourceType to parameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to secretparameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to stateconfiguration, specify ResourceIds in the ["StateConfigurationId 1","StateConfigurationId 2"] format.
	//
	// 	- If you set ResourceType to application, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["templateName1","templateName2"]
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource for which you want to modify the resource group. Valid values:
	//
	// 	- template: template.
	//
	// 	- parameter: parameter.
	//
	// 	- secretparameter: encryption parameter.
	//
	// 	- stateconfiguration: desired-state configuration.
	//
	// 	- application: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. The number of keys ranges from 1 to 20.
	//
	// example:
	//
	// ["k1","k2"]
	TagKeysShrink *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *UntagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetRegionId(v string) *UntagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeysShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
