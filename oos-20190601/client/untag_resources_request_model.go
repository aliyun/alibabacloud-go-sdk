// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v map[string]interface{}) *UntagResourcesRequest
	GetResourceIds() map[string]interface{}
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v map[string]interface{}) *UntagResourcesRequest
	GetTagKeys() map[string]interface{}
}

type UntagResourcesRequest struct {
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
	ResourceIds map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
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
	TagKeys map[string]interface{} `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetResourceIds() map[string]interface{} {
	return s.ResourceIds
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKeys() map[string]interface{} {
	return s.TagKeys
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v map[string]interface{}) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v map[string]interface{}) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
