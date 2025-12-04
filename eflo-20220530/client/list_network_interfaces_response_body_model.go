// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListNetworkInterfacesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListNetworkInterfacesResponseBody
	GetCode() *int32
	SetContent(v *ListNetworkInterfacesResponseBodyContent) *ListNetworkInterfacesResponseBody
	GetContent() *ListNetworkInterfacesResponseBodyContent
	SetMessage(v string) *ListNetworkInterfacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNetworkInterfacesResponseBody
	GetRequestId() *string
}

type ListNetworkInterfacesResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListNetworkInterfacesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListNetworkInterfacesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListNetworkInterfacesResponseBody) GetContent() *ListNetworkInterfacesResponseBodyContent {
	return s.Content
}

func (s *ListNetworkInterfacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkInterfacesResponseBody) SetAccessDeniedDetail(v string) *ListNetworkInterfacesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetCode(v int32) *ListNetworkInterfacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetContent(v *ListNetworkInterfacesResponseBodyContent) *ListNetworkInterfacesResponseBody {
	s.Content = v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetMessage(v string) *ListNetworkInterfacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) SetRequestId(v string) *ListNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkInterfacesResponseBodyContent struct {
	// The response parameters.
	Data []*ListNetworkInterfacesResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContent) GetData() []*ListNetworkInterfacesResponseBodyContentData {
	return s.Data
}

