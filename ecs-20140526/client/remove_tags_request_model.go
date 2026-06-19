// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *RemoveTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveTagsRequest
	GetRegionId() *string
	SetResourceId(v string) *RemoveTagsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *RemoveTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *RemoveTagsRequest
	GetResourceType() *string
	SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest
	GetTag() []*RemoveTagsRequestTag
}

type RemoveTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region to which the resource belongs. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource from which you want to unbind tags. For example, when the resource type (ResourceType) is instance, the resource ID is the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-946ntx4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - instance: ECS instance.
	//
	// - disk: disk.
	//
	// - snapshot: snapshot.
	//
	// - image: image.
	//
	// - securitygroup: security group.
	//
	// - volume: storage volume.
	//
	// - eni: elastic network interface (ENI).
	//
	// - ddh: dedicated host.
	//
	// - keypair: SSH key pair.
	//
	// - launchtemplate: launch template.
	//
	// - reservedinstance: reserved instance.
	//
	// - snapshotpolicy: automatic snapshot policy.
	//
	// All valid values are in lowercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// snapshot
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	Tag []*RemoveTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s RemoveTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveTagsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RemoveTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RemoveTagsRequest) GetTag() []*RemoveTagsRequestTag {
	return s.Tag
}

func (s *RemoveTagsRequest) SetOwnerId(v int64) *RemoveTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetRegionId(v string) *RemoveTagsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceId(v string) *RemoveTagsRequest {
	s.ResourceId = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerAccount(v string) *RemoveTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerId(v int64) *RemoveTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceType(v string) *RemoveTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *RemoveTagsRequest) SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest {
	s.Tag = v
	return s
}

func (s *RemoveTagsRequest) Validate() error {
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

type RemoveTagsRequestTag struct {
	// The tag key of the resource. Valid values of N: 1 to 20. The tag key cannot be an empty string once specified. The tag key can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. Valid values of N: 1 to 20. The tag value can be an empty string once specified. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RemoveTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *RemoveTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *RemoveTagsRequestTag) SetKey(v string) *RemoveTagsRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsRequestTag) SetValue(v string) *RemoveTagsRequestTag {
	s.Value = &v
	return s
}

func (s *RemoveTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
