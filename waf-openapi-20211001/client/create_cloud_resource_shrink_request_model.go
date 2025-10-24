// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateCloudResourceShrinkRequest
	GetInstanceId() *string
	SetListenShrink(v string) *CreateCloudResourceShrinkRequest
	GetListenShrink() *string
	SetOwnerUserId(v string) *CreateCloudResourceShrinkRequest
	GetOwnerUserId() *string
	SetRedirectShrink(v string) *CreateCloudResourceShrinkRequest
	GetRedirectShrink() *string
	SetRegionId(v string) *CreateCloudResourceShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateCloudResourceShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*CreateCloudResourceShrinkRequestTag) *CreateCloudResourceShrinkRequest
	GetTag() []*CreateCloudResourceShrinkRequestTag
}

type CreateCloudResourceShrinkRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configurations.
	//
	// This parameter is required.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The forwarding configurations.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateCloudResourceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCloudResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudResourceShrinkRequest) GetListenShrink() *string {
	return s.ListenShrink
}

func (s *CreateCloudResourceShrinkRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateCloudResourceShrinkRequest) GetRedirectShrink() *string {
	return s.RedirectShrink
}

func (s *CreateCloudResourceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudResourceShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateCloudResourceShrinkRequest) GetTag() []*CreateCloudResourceShrinkRequestTag {
	return s.Tag
}

func (s *CreateCloudResourceShrinkRequest) SetInstanceId(v string) *CreateCloudResourceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetListenShrink(v string) *CreateCloudResourceShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetOwnerUserId(v string) *CreateCloudResourceShrinkRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetRedirectShrink(v string) *CreateCloudResourceShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetRegionId(v string) *CreateCloudResourceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetResourceManagerResourceGroupId(v string) *CreateCloudResourceShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateCloudResourceShrinkRequest) SetTag(v []*CreateCloudResourceShrinkRequestTag) *CreateCloudResourceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudResourceShrinkRequest) Validate() error {
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

type CreateCloudResourceShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudResourceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCloudResourceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCloudResourceShrinkRequestTag) SetKey(v string) *CreateCloudResourceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudResourceShrinkRequestTag) SetValue(v string) *CreateCloudResourceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCloudResourceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