func (s *ListNetworkInterfacesResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListNetworkInterfacesResponseBodyContent) SetData(v []*ListNetworkInterfacesResponseBodyContentData) *ListNetworkInterfacesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContent) SetTotal(v int64) *ListNetworkInterfacesResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkInterfacesResponseBodyContentData struct {
	// The time when the activation code was created.
	//
	// example:
	//
	// 1669734207000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The port number of the AD server.
	Ethernet []*string `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	// The gateway.
	//
	// example:
	//
	// 10.0.0.253
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.0.0.13
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The NC type.
	//
	// Valid value:
	//
	// 	- CUSTOM_LNI_INTEGRATION: two-network one-in-one architecture Lingjun hosting network interface controller.
	//
	// 	- CPU: CPU machine.
	//
	// 	- ELASTIC_6.2: Machine
	//
	// 	- GPU: GPU machine.
	//
	// 	- DEFAULT: the old CPU machine.
	//
	// 	- CUSTOM_LNI: two network separation architecture Lingjun hosting network interface controller.
	//
	// example:
	//
	// GPU
	NcType *string `json:"NcType,omitempty" xml:"NcType,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-2ze50voovmtswn328ogm
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The port name.
	//
	// example:
	//
	// bond0
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The ID of the machine to which the instance belongs.
	//
	// example:
	//
	// 2d53f5c204e7476dae69177e7fa6f19c
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary Private IP\\&MAC Address Collection
	PrivateIpAddressMacGroup []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	// network interface controller private secondary IP quota.
	//
	// example:
	//
	// 6
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-1234567890
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The address of the service network interface controller.
	//
	// example:
	//
	// 00-ff-84-15-ba-67
	ServiceMac *string `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet (Subnet) basic information.
	SubnetBaseInfo *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	// List of tags.
	Tags []*ListNetworkInterfacesResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1620939556166279
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Lingjun network segment (VPD) basic information.
	VpdBaseInfo *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetEthernet() []*string {
	return s.Ethernet
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetGateway() *string {
	return s.Gateway
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetIp() *string {
	return s.Ip
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetNcType() *string {
	return s.NcType
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetPrivateIpAddressMacGroup() []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	return s.PrivateIpAddressMacGroup
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetQuota() *int32 {
	return s.Quota
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetServiceMac() *string {
	return s.ServiceMac
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetSubnetBaseInfo() *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	return s.SubnetBaseInfo
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetTags() []*ListNetworkInterfacesResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetVpdBaseInfo() *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *ListNetworkInterfacesResponseBodyContentData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetEthernet(v []*string) *ListNetworkInterfacesResponseBodyContentData {
	s.Ethernet = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetGateway(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Gateway = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetIp(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Ip = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNcType(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NcType = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNetworkInterfaceId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNetworkInterfaceName(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetNodeId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.NodeId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetPrivateIpAddressMacGroup(v []*ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) *ListNetworkInterfacesResponseBodyContentData {
	s.PrivateIpAddressMacGroup = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetQuota(v int32) *ListNetworkInterfacesResponseBodyContentData {
	s.Quota = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetRegionId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetResourceGroupId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetServiceMac(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.ServiceMac = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetStatus(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetSubnetBaseInfo(v *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) *ListNetworkInterfacesResponseBodyContentData {
	s.SubnetBaseInfo = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetTags(v []*ListNetworkInterfacesResponseBodyContentDataTags) *ListNetworkInterfacesResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetTenantId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetVpdBaseInfo(v *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) *ListNetworkInterfacesResponseBodyContentData {
	s.VpdBaseInfo = v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) SetZoneId(v string) *ListNetworkInterfacesResponseBodyContentData {
	s.ZoneId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentData) Validate() error {
	if s.PrivateIpAddressMacGroup != nil {
		for _, item := range s.PrivateIpAddressMacGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubnetBaseInfo != nil {
		if err := s.SubnetBaseInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpdBaseInfo != nil {
		if err := s.VpdBaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup struct {
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Secondary private MAC address.
	//
	// example:
	//
	// 00:25:9d:00:20:20
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// The unique IP identifier.
	//
	// example:
	//
	// sip-1asjd3xg
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The secondary private IP address.
	//
	// example:
	//
	// 10.0.0.14
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The status of the cache reserve instance.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetDescription() *string {
	return s.Description
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetIpAddressMac() *string {
	return s.IpAddressMac
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetIpName() *string {
	return s.IpName
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetMessage() *string {
	return s.Message
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) GetStatus() *string {
	return s.Status
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetDescription(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Description = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetIpAddressMac(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.IpAddressMac = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetIpName(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.IpName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetMessage(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Message = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetPrivateIpAddress(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) SetStatus(v string) *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup {
	s.Status = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataPrivateIpAddressMacGroup) Validate() error {
	return dara.Validate(s)
}

type ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo struct {
	// The network segment of the Subnet.
	//
	// 	- The network segment of the subnet must be a proper subset of the network segment of Lingjun to which it belongs, and the mask must be between 16 bits and 29 bits, which can provide 8 to 65536 addresses. For example, the CIDR block of the Lingjun CIDR block is 192.168.0.0/16, and the CIDR blocks of the subnets under the Lingjun CIDR block are 192.168.0.0/17 to 192.168.0.0/29.
	//
	// 	- The first and last three IP addresses of each subnet segment are reserved by the system. For example, the CIDR blocks of the subnet are 192.168.1.0/24,192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// For more information about CIDR blocks, see the [What is CIDR?](https://www.alibabacloud.com/help/doc-detail/40637.htm#title-gu4-uzk-12r) section in the "Network FAQ" topic.
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// 10.0.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1623656472000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Subnet instance.
	//
	// example:
	//
	// subnet-yjnqn5ef
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) GetSubnetName() *string {
	return s.SubnetName
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetCidr(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetSubnetId(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.SubnetId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) SetSubnetName(v string) *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo {
	s.SubnetName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataSubnetBaseInfo) Validate() error {
	return dara.Validate(s)
}

type ListNetworkInterfacesResponseBodyContentDataTags struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListNetworkInterfacesResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListNetworkInterfacesResponseBodyContentDataTags) SetTagKey(v string) *ListNetworkInterfacesResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataTags) SetTagValue(v string) *ListNetworkInterfacesResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}

type ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo struct {
	// The network segment of Lingjun network segment (VPD).
	//
	// 	- We recommend that you use an RFC private endpoint as the Lingjun CIDR block, such as 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16. In scenarios where the Doringjun CIDR block is connected to each other or where the Lingjun CIDR block is connected to a VPC, make sure that the addresses do not conflict with each other.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets as the primary IPv4 CIDR block of the VPD. This parameter is left empty by default.
	//
	// example:
	//
	// 10.0.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1668158213000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-d3isyds4
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetCidr(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetCreateTime(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetVpdId(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) SetVpdName(v string) *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *ListNetworkInterfacesResponseBodyContentDataVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
