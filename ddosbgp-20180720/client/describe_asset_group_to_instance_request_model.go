// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupToInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAssetGroupToInstanceRequest
	GetInstanceId() *string
	SetMemberUid(v string) *DescribeAssetGroupToInstanceRequest
	GetMemberUid() *string
	SetName(v string) *DescribeAssetGroupToInstanceRequest
	GetName() *string
	SetRegion(v string) *DescribeAssetGroupToInstanceRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeAssetGroupToInstanceRequest
	GetRegionId() *string
	SetType(v string) *DescribeAssetGroupToInstanceRequest
	GetType() *string
}

type DescribeAssetGroupToInstanceRequest struct {
	// The ID of the instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// example:
	//
	// ddosbgp-cn-7212zaa5v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **waf**: WAF instance
	//
	// 	- **ga**: Global Accelerator (GA) instance
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupToInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAssetGroupToInstanceRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAssetGroupToInstanceRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAssetGroupToInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetGroupToInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAssetGroupToInstanceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAssetGroupToInstanceRequest) SetInstanceId(v string) *DescribeAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetMemberUid(v string) *DescribeAssetGroupToInstanceRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetName(v string) *DescribeAssetGroupToInstanceRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegion(v string) *DescribeAssetGroupToInstanceRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegionId(v string) *DescribeAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetType(v string) *DescribeAssetGroupToInstanceRequest {
	s.Type = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) Validate() error {
	return dara.Validate(s)
}
