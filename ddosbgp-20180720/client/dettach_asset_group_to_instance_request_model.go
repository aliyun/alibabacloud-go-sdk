// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDettachAssetGroupToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetGroupList(v []*DettachAssetGroupToInstanceRequestAssetGroupList) *DettachAssetGroupToInstanceRequest
	GetAssetGroupList() []*DettachAssetGroupToInstanceRequestAssetGroupList
	SetInstanceId(v string) *DettachAssetGroupToInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DettachAssetGroupToInstanceRequest
	GetRegionId() *string
}

type DettachAssetGroupToInstanceRequest struct {
	// The information about the asset that you want to dissociate.
	//
	// This parameter is required.
	AssetGroupList []*DettachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
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

func (s DettachAssetGroupToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequest) GetAssetGroupList() []*DettachAssetGroupToInstanceRequestAssetGroupList {
	return s.AssetGroupList
}

func (s *DettachAssetGroupToInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DettachAssetGroupToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DettachAssetGroupToInstanceRequest) SetAssetGroupList(v []*DettachAssetGroupToInstanceRequestAssetGroupList) *DettachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetRegionId(v string) *DettachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) Validate() error {
	if s.AssetGroupList != nil {
		for _, item := range s.AssetGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DettachAssetGroupToInstanceRequestAssetGroupList struct {
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **waf**: WAF instance
	//
	// 	- **ga**: Global Accelerator (GA) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return dara.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) GetName() *string {
	return s.Name
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) GetRegion() *string {
	return s.Region
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) GetType() *string {
	return s.Type
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) Validate() error {
	return dara.Validate(s)
}
