// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDettachAssetGroupToInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetGroupListShrink(v string) *DettachAssetGroupToInstanceShrinkRequest
	GetAssetGroupListShrink() *string
	SetInstanceId(v string) *DettachAssetGroupToInstanceShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *DettachAssetGroupToInstanceShrinkRequest
	GetRegionId() *string
}

type DettachAssetGroupToInstanceShrinkRequest struct {
	// The information about the asset that you want to dissociate.
	//
	// This parameter is required.
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-xxx
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

func (s DettachAssetGroupToInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DettachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceShrinkRequest) GetAssetGroupListShrink() *string {
	return s.AssetGroupListShrink
}

func (s *DettachAssetGroupToInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DettachAssetGroupToInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
