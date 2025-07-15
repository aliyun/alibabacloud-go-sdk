// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceIdsShrink(v string) *TagResourcesShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceType(v string) *TagResourcesShrinkRequest
	GetResourceType() *string
	SetTagsShrink(v string) *TagResourcesShrinkRequest
	GetTagsShrink() *string
}

type TagResourcesShrinkRequest struct {
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
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *TagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *TagResourcesShrinkRequest) SetRegionId(v string) *TagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceIdsShrink(v string) *TagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagsShrink(v string) *TagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
