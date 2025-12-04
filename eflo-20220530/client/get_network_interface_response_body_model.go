// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetNetworkInterfaceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetNetworkInterfaceResponseBody
	GetCode() *int32
	SetContent(v *GetNetworkInterfaceResponseBodyContent) *GetNetworkInterfaceResponseBody
	GetContent() *GetNetworkInterfaceResponseBodyContent
	SetMessage(v string) *GetNetworkInterfaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNetworkInterfaceResponseBody
	GetRequestId() *string
}

type GetNetworkInterfaceResponseBody struct {
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
	// The response data. (If a resource has dependent resources, the existing dependent resources are returned.)
	Content *GetNetworkInterfaceResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetNetworkInterfaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetNetworkInterfaceResponseBody) GetContent() *GetNetworkInterfaceResponseBodyContent {
	return s.Content
}

func (s *GetNetworkInterfaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkInterfaceResponseBody) SetAccessDeniedDetail(v string) *GetNetworkInterfaceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetCode(v int32) *GetNetworkInterfaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetContent(v *GetNetworkInterfaceResponseBodyContent) *GetNetworkInterfaceResponseBody {
	s.Content = v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetMessage(v string) *GetNetworkInterfaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) SetRequestId(v string) *GetNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkInterfaceResponseBodyContent struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Port
	Ethernet []*string `json:"Ethernet,omitempty" xml:"Ethernet,omitempty" type:"Repeated"`
	// Gateway
	//
	// example:
	//
	// 172.24.20.254
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 203.107.60.69
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// NC Type
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
	// DEFAULT
	NcType *string `json:"NcType,omitempty" xml:"NcType,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-f8z4scmfh0u4ewv6vdd8
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// ENI Name
	//
	// example:
	//
	// bond0
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// masterintranett2fdth5fkoocg
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Secondary Private IP\\&MAC Address Collection
	PrivateIpAddressMacGroup []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup `json:"PrivateIpAddressMacGroup,omitempty" xml:"PrivateIpAddressMacGroup,omitempty" type:"Repeated"`
	// network interface controller private secondary IP limit
	//
	// example:
	//
	// 0
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
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Service network interface controller address
	//
	// example:
	//
	// 01-00-5e-00-00-16
	ServiceMac *string `json:"ServiceMac,omitempty" xml:"ServiceMac,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet (Subnet) basic information
	SubnetBaseInfo *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo `json:"SubnetBaseInfo,omitempty" xml:"SubnetBaseInfo,omitempty" type:"Struct"`
	// The details of the resource tags.
	Tags []*GetNetworkInterfaceResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Basic information of Lingjun network segment (VPD)
	VpdBaseInfo *GetNetworkInterfaceResponseBodyContentVpdBaseInfo `json:"VpdBaseInfo,omitempty" xml:"VpdBaseInfo,omitempty" type:"Struct"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNetworkInterfaceResponseBodyContent) GetEthernet() []*string {
	return s.Ethernet
}

func (s *GetNetworkInterfaceResponseBodyContent) GetGateway() *string {
	return s.Gateway
}

func (s *GetNetworkInterfaceResponseBodyContent) GetIp() *string {
	return s.Ip
}

func (s *GetNetworkInterfaceResponseBodyContent) GetNcType() *string {
	return s.NcType
}

func (s *GetNetworkInterfaceResponseBodyContent) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *GetNetworkInterfaceResponseBodyContent) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *GetNetworkInterfaceResponseBodyContent) GetNodeId() *string {
	return s.NodeId
}

func (s *GetNetworkInterfaceResponseBodyContent) GetPrivateIpAddressMacGroup() []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	return s.PrivateIpAddressMacGroup
}

func (s *GetNetworkInterfaceResponseBodyContent) GetQuota() *int32 {
	return s.Quota
}

func (s *GetNetworkInterfaceResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNetworkInterfaceResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetNetworkInterfaceResponseBodyContent) GetServiceMac() *string {
	return s.ServiceMac
}

func (s *GetNetworkInterfaceResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetNetworkInterfaceResponseBodyContent) GetSubnetBaseInfo() *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	return s.SubnetBaseInfo
}

func (s *GetNetworkInterfaceResponseBodyContent) GetTags() []*GetNetworkInterfaceResponseBodyContentTags {
	return s.Tags
}

func (s *GetNetworkInterfaceResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetNetworkInterfaceResponseBodyContent) GetVpdBaseInfo() *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	return s.VpdBaseInfo
}

func (s *GetNetworkInterfaceResponseBodyContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetNetworkInterfaceResponseBodyContent) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetEthernet(v []*string) *GetNetworkInterfaceResponseBodyContent {
	s.Ethernet = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetGateway(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Gateway = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetIp(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Ip = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNcType(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NcType = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNetworkInterfaceId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NetworkInterfaceId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNetworkInterfaceName(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NetworkInterfaceName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetNodeId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.NodeId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetPrivateIpAddressMacGroup(v []*GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) *GetNetworkInterfaceResponseBodyContent {
	s.PrivateIpAddressMacGroup = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetQuota(v int32) *GetNetworkInterfaceResponseBodyContent {
	s.Quota = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetRegionId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetResourceGroupId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetServiceMac(v string) *GetNetworkInterfaceResponseBodyContent {
	s.ServiceMac = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetStatus(v string) *GetNetworkInterfaceResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetSubnetBaseInfo(v *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) *GetNetworkInterfaceResponseBodyContent {
	s.SubnetBaseInfo = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetTags(v []*GetNetworkInterfaceResponseBodyContentTags) *GetNetworkInterfaceResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetTenantId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetVpdBaseInfo(v *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) *GetNetworkInterfaceResponseBodyContent {
	s.VpdBaseInfo = v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) SetZoneId(v string) *GetNetworkInterfaceResponseBodyContent {
	s.ZoneId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContent) Validate() error {
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

type GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup struct {
	// The instance description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Secondary private MAC address
	//
	// example:
	//
	// 01-00-5e-00-00-16
	IpAddressMac *string `json:"IpAddressMac,omitempty" xml:"IpAddressMac,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-xxxx
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Secondary private IP address
	//
	// example:
	//
	// 172.23.161.57
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The state of the rule.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetDescription() *string {
	return s.Description
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetIpAddressMac() *string {
	return s.IpAddressMac
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetIpName() *string {
	return s.IpName
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetMessage() *string {
	return s.Message
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) GetStatus() *string {
	return s.Status
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetDescription(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Description = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetIpAddressMac(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.IpAddressMac = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetIpName(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.IpName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetMessage(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Message = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetPrivateIpAddress(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.PrivateIpAddress = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) SetStatus(v string) *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup {
	s.Status = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentPrivateIpAddressMacGroup) Validate() error {
	return dara.Validate(s)
}

type GetNetworkInterfaceResponseBodyContentSubnetBaseInfo struct {
	// Network address segment
	//
	// example:
	//
	// 116.233.21.57/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Subnet instance.
	//
	// example:
	//
	// subnet-urb01blo
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The name of the Subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GetSubnetId() *string {
	return s.SubnetId
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) GetSubnetName() *string {
	return s.SubnetName
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetCidr(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetSubnetId(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.SubnetId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) SetSubnetName(v string) *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo {
	s.SubnetName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentSubnetBaseInfo) Validate() error {
	return dara.Validate(s)
}

type GetNetworkInterfaceResponseBodyContentTags struct {
	// The tag key.
	//
	// example:
	//
	// key-test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetNetworkInterfaceResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetNetworkInterfaceResponseBodyContentTags) SetTagKey(v string) *GetNetworkInterfaceResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentTags) SetTagValue(v string) *GetNetworkInterfaceResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}

type GetNetworkInterfaceResponseBodyContentVpdBaseInfo struct {
	// The network segment of the Lingjun subnet.
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
	// 172.18.0.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1678273219000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-ppdunxzc
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-1
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s GetNetworkInterfaceResponseBodyContentVpdBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GetCidr() *string {
	return s.Cidr
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) GetVpdName() *string {
	return s.VpdName
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetCidr(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.Cidr = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetCreateTime(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetVpdId(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.VpdId = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) SetVpdName(v string) *GetNetworkInterfaceResponseBodyContentVpdBaseInfo {
	s.VpdName = &v
	return s
}

func (s *GetNetworkInterfaceResponseBodyContentVpdBaseInfo) Validate() error {
	return dara.Validate(s)
}
