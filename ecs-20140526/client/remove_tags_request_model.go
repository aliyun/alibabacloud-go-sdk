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
	// The ID of the region where the resource is located. Call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource from which to remove tags. For example, if ResourceType is set to instance, this parameter is the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-946ntx4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - instance: an ECS instance.
	//
	// - disk: a disk.
	//
	// - snapshot: a snapshot.
	//
	// - image: an image.
	//
	// - securitygroup: a security group.
	//
	// - volume: a volume.
	//
	// - eni: an elastic network interface.
	//
	// - ddh: a dedicated host.
	//
	// - keypair: an SSH key pair.
	//
	// - launchtemplate: a launch template.
	//
	// - reservedinstance: a reserved instance.
	//
	// - snapshotpolicy: an automatic snapshot policy.
	//
	// All values must be in lowercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// snapshot
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// A list of tags.
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
	// The key of the tag.
	//
	// > For compatibility, we recommend that you use the Tag.N.Key parameter.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value can be an empty string and up to 128 characters long. It cannot start with aliyun or acs:, and cannot contain http\\:// or https\\://.
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
