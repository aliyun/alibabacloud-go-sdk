// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceIds(v string) *TagResourcesRequest
	GetResourceIds() *string
	SetResourceRegionId(v string) *TagResourcesRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v string) *TagResourcesRequest
	GetTags() *string
}

type TagResourcesRequest struct {
	// The IDs of the resources. You can specify up to 20 IDs in the format of a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["000e5836-xxxx-xxxx-xxxx-0d6ab2ac4877"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
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
	// The key-value pairs. When you set this parameter, take note of the following limits:
	//
	// 	- You can add up to 20 tags to a resource.
	//
	// 	- The tag key cannot start with **aliyun*	- or **acs:**. It cannot contain **http://*	- or **https://**.
	//
	// 	- The tag key or tag value can be up to 128 characters in length, and can contain letters, digits, hyphens (-), commas (,), asterisks (\\*), forward slashes (/), question marks (?), and colons (:).
	//
	// 	- Set this parameter to a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"key":"key1","value":"v1"},{"key":"key2","value":"v2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *TagResourcesRequest) SetResourceIds(v string) *TagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *TagResourcesRequest) SetResourceRegionId(v string) *TagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
