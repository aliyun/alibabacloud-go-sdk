// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *TagResourcesRequest
	GetResourceId() []*string
	SetResourcePrincipalName(v []*string) *TagResourcesRequest
	GetResourcePrincipalName() []*string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest
	GetTag() []*TagResourcesRequestTag
}

type TagResourcesRequest struct {
	// The ID of resource N.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource ID is the ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `ResourceId` and `ResourcePrincipalName`.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The name of resource N.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource name is the name of the RAM user.
	//
	// >  You must specify only one of the following parameters: `ResourceId` and `ResourcePrincipalName`.
	//
	// example:
	//
	// TagResources
	ResourcePrincipalName []*string `json:"ResourcePrincipalName,omitempty" xml:"ResourcePrincipalName,omitempty" type:"Repeated"`
	// The type of the resource. Valid value:
	//
	// 	- user: a Resource Access Management (RAM) user.
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. You cannot specify empty strings as tag keys. The tag key can be up to 128 characters in length. The tag key cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourcePrincipalName() []*string {
	return s.ResourcePrincipalName
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTag() []*TagResourcesRequestTag {
	return s.Tag
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourcePrincipalName(v []*string) *TagResourcesRequest {
	s.ResourcePrincipalName = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
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

type TagResourcesRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. You cannot specify empty strings as tag keys. The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
