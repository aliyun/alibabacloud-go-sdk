// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagResourcesRequest
	GetAll() *bool
	SetResourceId(v []*string) *UnTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UnTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*UnTagResourcesRequestTag) *UnTagResourcesRequest
	GetTag() []*UnTagResourcesRequestTag
	SetTagKey(v []*string) *UnTagResourcesRequest
	GetTagKey() []*string
}

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The IDs of the resources from which you want to remove tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-123
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// 	- **acl**: an access control list (ACL)
	//
	// 	- **loadbalancer**: an Application Load Balancer (ALB) instance
	//
	// 	- **securitypolicy**: a security policy
	//
	// 	- **servergroup**: a server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to remove.
	//
	// example:
	//
	// test
	Tag []*UnTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The keys of the tags that you want to remove.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagResourcesRequest) GetTag() []*UnTagResourcesRequestTag {
	return s.Tag
}

func (s *UnTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTag(v []*UnTagResourcesRequestTag) *UnTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagResourcesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UnTagResourcesRequestTag struct {
	// The key of the tag that you want to remove. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to remove. The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UnTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *UnTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *UnTagResourcesRequestTag) SetKey(v string) *UnTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *UnTagResourcesRequestTag) SetValue(v string) *UnTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *UnTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
