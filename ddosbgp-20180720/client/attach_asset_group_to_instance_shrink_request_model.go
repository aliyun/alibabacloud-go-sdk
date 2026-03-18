// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAssetGroupToInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetGroupListShrink(v string) *AttachAssetGroupToInstanceShrinkRequest
	GetAssetGroupListShrink() *string
	SetInstanceId(v string) *AttachAssetGroupToInstanceShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachAssetGroupToInstanceShrinkRequest
	GetRegionId() *string
}

type AttachAssetGroupToInstanceShrinkRequest struct {
	// The information about the asset to be associated.
	//
	// This parameter is required.
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	// The ID of the instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachAssetGroupToInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceShrinkRequest) GetAssetGroupListShrink() *string {
	return s.AssetGroupListShrink
}

func (s *AttachAssetGroupToInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachAssetGroupToInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
