// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v map[string]interface{}) *ListTagResourcesRequest
	GetResourceIds() map[string]interface{}
	SetResourceRegionId(v string) *ListTagResourcesRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTags(v map[string]interface{}) *ListTagResourcesRequest
	GetTags() map[string]interface{}
}

type ListTagResourcesRequest struct {
	// The IDs of the resources. You can specify up to 20 IDs. Set this parameter to a JSON array.
	//
	// example:
	//
	// ["000e5836-xxxx-xxxx-xxxx-0d6ab2ac4877"]
	ResourceIds map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The region in which the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **application**: Enterprise Distributed Application Service (EDAS) application
	//
	// 	- **cluster**: EDAS cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key-value pairs that specify the tags.
	//
	// 	- You can add up to 20 tags to a resource.
	//
	// 	- The key cannot start with **aliyun*	- or **acs:*	- and cannot contain **http://*	- or **https://**.
	//
	// 	- The tag key or tag value can be up to 128 characters in length, and can contain letters, digits, hyphens (-), commas (,), asterisks (\\*), forward slashes (/), question marks (?), and colons (:).
	//
	// 	- Set this parameter to a JSON array.
	//
	// example:
	//
	// [{"key":"key1","value":"v1"},{"key":"key2","value":"v2"}]
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetResourceIds() map[string]interface{} {
	return s.ResourceIds
}

func (s *ListTagResourcesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListTagResourcesRequest) SetResourceIds(v map[string]interface{}) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceRegionId(v string) *ListTagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v map[string]interface{}) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
