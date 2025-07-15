// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v map[string]interface{}) *TagResourcesRequest
	GetResourceIds() map[string]interface{}
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v map[string]interface{}) *TagResourcesRequest
	GetTags() map[string]interface{}
}

type TagResourcesRequest struct {
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
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceIds() map[string]interface{} {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v map[string]interface{}) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v map[string]interface{}) *TagResourcesRequest {
	s.Tags = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
