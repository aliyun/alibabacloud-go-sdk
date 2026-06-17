// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeRiskEventGroupResponseBodyDataList) *DescribeRiskEventGroupResponseBody
	GetDataList() []*DescribeRiskEventGroupResponseBodyDataList
	SetRequestId(v string) *DescribeRiskEventGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRiskEventGroupResponseBody
	GetTotalCount() *int32
}

type DescribeRiskEventGroupResponseBody struct {
	// The list of returned data.
	DataList []*DescribeRiskEventGroupResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558FE7E6EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of risk events.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskEventGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBody) GetDataList() []*DescribeRiskEventGroupResponseBodyDataList {
	return s.DataList
}

func (s *DescribeRiskEventGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRiskEventGroupResponseBody) SetDataList(v []*DescribeRiskEventGroupResponseBodyDataList) *DescribeRiskEventGroupResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetRequestId(v string) *DescribeRiskEventGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetTotalCount(v int32) *DescribeRiskEventGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskEventGroupResponseBodyDataList struct {
	// The name of the attacked application.
	//
	// example:
	//
	// MySql
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The type of the attack. Valid values:
	//
	// - **1**: abnormal connection
	//
	// - **2**: command execution
	//
	// - **3**: brute-force attack
	//
	// - **4**: scan
	//
	// - **5**: other
	//
	// - **6**: information leakage
	//
	// - **7**: DoS attack
	//
	// - **8**: overflow attack
	//
	// - **9**: web attack
	//
	// - **10**: backdoor trojan
	//
	// - **11**: virus or worm
	//
	// - **12**: mining behavior
	//
	// - **13**: reverse shell
	//
	// example:
	//
	// 1
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The description of the intrusion prevention event.
	//
	// example:
	//
	// A directory traversal attack was detected in Web access for HTTP requests
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The traffic direction of the intrusion prevention event. Valid values:
	//
	// - **in**: inbound
	//
	// - **out**: outbound
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address. The intrusion prevention event contains this destination IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The number of intrusion prevention events.
	//
	// example:
	//
	// 100
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The ID of the intrusion prevention event.
	//
	// example:
	//
	// 2b58efae-4c4b-4d96-9544-a586fb1f****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the intrusion prevention event.
	//
	// example:
	//
	// Web Directory Traversal Attack
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The time when the intrusion event first occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1534408189
	FirstEventTime *int32 `json:"FirstEventTime,omitempty" xml:"FirstEventTime,omitempty"`
	// The geolocation information of the IP address. This struct contains the **CityId**, **CityName**, **CountryId**, and **CountryName*	- parameters.<br>
	//
	// **CityId*	- indicates the city ID of the IP address. **CityName*	- indicates the city name of the IP address. **CountryId*	- indicates the country ID of the IP address. **CountryName*	- indicates the country name of the IP address.<br>
	IPLocationInfo *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo `json:"IPLocationInfo,omitempty" xml:"IPLocationInfo,omitempty" type:"Struct"`
	// The time when the intrusion prevention event last occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1534408267
	LastEventTime *int32 `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	// The information about the private IP address in the intrusion prevention event. The value is an array that consists of the **RegionNo**, **ResourceInstanceId**, **ResourceInstanceName**, and **ResourcePrivateIP*	- parameters.<br>
	//
	// **RegionNo*	- indicates the region ID of the IP address. **ResourceInstanceId*	- indicates the ID of the instance to which the IP address belongs. **ResourceInstanceName*	- indicates the name of the instance to which the IP address belongs. **ResourcePrivateIP*	- indicates the IP address.<br>
	ResourcePrivateIPList []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList `json:"ResourcePrivateIPList,omitempty" xml:"ResourcePrivateIPList,omitempty" type:"Repeated"`
	// The type of the public IP address. Valid values:
	//
	// - **EIP**: an elastic IP address (EIP)
	//
	// - **EcsPublicIP**: an ECS public IP address
	//
	// - **EcsEIP**: an ECS EIP
	//
	// - **NatPublicIP**: a NAT public IP address
	//
	// - **NatEIP**: a NAT EIP
	//
	// - **SlbPublicIp**: an SLB public IP address
	//
	// example:
	//
	// EcsPublicIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the rule that is used to defend against the intrusion prevention event.
	//
	// example:
	//
	// 1000****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The handling status. Valid values:
	//
	// - **1**: Alert
	//
	// - **2**: Block
	//
	// example:
	//
	// 2
	RuleResult *int32 `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The source of the rule that is used to detect the intrusion prevention event. Valid values:
	//
	// - **1**: basic protection
	//
	// - **2**: virtual patching
	//
	// - **4**: threat intelligence
	//
	// example:
	//
	// 1
	RuleSource *int32 `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The source IP address. The intrusion prevention event contains this source IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// Deprecated
	//
	// The tag of the source IP address. This tag is used to identify whether the IP address is a back-to-origin IP address of an Alibaba Cloud service.
	//
	// example:
	//
	// WAF Back-to-origin Address
	SrcIPTag *string `json:"SrcIPTag,omitempty" xml:"SrcIPTag,omitempty"`
	// The list of IP address tags.
	SrcIPTags []*string `json:"SrcIPTags,omitempty" xml:"SrcIPTags,omitempty" type:"Repeated"`
	// The list of source private IP addresses in the intrusion prevention event.
	SrcPrivateIPList []*string `json:"SrcPrivateIPList,omitempty" xml:"SrcPrivateIPList,omitempty" type:"Repeated"`
	// The tag for major event support.
	//
	// example:
	//
	// test-tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The information about the destination VPC of the intrusion prevention event. This struct contains the **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo*	- parameters.<br>
	//
	// **EcsInstanceId*	- indicates the ID of the ECS instance in the VPC. **EcsInstanceName*	- indicates the name of the ECS instance in the VPC. **NetworkInstanceId*	- indicates the ID of the VPC. **NetworkInstanceName*	- indicates the name of the VPC. **RegionNo*	- indicates the region ID of the VPC.<br>
	VpcDstInfo *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo `json:"VpcDstInfo,omitempty" xml:"VpcDstInfo,omitempty" type:"Struct"`
	// The information about the source VPC of the intrusion prevention event. This struct contains the **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo*	- parameters.<br>
	//
	// **EcsInstanceId*	- indicates the ID of the ECS instance in the VPC. **EcsInstanceName*	- indicates the name of the ECS instance in the VPC. **NetworkInstanceId*	- indicates the ID of the VPC. **NetworkInstanceName*	- indicates the name of the VPC. **RegionNo*	- indicates the region ID of the VPC.<br>
	VpcSrcInfo *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo `json:"VpcSrcInfo,omitempty" xml:"VpcSrcInfo,omitempty" type:"Struct"`
	// The risk level of the intrusion prevention event. Valid values:
	//
	// - **1**: low
	//
	// - **2**: medium
	//
	// - **3**: high
	//
	// example:
	//
	// 1
	VulLevel *int32 `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetAttackApp() *string {
	return s.AttackApp
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetAttackType() *int32 {
	return s.AttackType
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetDirection() *string {
	return s.Direction
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetEventCount() *int32 {
	return s.EventCount
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetEventId() *string {
	return s.EventId
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetFirstEventTime() *int32 {
	return s.FirstEventTime
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetIPLocationInfo() *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	return s.IPLocationInfo
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetLastEventTime() *int32 {
	return s.LastEventTime
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetResourcePrivateIPList() []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	return s.ResourcePrivateIPList
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetRuleResult() *int32 {
	return s.RuleResult
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetRuleSource() *int32 {
	return s.RuleSource
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetSrcIPTag() *string {
	return s.SrcIPTag
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetSrcIPTags() []*string {
	return s.SrcIPTags
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetSrcPrivateIPList() []*string {
	return s.SrcPrivateIPList
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetTag() *string {
	return s.Tag
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetVpcDstInfo() *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	return s.VpcDstInfo
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetVpcSrcInfo() *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	return s.VpcSrcInfo
}

func (s *DescribeRiskEventGroupResponseBodyDataList) GetVulLevel() *int32 {
	return s.VulLevel
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackApp(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackApp = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackType(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDescription(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDirection(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDstIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventCount(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventCount = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventName(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetFirstEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.FirstEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetIPLocationInfo(v *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.IPLocationInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetLastEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.LastEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourcePrivateIPList(v []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourcePrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourceType(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleResult(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleSource(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIPTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIPTag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIPTags(v []*string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIPTags = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcPrivateIPList(v []*string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcPrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Tag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcDstInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcDstInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcSrcInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcSrcInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVulLevel(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.VulLevel = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) Validate() error {
	if s.IPLocationInfo != nil {
		if err := s.IPLocationInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcePrivateIPList != nil {
		for _, item := range s.ResourcePrivateIPList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpcDstInfo != nil {
		if err := s.VpcDstInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VpcSrcInfo != nil {
		if err := s.VpcSrcInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRiskEventGroupResponseBodyDataListIPLocationInfo struct {
	// The city ID.
	//
	// example:
	//
	// 510100
	CityId *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// The city name.
	//
	// example:
	//
	// Chengdu
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The country ID.
	//
	// example:
	//
	// CN
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The country name.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GetCityId() *string {
	return s.CityId
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GetCityName() *string {
	return s.CityName
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GetCountryId() *string {
	return s.CountryId
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList struct {
	// The region ID. This parameter indicates the region ID of the private IP address.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-wz92jf4scg2zb74p****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// LD-shenzhen-zy****
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 10.255.XX.XX
	ResourcePrivateIP *string `json:"ResourcePrivateIP,omitempty" xml:"ResourcePrivateIP,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GetResourcePrivateIP() *string {
	return s.ResourcePrivateIP
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourcePrivateIP(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourcePrivateIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskEventGroupResponseBodyDataListVpcDstInfo struct {
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-wz92jf4scg2zb74p****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// LD-shenzhen-zy****
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// VPC-SH-TX****
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The region ID. This parameter indicates the region ID of the destination VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo struct {
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-wz92jf4scg2zb74p****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// LD-shenzhen-zy****
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6e9a9zyokj2ywuo****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// VPC-SH-TX****
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The region ID. This parameter indicates the region ID of the source VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) Validate() error {
	return dara.Validate(s)
}
