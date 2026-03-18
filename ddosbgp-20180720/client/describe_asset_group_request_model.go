// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeAssetGroupRequest
	GetName() *string
	SetRegion(v string) *DescribeAssetGroupRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeAssetGroupRequest
	GetRegionId() *string
	SetType(v string) *DescribeAssetGroupRequest
	GetType() *string
}

type DescribeAssetGroupRequest struct {
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
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
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAssetGroupRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAssetGroupRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAssetGroupRequest) SetName(v string) *DescribeAssetGroupRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegion(v string) *DescribeAssetGroupRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegionId(v string) *DescribeAssetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetType(v string) *DescribeAssetGroupRequest {
	s.Type = &v
	return s
}

func (s *DescribeAssetGroupRequest) Validate() error {
	return dara.Validate(s)
}
