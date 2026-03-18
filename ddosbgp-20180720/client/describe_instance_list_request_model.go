// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdList(v string) *DescribeInstanceListRequest
	GetInstanceIdList() *string
	SetInstanceType(v string) *DescribeInstanceListRequest
	GetInstanceType() *string
	SetInstanceTypeList(v []*string) *DescribeInstanceListRequest
	GetInstanceTypeList() []*string
	SetIp(v string) *DescribeInstanceListRequest
	GetIp() *string
	SetIpVersion(v string) *DescribeInstanceListRequest
	GetIpVersion() *string
	SetOrderby(v string) *DescribeInstanceListRequest
	GetOrderby() *string
	SetOrderdire(v string) *DescribeInstanceListRequest
	GetOrderdire() *string
	SetPageNo(v int32) *DescribeInstanceListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeInstanceListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceListRequest
	GetRegionId() *string
	SetRemark(v string) *DescribeInstanceListRequest
	GetRemark() *string
	SetResourceGroupId(v string) *DescribeInstanceListRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest
	GetTag() []*DescribeInstanceListRequestTag
}

type DescribeInstanceListRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// ["ddosbgp-cn-oew1pjrk****"]
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The field that is used to sort the Anti-DDoS Origin instances. Set the value to **expireTime**, which indicates that the instances are sorted based on the expiration time.
	//
	// You can set the **Orderdire*	- parameter to specify the sorting method.
	//
	// example:
	//
	// 0
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The total number of Anti-DDoS Origin instances.
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The sorting method. Valid values:
	//
	// 	- **desc**: the descending order. This is the default value.
	//
	// 	- **asc**: the ascending order.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP address of the object that is protected by the Anti-DDoS Origin instance to query.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance to query resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// expireTime
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	// The tags that are added to the Anti-DDoS Origin instance.
	//
	// example:
	//
	// desc
	Orderdire *string `json:"Orderdire,omitempty" xml:"Orderdire,omitempty"`
	// The protocol type of the IP address asset that is protected by the Anti-DDoS Origin instance to query. Valid values:
	//
	// 	- **Ipv4**: IPv4
	//
	// 	- **Ipv6**: IPv6
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The mitigation plan of the Anti-DDoS Origin instance to query. Valid values:
	//
	// 	- **0**: the Professional mitigation plan
	//
	// 	- **1**: the Enterprise mitigation plan
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag that is added to the Anti-DDoS Origin instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The remarks of the Anti-DDoS Origin instance to query. Fuzzy match is supported.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The key of the tag that is added to the Anti-DDoS Origin instance.
	Tag []*DescribeInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) GetInstanceIdList() *string {
	return s.InstanceIdList
}

func (s *DescribeInstanceListRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceListRequest) GetInstanceTypeList() []*string {
	return s.InstanceTypeList
}

func (s *DescribeInstanceListRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstanceListRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceListRequest) GetOrderby() *string {
	return s.Orderby
}

func (s *DescribeInstanceListRequest) GetOrderdire() *string {
	return s.Orderdire
}

func (s *DescribeInstanceListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstanceListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceListRequest) GetTag() []*DescribeInstanceListRequestTag {
	return s.Tag
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceType(v string) *DescribeInstanceListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceTypeList(v []*string) *DescribeInstanceListRequest {
	s.InstanceTypeList = v
	return s
}

func (s *DescribeInstanceListRequest) SetIp(v string) *DescribeInstanceListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIpVersion(v string) *DescribeInstanceListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderby(v string) *DescribeInstanceListRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderdire(v string) *DescribeInstanceListRequest {
	s.Orderdire = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRegionId(v string) *DescribeInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRemark(v string) *DescribeInstanceListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListRequest) SetResourceGroupId(v string) *DescribeInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstanceListRequest) Validate() error {
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

type DescribeInstanceListRequestTag struct {
	// The mitigation plan of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mitigation plan of the Anti-DDoS Origin instance. Valid values:
	//
	// 	- 0: the Professional mitigation plan.
	//
	// 	- 1: the Enterprise mitigation plan.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceListRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceListRequestTag) SetKey(v string) *DescribeInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceListRequestTag) SetValue(v string) *DescribeInstanceListRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstanceListRequestTag) Validate() error {
	return dara.Validate(s)
}
