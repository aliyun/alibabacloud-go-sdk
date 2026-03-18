// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAssetGroupToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetGroupList(v []*AttachAssetGroupToInstanceRequestAssetGroupList) *AttachAssetGroupToInstanceRequest
	GetAssetGroupList() []*AttachAssetGroupToInstanceRequestAssetGroupList
	SetInstanceId(v string) *AttachAssetGroupToInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachAssetGroupToInstanceRequest
	GetRegionId() *string
}

type AttachAssetGroupToInstanceRequest struct {
	// The information about the asset to be associated.
	//
	// This parameter is required.
	AssetGroupList []*AttachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
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

func (s AttachAssetGroupToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequest) GetAssetGroupList() []*AttachAssetGroupToInstanceRequestAssetGroupList {
	return s.AssetGroupList
}

func (s *AttachAssetGroupToInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachAssetGroupToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachAssetGroupToInstanceRequest) SetAssetGroupList(v []*AttachAssetGroupToInstanceRequestAssetGroupList) *AttachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetRegionId(v string) *AttachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) Validate() error {
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

type AttachAssetGroupToInstanceRequestAssetGroupList struct {
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 1743970208320***
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset that you want to add. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-test-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return dara.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) GetName() *string {
	return s.Name
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) GetRegion() *string {
	return s.Region
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) GetType() *string {
	return s.Type
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetMemberUid(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.MemberUid = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) Validate() error {
	return dara.Validate(s)
}
